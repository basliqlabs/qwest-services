package httpserver

import (
	"fmt"

	"github.com/basliqlabs/qwest-services-auth/config"
	"github.com/basliqlabs/qwest-services-auth/delivery/httpserver/middleware"
	"github.com/basliqlabs/qwest-services-auth/delivery/httpserver/userhandler"
	"github.com/labstack/echo/v4"
)

type Server struct {
	cfg         config.Config
	Router      *echo.Echo
	userHandler userhandler.Handler
}

type Args struct {
	Config      config.Config
	UserHandler userhandler.Handler
}

func New(args Args) *Server {
	return &Server{
		cfg:         args.Config,
		Router:      echo.New(),
		userHandler: args.UserHandler,
	}
}

func (s *Server) Start() {
	s.Router.Use(middleware.TranslatorMiddleware())
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recovery())

	// * Healthcheck route
	s.Router.GET("/healthcheck", s.healthCheck)

	// * Users
	s.userHandler.SetUserRoutes(s.Router)

	// start the server
	addr := fmt.Sprintf(":%d", s.cfg.HttpServer.Port)

	if err := s.Router.Start(addr); err != nil {
		fmt.Println("router start error:", err)
	}
}
