package lawsuit

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/streadway/amqp"
)

const QueueProcessName = "stiuswalProcess"

func Process(db *gorm.DB, channel *amqp.Channel, lawsuitNumber string) (*Lawsuit, error) {
	l := New(lawsuitNumber)

	if err := Create(db, l); err != nil {
		return nil, err
	}

	body, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}

	queue, err := channel.QueueDeclare(QueueProcessName, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	publish := amqp.Publishing{DeliveryMode: amqp.Persistent, ContentType: "text/plain", Body: body}

	if err := channel.Publish("", queue.Name, false, false, publish); err != nil {
		return nil, err
	}

	return l, nil
}

type ProcessFinishedInput struct {
	OrderID       string          `json:"order-id"`
	LawsuitNumber string          `json:"lawsuit-number"`
	Criticized    bool            `json:"criticized"`
	Output        json.RawMessage `json:"output"`
}

func (i *ProcessFinishedInput) Ok() error {
	if _, err := uuid.Parse(i.OrderID); err != nil {
		return errors.New("order-id is invalid")
	}

	if !IsLawsuitNumber(i.LawsuitNumber) {
		return errors.New("lawsuit-number is invalid")
	}

	return nil
}

func ProcessFinished(db *gorm.DB, input *ProcessFinishedInput) (*Lawsuit, error) {
	l, err := GetByOrderIDAndLawsuitNumber(db, input.OrderID, input.LawsuitNumber)
	if err != nil {
		return nil, err
	}

	l.Processed = true
	l.Criticized = input.Criticized
	l.Output = postgres.Jsonb{input.Output}

	if err := Update(db, l); err != nil {
		return nil, err
	}

	return l, nil
}
