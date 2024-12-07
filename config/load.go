package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func Load(configPath string) Config {
	cfg := Config{}

	k := koanf.New(".")

	// check the dotenv file and load its data into dotenv `envMap`
	dotenv := NewEnv(EnvPrefix, ".env")

	dotenv.Load()

	// lowest precedence -> fallbacks
	err := k.Load(confmap.Provider(defaultConfig, "."), nil)

	if err != nil {
		panic("failed to initialize the default config")
	}

	// medium precedence -> overwrite variables with yaml configs
	err = k.Load(file.Provider(configPath), yaml.Parser())

	if err != nil {
		panic("failed to read configurations from config.yml")
	}

	// highest precedence -> overwrite variables with what's inside .env file
	err = k.Load(confmap.Provider(map[string]any{}, "."), nil)

	// deserialize all the loaded data into the config variable
	err = k.Unmarshal("", &cfg)

	return cfg
}
