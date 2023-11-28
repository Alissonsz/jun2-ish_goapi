package main

import (
	"fmt"
	"os"

	"github.com/alissonsz/jun2-ish_goapi/server"
	"github.com/alissonsz/jun2-ish_goapi/server/database"
	"github.com/alissonsz/jun2-ish_goapi/server/services/room"
)

func main() {
	dbCfg := database.ClientConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     5432,
		User:     "postgres",
		Password: "mysecretpassword",
		Dbname:   "jun2-ish_db"}
	dbConn, err := database.Setup(dbCfg)

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
