package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/quangtran88/anifni-authentication/adapters/services"
	"github.com/quangtran88/anifni-authentication/core/ports"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can not load env", err)
	}

	p := serviceAdapters.NewKafkaProducer()
	err = p.ProduceMultiple(context.Background(), "test", []ports.KafkaMessage{
		{Key: "1", Value: "a"},
		{Key: "2", Value: "b"},
		{Key: "3", Value: "c"},
		{Key: "4", Value: "d"},
	})
	if err != nil {
		log.Fatal(err)
	}

	err = p.ProduceMultiple(context.Background(), "test1", []ports.KafkaMessage{
		{Key: "1", Value: "a"},
		{Key: "2", Value: "b"},
		{Key: "3", Value: "c"},
		{Key: "4", Value: "d"},
	})
	if err != nil {
		log.Fatal(err)
	}
}
