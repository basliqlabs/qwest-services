package config

import (
	"github.com/basliqlabs/qwest-services/internal/repository/postgresql"
	"github.com/basliqlabs/qwest-services/pkg/jwtutil"
	"github.com/basliqlabs/qwest-services/pkg/logger"
	"github.com/basliqlabs/qwest-services/pkg/translation"
)

type HTTPServerConfig struct {
	Port uint `koanf:"port"`
}

type RepositoryConfig struct {
	Postgres postgresql.Config `koanf:"postgres"`
}

type AuthConfig struct {
	JWT jwtutil.JWTConfig `koanf:"jwt"`
}

type Config struct {
	Env        string             `koanf:"env"`
	HttpServer HTTPServerConfig   `koanf:"http_server"`
	Repository RepositoryConfig   `koanf:"repository"`
	Language   translation.Config `koanf:"language"`
	Logger     logger.Config      `koanf:"logger"`
	Auth       AuthConfig         `koanf:"auth"`
}
