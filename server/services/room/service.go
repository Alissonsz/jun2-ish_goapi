package room

import "github.com/alissonsz/jun2-ish_goapi/server/models"

type Service interface {
	// Create a new room
	Create(room PostPayload) (*models.Room, error)
	// Get a room by its id
	GetById(id int64) (*models.Room, error)
	// Save a chat message
	CreateChatMessage(roomId int64, message *models.ChatMessage) (*models.ChatMessage, error)
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

func (s *service) CreateChatMessage(roomId int64, message *models.ChatMessage) (*models.ChatMessage, error) {
	return s.Repository.CreateChatMessage(roomId, message)
}
