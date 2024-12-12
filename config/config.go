package config

import (
	"github.com/basliqlabs/qwest-services-auth/repository/postgresql"
	"github.com/basliqlabs/qwest-services-auth/translation"
)

type HTTPServerConfig struct {
	Port uint `koanf:"port"`
}

type RepositoryConfig struct {
	Postgres postgresql.Config `koanf:"postgres"`
}

type Config struct {
	HttpServer HTTPServerConfig   `koanf:"http_server"`
	Repository RepositoryConfig   `koanf:"repository"`
	Language   translation.Config `koanf:"language"`
}
