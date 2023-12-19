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
	CreatePlaylistItem(roomId int64, item *models.PlaylistItem) (*models.PlaylistItem, error)
	DeletePlaylistItem(itemId int64) (*models.PlaylistItem, error)
	GetPlaylistItems(roomId int64) ([]models.PlaylistItem, error)
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

	playlistItems, err := r.GetPlaylistItems(id)
	if err != nil {
		return nil, err
	}

	room.PlaylistItems = playlistItems

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

func (r *repositoryDB) CreatePlaylistItem(roomId int64, item *models.PlaylistItem) (*models.PlaylistItem, error) {
	var createdItem models.PlaylistItem
	err := r.db.QueryRowx(insertPlaylistItemQuery, roomId, item.VideoUrl, item.Name).StructScan(&createdItem)

	if err != nil {
		return nil, err
	}

	return &createdItem, err
}

func (r *repositoryDB) DeletePlaylistItem(itemId int64) (*models.PlaylistItem, error) {
	var deletedItem models.PlaylistItem
	err := r.db.QueryRowx(deletePlaylistItemQuery, itemId).StructScan(&deletedItem)

	if err != nil {
		return nil, err
	}

	return &deletedItem, err
}

func (r *repositoryDB) GetPlaylistItems(roomId int64) ([]models.PlaylistItem, error) {
	items := []models.PlaylistItem{}
	err := r.db.Select(&items, getPlaylistItemsQuery, roomId)

	if err != nil {
		return nil, err
	}

	return items, err
}

func (r *repositoryDB) getChatMessages(roomId int64) ([]models.ChatMessage, error) {
	messages := []models.ChatMessage{}
	err := r.db.Select(&messages, getChatMessagesQuery, roomId)

	if err != nil {
		return nil, err
	}

	return messages, err
}
