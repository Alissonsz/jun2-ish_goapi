package room

//go:generate mockery --name=repository --structname=RepositoryMock --inpackage --filename=repository_mock.go

import (
	"github.com/alissonsz/jun2-ish_goapi/server/models"
	"github.com/jmoiron/sqlx"
)

type repository interface {
	Create(room *models.Room) (*models.Room, error)
	GetById(id int64) (*models.Room, error)
	CreateChatMessage(roomId int64, message *models.ChatMessage) (*models.ChatMessage, error)
	Update(room *models.Room) (*models.Room, error)
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

	chatMessages, err := r.getChatMessages(id)
	if err != nil {
		return nil, err
	}

	room.Messages = chatMessages

	return &room, err
}

func (r *repositoryDB) CreateChatMessage(roomId int64, message *models.ChatMessage) (*models.ChatMessage, error) {
	var createdMessage models.ChatMessage
	err := r.db.QueryRowx(insertChatMessageQuery, roomId, message.Author, message.Content).StructScan(&createdMessage)

	if err != nil {
		return nil, err
	}

	return &createdMessage, err
}

func (r *repositoryDB) Update(room *models.Room) (*models.Room, error) {
	var updatedRoom models.Room
	err := r.db.QueryRowx(updateQuery, room.Id, room.Name, room.VideoUrl, room.Playing, room.Progress).StructScan(&updatedRoom)

	if err != nil {
		return nil, err
	}

	return &updatedRoom, err
}

func (r *repositoryDB) getChatMessages(roomId int64) ([]models.ChatMessage, error) {
	messages := []models.ChatMessage{}
	err := r.db.Select(&messages, getChatMessagesQuery, roomId)

	if err != nil {
		return nil, err
	}

	return messages, err
}
