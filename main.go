package main

import (
	"fmt"

	"github.com/alissonsz/jun2-ish_goapi/server"
	"github.com/alissonsz/jun2-ish_goapi/server/database"
)

func main() {
	dbConn, err := database.Setup()

	if err != nil {
		panic(err)
	}

	app := server.NewServer(dbConn)
	server.ConfigureRoutes(app)

	if err := app.Run(); err != nil {
		panic(err)
	}

	fmt.Println("Server started!!")
}
