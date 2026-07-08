package config

import (
	"os"
	"log"
)

type Config struct {
    Database Database
    Server   Server
    JWT      JWT
}

func MustLoad(prefix string) *Config {
	return &Config {
		Server: *MustLoadServer(prefix),
		Database: *MustLoadDatabase(prefix),
		JWT: *MustLoadJWT(prefix),
	}
}

func mustEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("missing environment variable %s", key)
	}
	return value
}
