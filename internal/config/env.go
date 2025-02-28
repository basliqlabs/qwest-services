package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
)

type EnvironmentVariables struct {
	Files  []string
	envMap map[string]string
	prefix string
}

func NewEnv(prefix string, file string, files ...string) *EnvironmentVariables {
	return &EnvironmentVariables{
		Files:  append([]string{file}, files...),
		prefix: prefix,
		envMap: make(map[string]string),
	}
}

func (e *EnvironmentVariables) Load() {
	err := godotenv.Load(e.Files...)

	if err != nil {
		panic("failed to read environment variables")
	}

	keys := make([]string, 0)

	for _, item := range os.Environ() {
		if strings.HasPrefix(item, e.prefix) {
			keys = append(keys, strings.Split(item, "=")[0])
		}
	}

	for _, key := range keys {
		e.envMap[key] = os.Getenv(key)
	}
}

func (e *EnvironmentVariables) Get(key string) string {
	envKey := fmt.Sprintf("%s%s", e.prefix, key)
	value, ok := e.envMap[envKey]

	if !ok {
		panic(fmt.Sprintf(`failed to get the value of "%v"`, envKey))
	}

	return value
}

func (e *EnvironmentVariables) GetNumber(key string) int {
	envKey := fmt.Sprintf("%s%s", e.prefix, key)
	value, ok := e.envMap[envKey]
	if !ok {
		panic(fmt.Sprintf(`failed to get the value of "%v"`, key))
	}

	v, err := strconv.Atoi(value)

	if err != nil {
		panic(err)
	}

	return v
}
