package store

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	DB_URL string
}

func New_config() *Config {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err.Error())
	}

	conn_str := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBNAME"), os.Getenv("PASSWORD"), os.Getenv("SSLMODE"))

	return &Config{DB_URL: conn_str}
}
