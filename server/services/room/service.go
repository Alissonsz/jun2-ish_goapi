package room

import "github.com/alissonsz/jun2-ish_goapi/server/models"

type Service interface {
	// Create a new room
	Create(room PostPayload) error
}

type service struct {
	Repository repository
}

func NewService(r repository) *service {
	return &service{
		Repository: r,
	}
}

func (s *service) Create(room PostPayload) error {
	return s.Repository.Create(&models.Room{
		Name:     room.Name,
		VideoUrl: room.VideoUrl,
		Playing:  false,
		Progress: 0,
	})
}
