package room

import (
	"fmt"
	"testing"

	"github.com/alissonsz/jun2-ish_goapi/server/database"
	"github.com/alissonsz/jun2-ish_goapi/server/models"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
	db *sqlx.DB
}

func (s *RepositoryTestSuite) SetupTest() {
	dbCfg := database.ClientConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "mysecretpassword",
		Dbname:   "jun2-ish_test_db"}

	testDb, err := database.Setup(dbCfg)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	s.db = testDb
}

func (s *RepositoryTestSuite) TestRoomRepository() {
	s.Run("Create", func() {
		s.T().Parallel()
		r := NewRepository(s.db)

		videoUrl := "https://www.youtube.com/watch?test"
		room := &models.Room{
			Name:     "test",
			VideoUrl: &videoUrl,
			Playing:  true,
			Progress: 5,
		}

		room, err := r.Create(room)
		s.NoError(err)

		insertedRoom, err := r.GetById(room.Id)
		s.NoError(err)

		s.Equal(room.Id, insertedRoom.Id)
		s.Equal(room.Name, insertedRoom.Name)
		s.Equal(room.VideoUrl, insertedRoom.VideoUrl)
		s.Equal(room.Playing, insertedRoom.Playing)
		s.Equal(room.Progress, insertedRoom.Progress)
	})

	s.Run("Get by id", func() {
		s.T().Parallel()
		r := NewRepository(s.db)

		room, err := s.buildRoom()
		s.NoError(err)

		insertedRoom, err := r.GetById(room.Id)
		s.NoError(err)

		s.Equal(room.Id, insertedRoom.Id)
		s.Equal(room.Name, insertedRoom.Name)
		s.Equal(room.VideoUrl, insertedRoom.VideoUrl)
		s.Equal(room.Playing, insertedRoom.Playing)
		s.Equal(room.Progress, insertedRoom.Progress)
	})

	s.Run("Update", func() {
		s.T().Parallel()
		r := NewRepository(s.db)

		room, err := s.buildRoom()
		s.NoError(err)

		updatedVideoUrl := "youtube.com/updated"

		room.Name = "updated name"
		room.VideoUrl = &updatedVideoUrl
		room.Playing = true

		_, err = r.Update(room)
		s.NoError(err)

		updatedRoomRetrieved, err := r.GetById(room.Id)
		s.NoError(err)

		s.Equal(room.Id, updatedRoomRetrieved.Id)
		s.Equal(room.Name, updatedRoomRetrieved.Name)
		s.Equal(room.VideoUrl, updatedRoomRetrieved.VideoUrl)
		s.Equal(room.Playing, updatedRoomRetrieved.Playing)
		s.Equal(room.Progress, updatedRoomRetrieved.Progress)
	})

	s.Run("Create chat message", func() {
		s.T().Parallel()
		r := NewRepository(s.db)

		room, err := s.buildRoom()
		s.NoError(err)

		chatMessage := &models.ChatMessage{
			Author:  "test",
			Content: "test message",
		}

		createdMessage, err := r.CreateChatMessage(room.Id, chatMessage)
		s.NoError(err)

		s.Equal(chatMessage.Author, createdMessage.Author)
		s.Equal(chatMessage.Content, createdMessage.Content)
		s.Equal(room.Id, createdMessage.RoomId)
	})

	s.Run("Create playlist item", func() {
		s.T().Parallel()
		r := NewRepository(s.db)

		room, err := s.buildRoom()
		s.NoError(err)

		playlistItem := &models.PlaylistItem{
			VideoUrl: "test",
			Name:     "test",
			RoomId:   room.Id,
		}

		createdItem, err := r.CreatePlaylistItem(room.Id, playlistItem)
		s.NoError(err)

		s.Equal(playlistItem.VideoUrl, createdItem.VideoUrl)
		s.Equal(playlistItem.Name, createdItem.Name)
		s.Equal(room.Id, createdItem.RoomId)
	})

	s.Run("Delete Playlist Item", func() {
		s.T().Parallel()
		r := NewRepository(s.db)

		room, err := s.buildRoom()
		s.NoError(err)

		playlistItem := &models.PlaylistItem{
			VideoUrl: "test",
			Name:     "test",
			RoomId:   room.Id,
		}

		_, err = r.CreatePlaylistItem(room.Id, playlistItem)
		s.NoError(err)

		createdItem, err := r.CreatePlaylistItem(room.Id, playlistItem)
		s.NoError(err)

		deletedItem, err := r.DeletePlaylistItem(createdItem.Id)
		s.NoError(err)

		s.NotNil(deletedItem.DeletedAt)

		playlistItems, err := r.GetPlaylistItems(room.Id)
		s.NoError(err)

		s.NotContains(playlistItems, deletedItem)
	})
}

func TestSuit(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (s *RepositoryTestSuite) buildRoom() (*models.Room, error) {
	r := NewRepository(s.db)

	videoUrl := "testing/room"
	roomToCreate := &models.Room{
		Name:     "test",
		VideoUrl: &videoUrl,
	}

	room, err := r.Create(roomToCreate)
	return room, err
}
