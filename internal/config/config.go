package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Bind      string
	Log_level string
	DB        string
}

func New_config() *Config {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err.Error())
	}

	log_l := os.Getenv("LOGLEVEL")
	b := fmt.Sprintf("%s:%s", os.Getenv("SERVERHOST"), os.Getenv("SERVERPORT"))
	conn_str := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBNAME"), os.Getenv("PASSWORD"), os.Getenv("SSLMODE"))
	return &Config{
		Log_level: log_l,
		Bind:      b,
		DB:        conn_str,
	}
}
