package wsclient

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	Id   string
	Conn *websocket.Conn
	send chan []byte
}

func (c *Client) readPump(room *wsRoom) {
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("error: %v", err)
			}

			room.removeClient(c)
			break
		}

		parsedMessage := DataMessage{}
		json.Unmarshal(message, &parsedMessage)

		room.handleMessage(parsedMessage)
	}
}

func (c *Client) writePump() {
	for message := range c.send {
		writer, err := c.Conn.NextWriter(websocket.TextMessage)
		if err != nil {
			fmt.Printf("error: %v \n", err)
		}

		writer.Write(message)
		writer.Close()
	}
}
