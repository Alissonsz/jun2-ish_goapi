package wsclient

import (
	"encoding/json"
	"fmt"

	"github.com/alissonsz/jun2-ish_goapi/server/models"
	"github.com/gorilla/websocket"
)

type Client struct {
	Id       string
	Nickname string
	Conn     *websocket.Conn
	send     chan []byte
}

type videoStateMessage struct {
	Type string `json:"type"`
	Data struct {
		Progress float64 `json:"progress"`
		Playing  bool    `json:"playing"`
	} `json:"data"`
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

func (c *Client) emitVideoState(room *models.Room) {
	videoState := &videoStateMessage{}
	videoState.Type = "videoState"
	videoState.Data.Playing = room.Playing
	videoState.Data.Progress = room.Progress

	jsonMessage, err := json.Marshal(videoState)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	c.send <- []byte(jsonMessage)
}
