package main

import (
	"fmt"

	"github.com/alissonsz/jun2-ish_goapi/server"
	"github.com/alissonsz/jun2-ish_goapi/server/database"
	"github.com/alissonsz/jun2-ish_goapi/server/services/room"
)

func main() {
	dbConn, err := database.Setup()

	if err != nil {
		panic(err)
	}

	roomService := room.NewService(room.NewRepository(dbConn))

	app := server.NewServer(dbConn)
	server.ConfigureRoutes(app, dbConn, roomService)

	if err := app.Run(); err != nil {
		panic(err)
	}

	fmt.Println("Server started!!")
}
