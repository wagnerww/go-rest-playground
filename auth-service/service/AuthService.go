package service

import (
	"github.com/wagnerww/go-rest-playground/models/dto/input"
	"github.com/wagnerww/go-rest-playground/models/dto/output"
)

type AuthService interface {
	VerifyUser(input.AuthToken) output.User
}

type AuthServiceImpl struct {
}

func NewAuthService() AuthService {
	return AuthServiceImpl{}
}

func (a AuthServiceImpl) VerifyUser(token input.AuthToken) output.User {

	var user output.User
	user.Id = "123"
	user.Name = "wagner"
	user.Email = "wagnerricardonet@gmail.com"

	return user
}
