package main

import (
	"fmt"
	"github.com/basliqlabs/qwest-services-auth/delivery/httpserver/userhandler"
	"github.com/basliqlabs/qwest-services-auth/translation"
	"github.com/basliqlabs/qwest-services-auth/validator"
	"github.com/basliqlabs/qwest-services-auth/validator/authvalidator"

	"github.com/basliqlabs/qwest-services-auth/config"
	"github.com/basliqlabs/qwest-services-auth/delivery/httpserver"
)

// TODO: context
// TODO: query params
// TODO: logging
// TODO: auth

func main() {
	cfg := config.Load("config.yml")

	fmt.Printf("%+v\n", cfg)

	translate := translation.New(cfg.Language)
	vldtr := validator.New(*translate)
	userValidator := authvalidator.New(*vldtr)
	userHandler := userhandler.New(userValidator)
	server := httpserver.New(httpserver.Args{
		UserHandler: *userHandler,
		Translate:   translate,
		Config:      cfg,
	})

	server.Start()
}
