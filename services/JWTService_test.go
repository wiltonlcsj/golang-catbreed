package services

import "testing"

func TestJwtService_CreateJwtTokenEmptySecret(t *testing.T) {
	jwtServiceEmpty := &JwtService{secretKey: ""}
	_, err := jwtServiceEmpty.CreateJwtToken(1)
	if err == nil{
		t.Fail()
	}
}

func TestJwtService_CreateJwtTokenValidSecret(t *testing.T) {
	jwtServiceNotEmpty := &JwtService{secretKey: "secret-test"}
	_, err := jwtServiceNotEmpty.CreateJwtToken(1)
	if err != nil{
		t.Fail()
	}
}

func TestJwtService_VerifyTokenValidPeriod(t *testing.T) {
	jwtServiceNotEmpty := &JwtService{secretKey: "secret-test"}
	result, err := jwtServiceNotEmpty.CreateJwtToken(1)
	if err != nil{
		t.Fail()
	}

	err = jwtServiceNotEmpty.VerifyToken(result)
	if err != nil {
		t.Fail()
	}
}

func TestJwtService_VerifyTokenExpiredPeriod(t *testing.T) {
	jwtServiceNotEmpty := &JwtService{secretKey: "my-secret"}
	expiredToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTIxNTU1MDUsIlVzZXJJZCI6MX0.9mQTp5EnzAel8S40DOPAWaxI8XPy3TZs_uPJP5IbxPE"
	err := jwtServiceNotEmpty.VerifyToken(expiredToken)
	if err == nil {
		t.Fail()
	}
}

func TestJwtService_VerifyTokenInvalidToken(t *testing.T) {
	jwtServiceNotEmpty := &JwtService{secretKey: "my-secret"}
	invalidToken := "eyJsfaciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTIxN123MDUsIlVzZXJJZCI6MX0.9mQTp5EnzAel8ABCDOPAWaxI8XPy3TZs_uPJP5IbxPE"
	err := jwtServiceNotEmpty.VerifyToken(invalidToken)
	if err == nil {
		t.Fail()
	}
}
