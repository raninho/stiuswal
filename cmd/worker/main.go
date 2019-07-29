package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/streadway/amqp"

	"github.com/raninho/stiuswal/internal/lawsuit"
)

var (
	amqpURI    string
	webhookURI string
)

func init() {
	flag.StringVar(&amqpURI, "amqpURI", os.Getenv("AMQP_URI"), "-amqpURI=amqp://guest:guest@localhost:5672/")
	flag.StringVar(&webhookURI, "webhookURI", os.Getenv("WEBHOOK_URI"), "-webhookURI=http://localhost:8080/webhooks/finish")
}

func main() {
	if webhookURI == "" {
		webhookURI = "http://localhost:8080/webhooks/finish"
	}

	if amqpURI == "" {
		amqpURI = "amqp://guest:guest@localhost:5672/"
	}

	amqpConn, err := amqp.Dial(amqpURI)
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

			lss, err := lawsuit.DoCrawler(task.LawsuitNumber)
			lssdWithBytes, _ := json.Marshal(lss)

			payload := new(lawsuit.ProcessFinishedInput)
			payload.OrderID = task.OrderID
			payload.LawsuitNumber = task.LawsuitNumber
			payload.Output = lssdWithBytes

			body := &bytes.Buffer{}
			_ = json.NewEncoder(body).Encode(payload)


			req, err := http.NewRequest("POST", webhookURI, body)
			if err != nil {
				log.Println("NewRequest:", err.Error())
				err = d.Nack(false, true)
				continue
			}

			var client = &http.Client{}

			resp, err := client.Do(req)
			if err != nil {
				log.Println("client:", err.Error())
				err = d.Nack(false, true)
				continue
			}

			if resp.StatusCode != http.StatusOK {
				err = d.Nack(false, true)
				log.Println("resp.StatusCode != http.StatusOK: ", resp.StatusCode)
				continue
			}

			_ = d.Ack(false)
		}

	}()

	<-forever
}
