package store_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var (
	db_URL string
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err.Error())
	}
	db_URL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"), os.Getenv("DBNAME"), os.Getenv("PASSWORD"), os.Getenv("SSLMODE")) //Сделать считывание параметров из env

	os.Exit(m.Run())
}
