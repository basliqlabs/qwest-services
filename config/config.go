package config

type AppConfig struct {
}

type Config struct {
	AppConfig AppConfig `koanf:"app_config"`
}
