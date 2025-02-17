package config

import (
	"github.com/basliqlabs/qwest-services-auth/pkg/logger"
	"github.com/basliqlabs/qwest-services-auth/pkg/translation"
	"github.com/basliqlabs/qwest-services-auth/repository/postgresql"
)

type HTTPServerConfig struct {
	Port uint `koanf:"port"`
}

type RepositoryConfig struct {
	Postgres postgresql.Config `koanf:"postgres"`
}

type Config struct {
	Env        string             `koanf:"env"`
	HttpServer HTTPServerConfig   `koanf:"http_server"`
	Repository RepositoryConfig   `koanf:"repository"`
	Language   translation.Config `koanf:"language"`
	Logger     logger.Config      `koanf:"logger"`
}
