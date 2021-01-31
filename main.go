package main

import (
	"github.com/subosito/gotenv"
	"github.com/wiltonlcsj/hgBackendTest/router"
)

func init() {
	_ = gotenv.Load()
}

func main() {
	router.New()
}
