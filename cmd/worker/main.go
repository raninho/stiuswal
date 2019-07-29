package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/raninho/stiuswal/internal/lawsuit"
	"github.com/streadway/amqp"
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

	messageChannel, err := amqpChannel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {

		for d := range messageChannel {
			log.Printf("Received a message: %s", d.Body)

			task := new(lawsuit.Lawsuit)

			err := json.Unmarshal(d.Body, task)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
				break
			}

			fmt.Println(task.OrderID, task.LawsuitNumber)

			payload := new(lawsuit.ProcessFinishedInput)
			payload.OrderID = task.OrderID
			payload.LawsuitNumber = task.LawsuitNumber

			payloadWithBytes, _ := json.Marshal(payload)
			payload.Output = payloadWithBytes

			body := &bytes.Buffer{}
			json.NewEncoder(body).Encode(payload);

			//https://www2.tjal.jus.br/cpopg/search.do?conversationId=&dadosConsulta.localPesquisa.cdLocal=-1&cbPesquisa=NUMPROC&dadosConsulta.tipoNuProcesso=UNIFICADO&numeroDigitoAnoUnificado=0710802-55.2018&foroNumeroUnificado=0001&dadosConsulta.valorConsultaNuUnificado=0710802-55.2018.8.02.0001&dadosConsulta.valorConsulta=&uuidCaptcha=
			//https://www2.tjal.jus.br/cposg5/search.do?conversationId=&paginaConsulta=1&cbPesquisa=NUMPROC&tipoNuProcesso=UNIFICADO&numeroDigitoAnoUnificado=0710802-55.2018&foroNumeroUnificado=0001&dePesquisaNuUnificado=0710802-55.2018.8.02.0001&dePesquisa=&uuidCaptcha=

			//https://esaj.tjms.jus.br/cpopg5/search.do?conversationId=&dadosConsulta.localPesquisa.cdLocal=-1&cbPesquisa=NUMPROC&dadosConsulta.tipoNuProcesso=UNIFICADO&numeroDigitoAnoUnificado=0710802-55.2018&foroNumeroUnificado=0001&dadosConsulta.valorConsultaNuUnificado=0710802-55.2018.8.12.0001&dadosConsulta.valorConsulta=&uuidCaptcha=
			//https://esaj.tjms.jus.br/cposg5/search.do?conversationId=&paginaConsulta=1&localPesquisa.cdLocal=-1&cbPesquisa=NUMPROC&tipoNuProcesso=UNIFICADO&numeroDigitoAnoUnificado=0821901-51.2018&foroNumeroUnificado=0001&dePesquisaNuUnificado=0821901-51.2018.8.12.0001&dePesquisa=&uuidCaptcha=

			var client = &http.Client{}

			req, err := http.NewRequest("POST", webhookURI, body)
			if err != nil {
				log.Println("NewRequest:", err.Error())
			}

			resp, err := client.Do(req)
			if err != nil {
				log.Println("client:", err.Error())
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
