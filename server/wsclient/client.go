package wsclient

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type Client struct {
	Id   string
	Conn *websocket.Conn
	send chan []byte
}

type DataMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type WSClient struct {
	clients   []*Client
	register  chan *Client
	broadcast chan []byte
}

func NewWSClient() *WSClient {
	wsClient := &WSClient{clients: []*Client{}, register: make(chan *Client), broadcast: make(chan []byte)}
	go wsClient.registerChannels()
	return wsClient
}

func (c *WSClient) UpgradeAndRegister(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newClient := &Client{Id: uuid.NewString(), Conn: conn, send: make(chan []byte)}
	c.register <- newClient
}

func (c *WSClient) addClient(client *Client) {
	c.clients = append(c.clients, client)
}

func (c *WSClient) registerChannels() {
	go c.broadcastPump()

	for client := range c.register {
		c.addClient(client)
		go client.readPump(c)
		go client.writePump()
		client.send <- []byte("Welcome to the chat!")
		c.broadcast <- []byte(fmt.Sprintf("User %s joined the chat!", client.Id))
	}
}

func (c *WSClient) broadcastPump() {
	for message := range c.broadcast {
		for _, client := range c.clients {
			client.send <- message
		}
	}
}

func (c *Client) readPump(room *WSClient) {
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("error: %v", err)
			}
			fmt.Print("Client disconnected \n")
			c.Conn.Close()
			break
		}

		parsedMessage := DataMessage{}
		json.Unmarshal(message, &parsedMessage)

		if parsedMessage.Type == "message" {
			room.broadcast <- []byte(fmt.Sprintf("User %s: %s", c.Id, parsedMessage.Data))
		}
	}
}

func (c *Client) writePump() {
	for message := range c.send {
		writer, err := c.Conn.NextWriter(websocket.TextMessage)
		if err != nil {
			fmt.Printf("error: %v", err)
		}

		writer.Write(message)
		writer.Close()
	}
}
