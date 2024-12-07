package main

import (
	"fmt"
	"github.com/basliqlabs/qwest-services-auth/config"
	"github.com/basliqlabs/qwest-services-auth/delivery/httpserver"
)

// TODO: repository and migration
// TODO: rich error
// TODO: hot reload
// TODO: logging
// TODO: envelope

func main() {
	cfg := config.Load("config.yml")

	fmt.Printf("%+v\n", cfg)

	server := httpserver.New(cfg)
	server.Start()
}
