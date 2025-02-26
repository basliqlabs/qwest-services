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
	"github.com/basliqlabs/qwest-services/internal/delivery/httpserver/middleware"
	"github.com/basliqlabs/qwest-services/internal/delivery/httpserver/userhandler"
	"github.com/basliqlabs/qwest-services/internal/repository/postgresql"
	"github.com/basliqlabs/qwest-services/internal/repository/postgresql/postgresqlrefreshtoken"
	"github.com/basliqlabs/qwest-services/internal/repository/postgresql/postgresqluser"
	"github.com/basliqlabs/qwest-services/internal/service/tokenservice"
	"github.com/basliqlabs/qwest-services/internal/service/userservice"
	"github.com/basliqlabs/qwest-services/internal/validator"
	"github.com/basliqlabs/qwest-services/internal/validator/uservalidator"

	"github.com/basliqlabs/qwest-services/internal/config"
	"github.com/basliqlabs/qwest-services/internal/delivery/httpserver"
	"github.com/basliqlabs/qwest-services/pkg/jwtutil"
	"github.com/basliqlabs/qwest-services/pkg/logger"
	"github.com/basliqlabs/qwest-services/pkg/translation"
)

func main() {
	cfg := config.Load("config.yml")
	logger.Init(cfg.Logger, cfg.Env)
	translation.Init(cfg.Language)
	jwtutil.Init(cfg.Auth.JWT)

	mainRepo := postgresql.New(cfg.Repository.Postgres)
	userRepo := postgresqluser.New(mainRepo)
	tokenRepo := postgresqlrefreshtoken.New(mainRepo)
	userSvc := userservice.New(userRepo, tokenRepo, cfg.Auth)
	tokenSvc := tokenservice.New(tokenRepo, cfg.Auth)
	
	mainValidator := validator.New()
	userValidator := uservalidator.New(mainValidator)
	authMiddleware := middleware.NewAuthMiddleware(tokenRepo)
	userHandler := userhandler.New(userValidator, userSvc, tokenSvc, authMiddleware)

	server := httpserver.New(httpserver.Args{
		UserHandler: *userHandler,
		Config:      cfg,
	})

	server.Start()
}
