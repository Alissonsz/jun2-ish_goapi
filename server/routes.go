package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	roomHandler "github.com/alissonsz/jun2-ish_goapi/server/api/room"
	roomService "github.com/alissonsz/jun2-ish_goapi/server/services/room"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
)

var upgrader = websocket.Upgrader{}

type DataMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func ConfigureRoutes(server *Server, dbConn *sqlx.DB, roomService roomService.Service) {
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

	server.router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					fmt.Printf("error: %v", err)
				}
				break
			}

			parsedMessage := DataMessage{}
			json.Unmarshal(message, &parsedMessage)

			if parsedMessage.Type == "message" {
				writer, err := conn.NextWriter(websocket.TextMessage)
				if err != nil {
					fmt.Printf("error: %v", err)
				}

				writer.Write([]byte("You said: " + parsedMessage.Data))
				writer.Close()
			}
		}
	})
}
