package test

import (
	"github.com/streadway/amqp"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDBTest() *gorm.DB {
	dbURI := os.Getenv("POSTGRESDBTEST_URI")
	if dbURI == "" {
		dbURI = "postgres://postgres:@localhost:5432/testdb?sslmode=disable"
	}

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Println("Error DB: " + err.Error() + " " + dbURI)
		os.Exit(1)
	}
	return db
}

//NewRedisTest ...
func NewRedisTest() *redis.Client {
	redisURI := os.Getenv("REDISTEST_URI")
	if redisURI == "" {
		redisURI = "redis://localhost"
	}

	option, err := redis.ParseURL(redisURI)
	if err != nil {
		log.Println("Error Cache: " + err.Error() + " " + redisURI)
		os.Exit(1)
	}

	cache := redis.NewClient(option)
	_, errPing := cache.Ping().Result()
	if errPing != nil {
		log.Println("Error Cache: " + err.Error() + " " + redisURI)
		os.Exit(1)
	}

	return cache
}

func NewQueueTest() *amqp.Channel {
	amqpURI := os.Getenv("AMQPTEST_URI")
	if amqpURI == "" {
		amqpURI = "amqp://guest:guest@localhost:5672"
	}

	amqpConn, err := amqp.Dial(amqpURI)
	if err != nil {
		panic("Error Amqp: " + err.Error())
	}
	//defer amqpConn.Close()

	amqpChannel, err := amqpConn.Channel()
	if err != nil {
		panic("Error amqpChannel: " + err.Error())
	}

	return amqpChannel
}
