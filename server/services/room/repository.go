package room

//go:generate mockery --name=repository --structname=RepositoryMock --inpackage --filename=repository_mock.go

import (
	"github.com/alissonsz/jun2-ish_goapi/server/models"
	"github.com/jmoiron/sqlx"
)

type repository interface {
	Create(room *models.Room) (*models.Room, error)
	GetById(id int64) (*models.Room, error)
}

type repositoryDB struct {
	db *sqlx.DB
}

func NewRepository(dbConn *sqlx.DB) *repositoryDB {
	return &repositoryDB{db: dbConn}
}

func (r *repositoryDB) Create(room *models.Room) (*models.Room, error) {

	var createdRoom models.Room
	err := r.db.QueryRowx(insertQuery, room.Name, room.VideoUrl, room.Playing, room.Progress).StructScan(&createdRoom)

	if err != nil {
		return nil, err
	}

	return &createdRoom, err
}

func (r *repositoryDB) GetById(id int64) (*models.Room, error) {
	room := models.Room{}
	err := r.db.Get(&room, getByIdQuery, id)
	if err != nil {
		return nil, err
	}

	return &room, err
}
