package internal

import (
	"os"
	"strconv"
)

type Config struct {
	Host string `env:"HOST"`
	Port int    `env:"PORT"`
}

func CfgLoad() (*Config, error) {
	cfg := &Config{}

	cfg.Host = getEnv("HOST", "0.0.0.0")

	port, err := strconv.Atoi(getEnv("PORT", "6379"))
	if err != nil {
		return nil, err
	}

	cfg.Port = port

	return cfg, nil
}

func getEnv(key string, fallback string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return v
}
