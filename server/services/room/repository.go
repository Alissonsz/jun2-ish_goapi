package room

//go:generate mockery --name=repository --structname=RepositoryMock --inpackage --filename=repository_mock.go

import (
	"encoding/json"
	"fmt"

	"github.com/alissonsz/jun2-ish_goapi/server/models"
	"github.com/jmoiron/sqlx"
)

type repository interface {
	Create(room *models.Room) (*int64, error)
}

type repositoryDB struct {
	db *sqlx.DB
}

func NewRepository(dbConn *sqlx.DB) *repositoryDB {
	return &repositoryDB{db: dbConn}
}

func (r *repositoryDB) Create(room *models.Room) (*int64, error) {
	rawRoom, _ := json.Marshal(room)

	fmt.Println("Creating room...")
	fmt.Printf("%s \n", rawRoom)

	var roomId int64
	err := r.db.QueryRowx(insertQuery, room.Name, room.VideoUrl, room.Playing, room.Progress).Scan(&roomId)
	if err != nil {
		return nil, err
	}

	return &roomId, err
}
