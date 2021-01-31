package services

import (
	"golang.org/x/crypto/bcrypt"
)

type IUserSecurityService interface {
	ComparePasswords(hashedPwd string, plainPwd []byte) bool
	HashAndSalt(pwd []byte) string
}

type UserSecurityService struct {
	HashCost int
}

func NewSecurityService() *UserSecurityService {
	return &UserSecurityService{
		HashCost: 8,
	}
}

func (service *UserSecurityService) HashAndSalt(password string) (string, error) {
	byteHash := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(byteHash, service.HashCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (service *UserSecurityService) ComparePasswords(hashedPwd string, plainPwd string) error {
	byteHash := []byte(hashedPwd)
	bytePass := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePass)
	if err != nil {
		return err
	}

	return nil
}
