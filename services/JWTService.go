package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type IJwtService interface {
	CreateJwtToken(userId uint8) (string, error)
	VerifyToken(tokenString string) (*jwt.Token, error)
}

type JwtService struct{
	secretKey string
}

func NewJwtService() *JwtService {
	return &JwtService{
		secretKey: os.Getenv("SECRET_KEY"),
	}
}

func (service *JwtService) CreateJwtToken(userId uint8) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["userId"] = userId
	atClaims["expireAt"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(service.secretKey))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *JwtService) VerifyToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})

	if err != nil {
		return err
	}

	return nil
}
