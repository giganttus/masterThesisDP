package services

import (
	"context"
	"fmt"
	"trackingApp/graph/models"
	"trackingApp/middleware"
)

var jwtID string
var accessToken string

func (s *Services) Login(ctx context.Context, input models.LoginInput) (string, error) {

	user, err := s.UsersRepo.GetUserByField("username", input.Username)

	if err != nil {
		return "", ErrBadCredentials
	}
	if s.UsersRepo.IsActive(user.ID) != true {
		return "", ErrBadCredentials
	}
	//Compare password
	err = user.ComparePassword(input.Password)
	if err != nil {
		return "", ErrBadCredentials
	}

	//Create JWT Token
	token, err2 := middleware.GenerateToken(*user)

	if err2 != nil {
		fmt.Println("Generate token error")
		return "", err2
	}
	return fmt.Sprintf("%s", token), nil

}
