package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"
)

type Handler struct {
	DB    *gorm.DB
	Cache *redis.Client
	Queue *amqp.Channel
}

func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintln(w, &buf)
}
