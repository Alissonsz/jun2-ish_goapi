package room

import "github.com/alissonsz/jun2-ish_goapi/server/models"

type Service interface {
	// Create a new room
	Create(room PostPayload) (*models.Room, error)
	// Get a room by its id
	GetById(id int64) (*models.Room, error)
	// Update a room
	Update(room *models.Room) (*models.Room, error)
	// Save a chat message
	CreateChatMessage(roomId int64, message *models.ChatMessage) (*models.ChatMessage, error)
	// Save a playlist item
	CreatePlaylistItem(roomId int64, item *models.PlaylistItem) (*models.PlaylistItem, error)
	// Delete a playlist item
	DeletePlaylistItem(itemId int64) (*models.PlaylistItem, error)
	// Get all playlist items
	GetPlaylistItems(roomId int64) ([]models.PlaylistItem, error)
}

type service struct {
	Repository repository
}

func NewService(r repository) *service {
	return &service{
		Repository: r,
	}
}

func (s *service) Create(room PostPayload) (*models.Room, error) {
	return s.Repository.Create(&models.Room{
		Name:     room.Name,
		VideoUrl: room.VideoUrl,
		Playing:  false,
		Progress: 0,
	})
}

func (s *service) GetById(id int64) (*models.Room, error) {
	return s.Repository.GetById(id)
}

func (s *service) Update(room *models.Room) (*models.Room, error) {
	return s.Repository.Update(room)
}

func (s *service) CreateChatMessage(roomId int64, message *models.ChatMessage) (*models.ChatMessage, error) {
	return s.Repository.CreateChatMessage(roomId, message)
}

func (s *service) CreatePlaylistItem(roomId int64, item *models.PlaylistItem) (*models.PlaylistItem, error) {
	return s.Repository.CreatePlaylistItem(roomId, item)
}

func (s *service) DeletePlaylistItem(itemId int64) (*models.PlaylistItem, error) {
	return s.Repository.DeletePlaylistItem(itemId)
}

func (s *service) GetPlaylistItems(roomId int64) ([]models.PlaylistItem, error) {
	return s.Repository.GetPlaylistItems(roomId)
}
