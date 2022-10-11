package store

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	conf            *Config
	db              *sqlx.DB
	user_repository *User_repository
}

func New(conf *Config) *Store {
	return &Store{conf: conf}
}

func (s *Store) Open() error {
	db, err := sqlx.Open("postgres", s.conf.DB_URL)

	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

//Метод для исполизования userreposirory из хранилища
func (s *Store) User() *User_repository {
	if s.user_repository != nil {
		return s.user_repository
	}

	s.user_repository = &User_repository{
		store: s,
	}

	return s.user_repository
}
