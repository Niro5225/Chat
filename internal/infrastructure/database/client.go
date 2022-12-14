package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"chat-app/internal/config"
)

func NewDB(cfg config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.DB)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
