package main

import (
	"fmt"
	"github.com/joho/godotenv"
	grpcHandler "github.com/quangtran88/anifni-authentication/adapters/handlers/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Println("Error loading .env file")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 6000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	grpcHandler.InitGRPCServices(s)

	log.Printf("server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
