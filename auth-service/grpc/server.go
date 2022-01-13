package main

import (
	"log"
	"net"

	"github.com/wagnerww/go-rest-playground/grpc/pb"
	"github.com/wagnerww/go-rest-playground/grpc/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, service.NewAuthService())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}

}
