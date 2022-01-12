package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()

}
