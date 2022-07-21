package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
	"trackingApp/graph/models"
)

func GenerateToken(user models.User) (string, error) {

	// Setting up the secret key

	os.Getenv("ACCESS_KEY")

	nowTime := time.Now().Unix()

	// Creating access accessToken header and payload
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"accountId":  user.ID,
		"issuedAt":   nowTime, // issued since now
		"expires_At": time.Now().Add(time.Hour * 24),
	})

	//var err error

	token, err := accessToken.SignedString([]byte(os.Getenv("ACCESS_KEY"))) // signature
	//
	if err != nil {
		//If there is an errorlist in creating the JWT accessToken
		return "", errors.New("jwt Fail")
	}

	return token, nil
}

func ParseToken(tokenParam string) (string, error) {
	token, err := jwt.Parse(tokenParam, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_KEY")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["accountId"].(string)
		return id, nil
	} else {
		return "", err
	}
}
