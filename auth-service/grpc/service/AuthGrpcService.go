package service

import (
	"context"
	"fmt"

	"github.com/wagnerww/go-rest-playground/grpc/pb"
	"github.com/wagnerww/go-rest-playground/models/dto/input"
	"github.com/wagnerww/go-rest-playground/service"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (*AuthService) VerifyUser(ctx context.Context, req *pb.AuthToken) (*pb.User, error) {
	fmt.Println("token" + req.Token)

	var authToken input.AuthToken
	authToken.Token = req.Token

	authService := service.NewAuthService()
	userOutput := authService.VerifyUser(authToken)

	return &pb.User{
		Id:    userOutput.Id,
		Name:  userOutput.Name,
		Email: userOutput.Email,
	}, nil
}
