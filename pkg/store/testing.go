package store

import (
	"fmt"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func Test_store(t *testing.T, db_url string) (*Store, func(...string)) {
	t.Helper()
	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err.Error())
	}
	c := New_config()
	c.DB_URL = db_url
	s := New(c)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}

}
