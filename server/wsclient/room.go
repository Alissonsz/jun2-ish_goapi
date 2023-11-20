package wsclient

import (
	"fmt"
	"net/http"

	"github.com/alissonsz/jun2-ish_goapi/server/services/room"
	"github.com/google/uuid"
)

type wsRoom struct {
	Id          int64
	roomService room.Service
	clients     []*Client
	register    chan *Client
	broadcast   chan []byte
}

func NewWsRoom(id int64, roomService room.Service) *wsRoom {
	wsRoom := &wsRoom{
		Id:          id,
		roomService: roomService,
		clients:     []*Client{},
		register:    make(chan *Client),
		broadcast:   make(chan []byte),
	}

	wsRoom.registerChannels()
	return wsRoom
}

func (wsRoom *wsRoom) UpgradeAndRegister(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newClient := &Client{Id: uuid.NewString(), Conn: conn, send: make(chan []byte)}
	wsRoom.register <- newClient
}

func (wsRoom *wsRoom) addClient(client *Client) {
	wsRoom.clients = append(wsRoom.clients, client)
}

func (wsRoom *wsRoom) registerChannels() {
	fmt.Printf("Registering channels for room %d \n", wsRoom.Id)
	go wsRoom.broadcastPump()

	go func() {
		for client := range wsRoom.register {
			wsRoom.addClient(client)
			go client.readPump(wsRoom)
			go client.writePump()
			client.send <- []byte("Welcome to the room!")
			wsRoom.broadcast <- []byte(fmt.Sprintf("User %s joined the room!", client.Id))

			fmt.Printf("New client joined room %d, now %d clients \n", wsRoom.Id, len(wsRoom.clients))
		}
	}()
}

func (wsRoom *wsRoom) broadcastPump() {
	for message := range wsRoom.broadcast {
		for _, client := range wsRoom.clients {
			client.send <- message
		}
	}
}

func (wsRoom *wsRoom) removeClient(clientToRemove *Client) {
	for i, client := range wsRoom.clients {
		if client.Id == clientToRemove.Id {
			clientToRemove.Conn.Close()
			close(clientToRemove.send)
			wsRoom.clients = append(wsRoom.clients[:i], wsRoom.clients[i+1:]...)
			wsRoom.broadcast <- []byte(fmt.Sprintf("User %s left the room!", clientToRemove.Id))
			break
		}
	}
}
