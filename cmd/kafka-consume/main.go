package main

import (
	"github.com/joho/godotenv"
	kafkaHandlers "github.com/quangtran88/anifni-authentication/adapters/handlers/kafka"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can not load env", err)
	}
	kafkaHandlers.InitKafkaHandlers()
}
