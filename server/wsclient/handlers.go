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

func (wsRoom *wsRoom) handleMessage(message DataMessage) {
	switch message.Type {
	case "newMessage":
		userMessage := UserMessage{}
		err := json.Unmarshal(message.Data, &userMessage)

		if err != nil {
			fmt.Printf("error: %v", err)
		} else {
			wsRoom.handleChatMessage(userMessage)

			jsonMessage, err := json.Marshal(message)
			if err != nil {
				fmt.Printf("error: %v", err)
			}

			wsRoom.broadcast <- []byte(jsonMessage)
		}
	default:
		fmt.Println(message)
	}
}

func (wsRoom *wsRoom) handleChatMessage(message UserMessage) error {
	chatMessage := &models.ChatMessage{
		Author:  message.Author,
		Content: message.Content,
	}

	createdMessage, err := wsRoom.roomService.CreateChatMessage(wsRoom.Id, chatMessage)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}

	wsRoom.broadcast <- []byte(fmt.Sprintf("%s: %s", createdMessage.Author, createdMessage.Content))
	return nil
}
