//	@title			Qwest API
//	@version		1.0
//	@description	Qwest services

//	@host		localhost:15340
//	@BasePath	/

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
package main

import (
	"github.com/basliqlabs/qwest-services/internal/delivery/httpserver/userhandler"
	"github.com/basliqlabs/qwest-services/internal/repository/postgresql"
	"github.com/basliqlabs/qwest-services/internal/repository/postgresql/postgresqluser"
	"github.com/basliqlabs/qwest-services/internal/service/userservice"
	"github.com/basliqlabs/qwest-services/internal/validator"
	"github.com/basliqlabs/qwest-services/internal/validator/uservalidator"

	"github.com/basliqlabs/qwest-services/internal/config"
	"github.com/basliqlabs/qwest-services/internal/delivery/httpserver"
	"github.com/basliqlabs/qwest-services/pkg/logger"
	"github.com/basliqlabs/qwest-services/pkg/translation"
)

func main() {
	cfg := config.Load("config.yml")
	logger.Init(cfg.Logger, cfg.Env)
	translation.Init(cfg.Language)

	mainRepo := postgresql.New(cfg.Repository.Postgres)
	userRepo := postgresqluser.New(mainRepo)

	userSvc := userservice.New(userRepo)

	mainValidator := validator.New()
	userValidator := uservalidator.New(mainValidator)
	userHandler := userhandler.New(userValidator, userSvc)

	server := httpserver.New(httpserver.Args{
		UserHandler: *userHandler,
		Config:      cfg,
	})

	server.Start()
}
