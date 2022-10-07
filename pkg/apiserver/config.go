package apiserver

import (
	"chat/pkg/store"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	bind      string
	log_level string
	Store     *store.Config
}

func New_config() *Config {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err.Error())
	}

	log_l := os.Getenv("LOGLEVEL")
	b := fmt.Sprintf("%s:%s", os.Getenv("SERVERHOST"), os.Getenv("SERVERPORT"))
	return &Config{
		log_level: log_l,
		bind:      b,
		Store:     store.New_config(),
	}
}
