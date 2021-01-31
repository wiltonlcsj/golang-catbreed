package services

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type IJwtService interface {
	CreateJwtToken(userId uint8) (string, error)
	VerifyToken(tokenString string) (*jwt.Token, error)
}

type JwtService struct {
	secretKey string
}

type JwtCustomClaims struct {
	jwt.StandardClaims
	UserId int64
}

func NewJwtService() *JwtService {
	return &JwtService{
		secretKey: os.Getenv("SECRET_KEY"),
	}
}

func (service *JwtService) CreateJwtToken(userId int64) (string, error) {
	atClaims := JwtCustomClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(service.secretKey))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *JwtService) VerifyToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok {
		return errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return errors.New("JWT is expired")
	}

	return nil
}
