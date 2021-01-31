package main

import (
	"github.com/subosito/gotenv"
	"github.com/wiltonlcsj/hgBackendTest/router"
)

func init() {
	_ = gotenv.Load()
}

// @securityDefinitions.apiKey bearerAuth
// @in header
// @name Authorization
func main() {
	router.New()
}
