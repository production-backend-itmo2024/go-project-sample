package db_postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Dbname   string
	Password string
	SSL      string
}

func NewPostrgresDB(cfg Config) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Dbname,
		cfg.Password,
		cfg.SSL,
	)

	db, err := sqlx.Open("postgres", dataSourceName)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
