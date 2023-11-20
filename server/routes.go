package server

import (
	"fmt"
	"net/http"

	roomHandler "github.com/alissonsz/jun2-ish_goapi/server/api/room"
	roomService "github.com/alissonsz/jun2-ish_goapi/server/services/room"
	"github.com/alissonsz/jun2-ish_goapi/server/wsclient"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
)

func ConfigureRoutes(server *Server, dbConn *sqlx.DB, roomService roomService.Service) {
	wsClient := wsclient.NewWsHub(roomService)

	server.router.Use(middleware.Logger)
	server.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Heeeeelloo :)!!"))
	})

	server.router.Post("/room", func(w http.ResponseWriter, r *http.Request) {
		handler := &roomHandler.CreateRoomHandler{Service: roomService}
		err := handler.Handle(w, r)

		if err != nil {
			fmt.Printf("error: %v", err)
		}
	})

	server.router.Get("/room/{id}", func(w http.ResponseWriter, r *http.Request) {
		handler := &roomHandler.GetRoomByIdHandler{Service: roomService}
		err := handler.Handle(w, r)

		if err != nil {
			fmt.Printf("error: %v", err)
		}
	})

	server.router.HandleFunc("/ws", wsClient.UpgradeAndRegister)
}
