package config

type HTTPServerConfig struct {
	Port uint `koanf:"port"`
}

type Config struct {
	HttpServer HTTPServerConfig `koanf:"http_server"`
}
