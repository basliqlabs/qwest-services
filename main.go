package main

import (
	"github.com/basliqlabs/qwest-services-auth/delivery/httpserver/userhandler"
	"github.com/basliqlabs/qwest-services-auth/translation"
	"github.com/basliqlabs/qwest-services-auth/validator"
	"github.com/basliqlabs/qwest-services-auth/validator/authvalidator"

	"github.com/basliqlabs/qwest-services-auth/config"
	"github.com/basliqlabs/qwest-services-auth/delivery/httpserver"
	"github.com/basliqlabs/qwest-services-auth/pkg/logger"
)

func main() {
	cfg := config.Load("config.yml")

	logger.Init(cfg.Logger, cfg.Env)

	translate := translation.New(cfg.Language)
	mainValidator := validator.New(translate)

	userValidator := authvalidator.New(mainValidator)
	userHandler := userhandler.New(userValidator)

	server := httpserver.New(httpserver.Args{
		UserHandler: *userHandler,
		Translate:   translate,
		Config:      cfg,
	})

	server.Start()
}
