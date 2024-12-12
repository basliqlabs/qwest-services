package main

import (
	"fmt"
	"github.com/basliqlabs/qwest-services-auth/config"
	"github.com/basliqlabs/qwest-services-auth/delivery/httpserver"
)

// TODO: pre-commit (husky)
// TODO: auth
// TODO: language
// TODO: logging
// TODO: envelope

func main() {
	cfg := config.Load("config.yml")

	fmt.Printf("%+v\n", cfg)

	server := httpserver.New(cfg)
	server.Start()
}
