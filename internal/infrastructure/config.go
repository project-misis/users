package infrastructure

import (
	"log"
	"os"
	"slices"
	"strings"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Port     string `env:"PORT"`
	Database struct {
		Host     string `env:"DB_HOST"`
		Port     string `env:"DB_PORT"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Name     string `env:"DB_NAME"`
	}
	Hash struct {
		Memory      uint32 `env:"HASH_MEMORY"`
		Iterations  uint32
		Parallelism uint8
		SaltLength  uint32
		KeyLength   uint32
	}
	Debug bool
}

func LoadConfig() *Config {
	godotenv.Load()
	var config Config
	if err := env.Parse(&config); err != nil {
		log.Fatal("can't get config")
		return nil
	}
	config.Debug = slices.Contains(
		[]string{"true", "1", "on", "yes"},
		strings.ToLower(os.Getenv("DEBUG")),
	)
	return &config
}
