package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ClientConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "jun2-ish_db"
)

func Setup(cfg ClientConfig) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("sslmode=disable host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Dbname)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")
	return db, nil
}
