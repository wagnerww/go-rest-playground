package service

import (
	"context"
	"fmt"
	"log"

	"github.com/wagnerww/go-rest-playground.git/grpc/pb"
	"google.golang.org/grpc"
)

type AuthService interface {
	validateUser()
}

type AuthServiceeImpl struct {
}

func NewAuthService() AuthService {
	return AuthServiceeImpl{}
}

func (p AuthServiceeImpl) validateUser() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connection to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewAuthServiceClient(connection)

	req := &pb.AuthToken{
		Token: "23423424234",
	}

	res, err := client.VerifyUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Coud not make gRPC request: %v", err)
	}

	fmt.Println(res)

}
