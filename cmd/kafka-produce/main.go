package main

import (
	"context"
	"github.com/joho/godotenv"
	basePorts "github.com/quangtran88/anifni-base/libs/ports"
	baseServices "github.com/quangtran88/anifni-base/libs/services"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can not load env", err)
	}

	p := baseServices.NewKafkaProducer()
	err = p.ProduceMultiple(context.Background(), "test", []basePorts.EventMessage{
		{Key: "1", Value: "a"},
		{Key: "2", Value: "b"},
		{Key: "3", Value: "c"},
		{Key: "4", Value: "d"},
	})
	if err != nil {
		log.Fatal(err)
	}

	err = p.ProduceMultiple(context.Background(), "test1", []basePorts.EventMessage{
		{Key: "1", Value: "a"},
		{Key: "2", Value: "b"},
		{Key: "3", Value: "c"},
		{Key: "4", Value: "d"},
	})
	if err != nil {
		log.Fatal(err)
	}
}
