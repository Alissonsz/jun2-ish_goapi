package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type DataMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

func ConfigureRoutes(server *Server) {
	server.router.Use(middleware.Logger)
	server.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Heeeeelloo :)!!"))
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
