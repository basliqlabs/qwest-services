// @title           Qwest Authentication API
// @version         1.0
// @description     Authentication service for Qwest platform
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:15340
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"github.com/basliqlabs/qwest-services-auth/delivery/httpserver/userhandler"
	"github.com/basliqlabs/qwest-services-auth/validator"
	"github.com/basliqlabs/qwest-services-auth/validator/authvalidator"

	"github.com/basliqlabs/qwest-services-auth/config"
	"github.com/basliqlabs/qwest-services-auth/delivery/httpserver"
	"github.com/basliqlabs/qwest-services-auth/pkg/logger"
	"github.com/basliqlabs/qwest-services-auth/pkg/translation"
)

func main() {
	cfg := config.Load("config.yml")
	logger.Init(cfg.Logger, cfg.Env)
	translation.Init(cfg.Language)

	mainValidator := validator.New()
	userValidator := authvalidator.New(mainValidator)
	userHandler := userhandler.New(userValidator)

	server := httpserver.New(httpserver.Args{
		UserHandler: *userHandler,
		Config:      cfg,
	})

	server.Start()
}
