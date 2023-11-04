package room

import (
	"encoding/json"
	"fmt"

	"github.com/alissonsz/jun2-ish_goapi/server/models"
	"github.com/jmoiron/sqlx"
)

type repository interface {
	Create(room *models.Room) error
}

type repositoryDB struct {
	db *sqlx.DB
}

func NewRepository(dbConn *sqlx.DB) *repositoryDB {
	return &repositoryDB{db: dbConn}
}

func (r *repositoryDB) Create(room *models.Room) error {
	rawRoom, _ := json.Marshal(room)

	fmt.Println("Creating room...")
	fmt.Printf("%s \n", rawRoom)
	return nil

	// should call r.db.Exec(insertQuery, room.Name, room.VideoUrl, room.Playing, room.Progress) here
}
