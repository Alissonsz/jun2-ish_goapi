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
}

type DataMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type WSClient struct {
	clients  []*Client
	register chan *Client
}

func NewWSClient() *WSClient {
	wsClient := &WSClient{clients: []*Client{}, register: make(chan *Client)}
	go wsClient.registerChannels()
	return wsClient
}

func (c *WSClient) UpgradeAndRegister(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newClient := &Client{Id: uuid.NewString(), Conn: conn}
	c.register <- newClient

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("error: %v", err)
			}
			fmt.Print("Client disconnected \n")
			conn.Close()
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
}

func (c *WSClient) addClient(client *Client) {
	c.clients = append(c.clients, client)
}

func (c *WSClient) notifyNewUser(client *Client) {
	for _, roomClient := range c.clients {
		if roomClient.Id != client.Id {
			roomClient.sendMessage(fmt.Sprintf("New user joined the chat! Id: %s", client.Id))
		}
	}
}

func (c *Client) sendMessage(message string) {
	writer, err := c.Conn.NextWriter(websocket.TextMessage)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	writer.Write([]byte(message))
	writer.Close()
}

func (c *WSClient) registerChannels() {
	for client := range c.register {
		c.addClient(client)
		client.sendMessage("Welcome to the chat!")
		c.notifyNewUser(client)
	}
}
