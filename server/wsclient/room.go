package wsclient

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alissonsz/jun2-ish_goapi/server/models"
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

type newUserJoinedMessage struct {
	Type string `json:"type"`
	Data struct {
		Nickname string `json:"progress"`
	} `json:"data"`
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

			roomData, err := wsRoom.roomService.GetById(wsRoom.Id)
			if err != nil {
				fmt.Printf("error: %v", err)
				return
			}

			wsRoom.emitNewUserJoined(client)
			client.emitVideoState(roomData)

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

func (wsRoom *wsRoom) emitNewUserJoined(client *Client) {
	message := &newUserJoinedMessage{Type: "newUserJoined"}
	message.Data.Nickname = client.Nickname

	jsonMessage, err := json.Marshal(message)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	wsRoom.broadcast <- []byte(jsonMessage)
}

func (wsRoom *wsRoom) emitVideoPlayingChanged(message *VideoPlayingChangedMessage) {
	type videoPlayingChangedMessage struct {
		Type string `json:"type"`
		Data struct {
			Progress float64 `json:"progress"`
			Playing  bool    `json:"playing"`
		} `json:"data"`
	}

	jsonMessage, err := json.Marshal(&videoPlayingChangedMessage{Type: "videoPlayingChanged", Data: *message})
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	wsRoom.broadcast <- []byte(jsonMessage)
}

func (wsRoom *wsRoom) emitVideoChanged(message *ChangeVideoMessage) {
	type videoChangedMessage struct {
		Type string `json:"type"`
		Data struct {
			VideoUrl string `json:"videoUrl"`
		} `json:"data"`
	}

	jsonMessage, err := json.Marshal(&videoChangedMessage{Type: "videoChanged", Data: *message})
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	wsRoom.broadcast <- []byte(jsonMessage)
}

func (wsRoom *wsRoom) emitAddedToPlaylist(playlistItem *models.PlaylistItem) {
	type messageData struct {
		Id       int64  `json:"id"`
		Name     string `json:"name"`
		VideoUrl string `json:"videoUrl"`
	}

	type addedToPlaylistMessage struct {
		Type string      `json:"type"`
		Data messageData `json:"data"`
	}

	data := &messageData{
		Id:       playlistItem.Id,
		Name:     playlistItem.Name,
		VideoUrl: playlistItem.VideoUrl,
	}

	jsonMessage, err := json.Marshal(&addedToPlaylistMessage{Type: "addedToPlaylist", Data: *data})
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	wsRoom.broadcast <- []byte(jsonMessage)
}
