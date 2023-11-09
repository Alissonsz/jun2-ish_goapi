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
	testDb, err := database.CreateTestDb()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	s.db = testDb
}

func (s *RepositoryTestSuite) TestRoomRepository() {
	r := NewRepository(s.db)

	videoUrl := "https://www.youtube.com/watch?test"
	room := &models.Room{
		Name:     "test",
		VideoUrl: &videoUrl,
		Playing:  true,
		Progress: 5,
	}

	roomId, err := r.Create(room)
	s.NoError(err)

	jsonInsertedRoom := models.Room{}
	err = r.db.Get(&jsonInsertedRoom, "SELECT * FROM room WHERE room_id = $1", *roomId)
	s.NoError(err)

	s.Equal(*roomId, jsonInsertedRoom.Id)
	s.Equal(room.Name, jsonInsertedRoom.Name)
	s.Equal(room.VideoUrl, jsonInsertedRoom.VideoUrl)
	s.Equal(room.Playing, jsonInsertedRoom.Playing)
	s.Equal(room.Progress, jsonInsertedRoom.Progress)
}

func TestSuit(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
