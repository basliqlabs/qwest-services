package main

import (
	"fmt"

	"github.com/basliqlabs/qwest-services-auth/config"
	"github.com/basliqlabs/qwest-services-auth/delivery/httpserver"
)

// TODO: envelope
// TODO: logging
// TODO: auth

func main() {
	cfg := config.Load("config.yml")

	fmt.Printf("%+v\n", cfg)

	server := httpserver.New(cfg)

	server.Start()
}
