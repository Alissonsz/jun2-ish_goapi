package wsclient

import (
	"encoding/json"
	"fmt"

	"github.com/alissonsz/jun2-ish_goapi/server/models"
)

type DataMessage struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type UserMessage struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}

func (c *WSClient) handleMessage(message DataMessage) {
	switch message.Type {
	case "message":
		userMessage := UserMessage{}
		err := json.Unmarshal(message.Data, &userMessage)

		if err != nil {
			fmt.Printf("error: %v", err)
		} else {
			c.handleChatMessage(userMessage)
		}
	}
}

func (c *WSClient) handleChatMessage(message UserMessage) {
	chatMessage := &models.ChatMessage{
		Author:  message.Author,
		Content: message.Content,
	}

	createdMessage, err := c.roomService.CreateChatMessage(1, chatMessage)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	c.broadcast <- []byte(fmt.Sprintf("%s: %s", createdMessage.Author, createdMessage.Content))
}
