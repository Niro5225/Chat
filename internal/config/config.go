package config

import (
	"chat-app/internal/repository"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Bind      string
	Log_level string
	DB        *repository.Config
}

func New_config() *Config {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err.Error())
	}

	log_l := os.Getenv("LOGLEVEL")
	b := fmt.Sprintf("%s:%s", os.Getenv("SERVERHOST"), os.Getenv("SERVERPORT"))
	return &Config{
		Log_level: log_l,
		Bind:      b,
		DB:        repository.New_config(),
	}
}
