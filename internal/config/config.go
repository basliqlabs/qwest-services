package config

import (
	"github.com/basliqlabs/qwest-services/internal/repository/postgresql"
	"github.com/basliqlabs/qwest-services/pkg/logger"
	"github.com/basliqlabs/qwest-services/pkg/translation"
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
