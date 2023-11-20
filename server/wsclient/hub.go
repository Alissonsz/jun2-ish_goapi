package wsclient

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alissonsz/jun2-ish_goapi/server/services/room"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Hub struct {
	roomService    room.Service
	rooms          []*wsRoom
	pendingClients []*Client
	register       chan *Client
}

type joinRoom struct {
	RoomId   int64  `json:"roomId"`
	Nickname string `json:"nickname"`
}

func NewWsHub(roomService room.Service) *Hub {
	return &Hub{roomService: roomService, rooms: []*wsRoom{}}
}

func (h *Hub) UpgradeAndRegister(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newClient := &Client{Id: uuid.NewString(), Conn: conn, send: make(chan []byte)}

	for {
		_, message, err := newClient.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("error: %v", err)
			}

			break
		}

		parsedMessage := DataMessage{}
		json.Unmarshal(message, &parsedMessage)

		if parsedMessage.Type == "joinRoom" {
			joinRoom := joinRoom{}
			json.Unmarshal(parsedMessage.Data, &joinRoom)

			h.addClientToRoom(newClient, joinRoom.RoomId)
			break
		}
	}
}

func (h *Hub) addClientToRoom(client *Client, roomId int64) {
	room := h.getRoomById(roomId)

	room.register <- client
}

func (h *Hub) getRoomById(roomId int64) *wsRoom {
	for _, room := range h.rooms {
		if room.Id == roomId {
			return room
		}
	}

	newRoom := NewWsRoom(roomId, h.roomService)
	h.rooms = append(h.rooms, newRoom)
	return newRoom
}
