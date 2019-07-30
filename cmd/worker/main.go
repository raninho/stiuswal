package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/streadway/amqp"

	"github.com/raninho/stiuswal/internal/lawsuit"
)

var (
	queueURI   string
	webhookURI string
)

func init() {
	flag.StringVar(&queueURI, "queueURI", os.Getenv("CLOUDAMQP_URL"), "-queueURI=amqp://guest:guest@localhost:5672/")
	flag.StringVar(&webhookURI, "webhookURI", os.Getenv("WEBHOOK_URI"), "-webhookURI=http://localhost:8080")
}

func main() {
	figure.NewFigure("Work stiuswal", "", true).Print()

	if webhookURI == "" {
		webhookURI = "http://localhost:8080"
	}

	if queueURI == "" {
		queueURI = "amqp://guest:guest@localhost:5672/"
	}

	amqpConn, err := amqp.Dial(queueURI)
	if err != nil {
		panic("Error Amqp: " + err.Error())
	}
	defer amqpConn.Close()

	amqpChannel, err := amqpConn.Channel()
	if err != nil {
		panic("Error amqpChannel: " + err.Error())
	}

	queue, err := amqpChannel.QueueDeclare(lawsuit.QueueProcessName, true, false, false, false, nil)
	if err != nil {
		panic("QueueDeclare: " + err.Error())
	}

	messageChannel, err := amqpChannel.Consume(queue.Name, "", false, false, false, false, nil)

	forever := make(chan bool)

	go func() {

		for d := range messageChannel {
			log.Printf("Received a message: %s", d.Body)

			task := new(lawsuit.Lawsuit)

			err := json.Unmarshal(d.Body, task)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
				err = d.Nack(false, true)
				continue
			}

			if err := CheckWebhookIdempotency(task.OrderID); err != nil {
				if err == ErrIdempotency {
					log.Println("CheckWebhookIdempotency:", err.Error())
					_ = d.Ack(false)
					continue
				}

				_ = d.Nack(false, true)
				continue
			}

			lss, err := lawsuit.DoCrawler(task.LawsuitNumber)
			if err != nil {
				log.Println("Error DoCrawler:", err.Error())
				err = d.Nack(false, true)
				continue
			}

			if err := DoWebhookFinish(task, lss); err != nil {
				log.Println("Error DoWebhookFinish:", err.Error())
				err = d.Nack(false, true)
				continue
			}

			_ = d.Ack(false)
		}

	}()

	<-forever
}

var (
	ErrFinish      = errors.New("Error Webhook Finish")
	ErrIdempotency = errors.New("Error Webhook Idempotency")
	ErrDontKnow    = errors.New("Error Dont Know")
)

func CheckWebhookIdempotency(orderID string) error {
	url := webhookURI + "/webhooks/idempotency/" + orderID

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	var client = &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("client:", err.Error())
		return err
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return ErrIdempotency
		}

		return ErrDontKnow
	}

	return nil
}

func DoWebhookFinish(task *lawsuit.Lawsuit, lss map[string]lawsuit.LawSuitCrawler) error {
	lssdWithBytes, _ := json.Marshal(lss)

	payload := new(lawsuit.ProcessFinishedInput)
	payload.OrderID = task.OrderID
	payload.LawsuitNumber = task.LawsuitNumber
	payload.Output = lssdWithBytes

	body := &bytes.Buffer{}
	_ = json.NewEncoder(body).Encode(payload)

	url := webhookURI + "/webhooks/finish"
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	var client = &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return ErrFinish
	}

	return nil
}
