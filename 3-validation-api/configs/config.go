package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	User UserConfig
	Smtp SmtpConfig
	Url  PathConfig
}

type UserConfig struct {
	Email    string
	Password string
	Address  string
}

type PathConfig struct {
	Port string
}

type SmtpConfig struct {
	SmtpServer string
	SmtpPort   string
}

func LoadConfig() *Config {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}
	return &Config{
		User: UserConfig{
			Email:    os.Getenv("EMAIL"),
			Password: os.Getenv("PASSWORD"),
			Address:  os.Getenv("ADDRESS"),
		}, Smtp: SmtpConfig{
			SmtpServer: os.Getenv("SMTPSERVER"),
			SmtpPort:   os.Getenv("SMTPPORT"),
		}, Url: PathConfig{
			Port: os.Getenv("PORT"),
		},
	}
}
