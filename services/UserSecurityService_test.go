package services

import "testing"

var (
	securityService = NewSecurityService()
)

func TestUserSecurityService_HashAndSaltEmptyString(t *testing.T) {
	_, err := securityService.HashAndSalt("")
	if err == nil{
		t.Fail()
	}
}

func TestUserSecurityService_HashAndSaltNotEmptyString(t *testing.T) {
	_, err := securityService.HashAndSalt("123456")
	if err != nil{
		t.Fail()
	}
}

func TestUserSecurityService_ComparePasswordsEqual(t *testing.T) {
	plainPass := "123456"
	result, err := securityService.HashAndSalt(plainPass)
	if err != nil{
		t.Fail()
	}

	err = securityService.ComparePasswords(result, plainPass)
	if err != nil {
		t.Fail()
	}
}

func TestUserSecurityService_ComparePasswordsNotEqual(t *testing.T) {
	plainPass := "123456"
	result, err := securityService.HashAndSalt("12345")
	if err != nil{
		t.Fail()
	}

	err = securityService.ComparePasswords(result, plainPass)
	if err == nil {
		t.Fail()
	}
}