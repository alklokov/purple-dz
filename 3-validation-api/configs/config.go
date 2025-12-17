package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	SMTPConf SMTPConfig
}

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err == nil {
		port, err := strconv.Atoi(os.Getenv("port"))
		if err == nil {
			return &Config{
				SMTPConfig{
					Host:     os.Getenv("host"),
					Port:     port,
					Username: os.Getenv("username"),
					Password: os.Getenv("password"),
				},
			}
		}
	}
	log.Println("Error loading .env file, using default config")
	return makeDefaultConfig()
}

func makeDefaultConfig() *Config {
	return &Config{
		SMTPConfig{
			Host:     "localhost",
			Port:     25,
			Username: "user",
			Password: "pwd",
		},
	}
}
