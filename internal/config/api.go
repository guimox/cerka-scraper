package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	BaseURL    string
	Server     ServerConfig
	RabbitMQ   RabbitMQ
}

type RabbitMQ struct {
	User     string
	Password string
	Host     string
	Port     string
}

type ServerConfig struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

const (
	defaultReadTimeout  = 5 * time.Second
	defaultWriteTimeout = 10 * time.Second
	defaultIdleTimeout  = 120 * time.Second
)

func LoadConfig() (Config, error) {
	_ = godotenv.Load()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		return Config{}, fmt.Errorf("SERVER_PORT environment variable is required")
	}

	config := Config{
		ServerPort: port,
		Server: ServerConfig{
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
			IdleTimeout:  defaultIdleTimeout,
		},
	}

	return config, nil
}

func (c Config) GetServerAddress() string {
	return fmt.Sprintf(":%s", c.ServerPort)
}
