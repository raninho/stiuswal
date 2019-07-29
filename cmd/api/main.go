package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/streadway/amqp"

	apiHandler "github.com/raninho/stiuswal/cmd/api/handler"
	apiRouter "github.com/raninho/stiuswal/cmd/api/router"
	"github.com/raninho/stiuswal/internal/lawsuit"
)

var (
	port        string
	redisURI    string
	postgresURI string
	amqpURI     string
)

func init() {
	flag.StringVar(&port, "port", os.Getenv("PORT"), "-port=8080")
	flag.StringVar(&postgresURI, "postgresURI", os.Getenv("DATABASE_URL"), "-postgresURI=postgres://postgres:@localhost/postgres?sslmode=disable")
	flag.StringVar(&redisURI, "redisURI", os.Getenv("REDIS_URL"), "-redisURI=redis://localhost")
	flag.StringVar(&amqpURI, "amqpURI", os.Getenv("CLOUDAMQP_URL"), "-amqpURI=amqp://guest:guest@localhost:5672/")
}

func main() {
	figure.NewFigure("API stiuswal", "", true).Print()

	flag.Parse()

	db, err := gorm.Open("postgres", postgresURI)
	if err != nil {
		panic("Error DB: " + err.Error())
	}
	defer db.Close()

	db.AutoMigrate(lawsuit.Lawsuit{})

	option, err := redis.ParseURL(redisURI)
	if err != nil {
		panic("Error Cache: " + err.Error())
	}

	cache := redis.NewClient(option)
	_, err = cache.Ping().Result()
	if err != nil {
		panic("Error Cache: " + err.Error())
	}
	defer cache.Close()

	amqpConn, err := amqp.Dial(amqpURI)
	if err != nil {
		panic("Error Amqp: " + err.Error())
	}
	defer amqpConn.Close()

	amqpChannel, err := amqpConn.Channel()
	if err != nil {
		panic("Error amqpChannel: " + err.Error())
	}

	h := &apiHandler.Handler{Cache: cache, DB: db, Queue: amqpChannel}
	routes := apiRouter.Router(h)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      routes,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 40,
	}

	if err := s.ListenAndServe(); err != nil {
		os.Exit(1)
	}
}
