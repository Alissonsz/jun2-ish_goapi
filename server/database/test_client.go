package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	testHost     = "localhost"
	testPort     = 5432
	testUser     = "postgres"
	testPassword = "mysecretpassword"
	testBbname   = "jun2-ish_test_db"
)

func CreateTestDb() (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("sslmode=disable host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		testHost, testPort, testUser, testPassword, testBbname)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")
	return db, nil
}
