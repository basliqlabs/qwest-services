package httpserver

import (
	"fmt"
	"github.com/basliqlabs/qwest-services-auth/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	cfg    config.Config
	Router *echo.Echo
}

func New(cfg config.Config) *Server {
	return &Server{
		cfg:    cfg,
		Router: echo.New(),
	}
}

func (s *Server) Start() {
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recover())

	// * Healthcheck route
	s.Router.GET("/healthcheck", s.healthCheck)

	// start the server
	addr := fmt.Sprintf(":%d", s.cfg.HttpServer.Port)

	if err := s.Router.Start(addr); err != nil {
		fmt.Println("router start error:", err)
	}
}
