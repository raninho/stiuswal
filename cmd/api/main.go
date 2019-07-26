package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	apiHandler "github.com/raninho/stiuswal/cmd/api/handler"
	apiRouter "github.com/raninho/stiuswal/cmd/api/router"
)

var (
	port        string
	redisURI    string
	postgresURI string
)

func init() {
	flag.StringVar(&port, "port", os.Getenv("PORT"), "-port=8080")
	flag.StringVar(&postgresURI, "postgresURI", os.Getenv("POSTGRES_URI"), "-postgresURI=postgres://postgres:@localhost/postgres?sslmode=disable")
	flag.StringVar(&redisURI, "redisURI", os.Getenv("REDIS_URI"), "-redisURI=redis://localhost")
}

func main() {
	fmt.Println("API stiuswal")

	flag.Parse()

	db, err := gorm.Open("postgres", postgresURI)
	if err != nil {
		panic("Error DB: " + err.Error())
	}
	defer db.Close()

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

	h := &apiHandler.Handler{Cache: cache, DB: db}
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
