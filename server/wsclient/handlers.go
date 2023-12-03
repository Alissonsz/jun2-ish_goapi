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

type ChatMessage struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}

type VideoPlayingChangedMessage struct {
	Progress float64 `json:"progress"`
	Playing  bool    `json:"playing"`
}

type ChangeVideoMessage struct {
	VideoUrl string `json:"videoUrl"`
}

type VideoSeekedMessage struct {
	SeekTo float64 `json:"seekTo"`
}

type addVideoToPlaylistMessage struct {
	VideoUrl string `json:"videoUrl"`
	Name     string `json:"name"`
}

func (wsRoom *wsRoom) handleMessage(message DataMessage) {
	switch message.Type {
	case "newMessage":
		userMessage := ChatMessage{}
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
	case "videoPlayingChanged":
		videoState := VideoPlayingChangedMessage{}
		err := json.Unmarshal(message.Data, &videoState)

		if err != nil {
			fmt.Printf("error: %v", err)
		} else {
			room, err := wsRoom.roomService.GetById(wsRoom.Id)
			if err != nil {
				fmt.Printf("error: %v", err)
				return
			}

			room.Playing = videoState.Playing
			room.Progress = videoState.Progress

			room, err = wsRoom.roomService.Update(room)
			if err != nil {
				fmt.Printf("error: %v", err)
				return
			}

			wsRoom.emitVideoPlayingChanged(&VideoPlayingChangedMessage{Progress: room.Progress, Playing: room.Playing})
		}
	case "changeVideo":
		changeVideoMessage := ChangeVideoMessage{}
		err := json.Unmarshal(message.Data, &changeVideoMessage)

		if err != nil {
			fmt.Printf("error: %v", err)
		} else {
			room, err := wsRoom.roomService.GetById(wsRoom.Id)
			if err != nil {
				fmt.Printf("error: %v", err)
				return
			}

			room.VideoUrl = &changeVideoMessage.VideoUrl
			room.Progress = 0

			room, err = wsRoom.roomService.Update(room)
			if err != nil {
				fmt.Printf("error: %v", err)
				return
			}

			wsRoom.emitVideoChanged(&ChangeVideoMessage{VideoUrl: *room.VideoUrl})
		}
	case "videoSeeked":
		videoSeekedMessage := VideoSeekedMessage{}
		err := json.Unmarshal(message.Data, &videoSeekedMessage)

		if err != nil {
			fmt.Printf("error: %v", err)
		} else {
			room, err := wsRoom.roomService.GetById(wsRoom.Id)
			if err != nil {
				fmt.Printf("error: %v", err)
				return
			}

			room.Progress = videoSeekedMessage.SeekTo

			_, err = wsRoom.roomService.Update(room)
			if err != nil {
				fmt.Printf("error: %v", err)
				return
			}

			jsonMessage, err := json.Marshal(message)
			if err != nil {
				fmt.Printf("error: %v", err)
			}

			wsRoom.broadcast <- []byte(jsonMessage)
		}
	case "addVideoToPlaylist":
		addVideoToPlaylistMessage := addVideoToPlaylistMessage{}

		err := json.Unmarshal(message.Data, &addVideoToPlaylistMessage)
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}

		playlistItem, err := wsRoom.roomService.CreatePlaylistItem(wsRoom.Id, &models.PlaylistItem{Name: addVideoToPlaylistMessage.Name, VideoUrl: addVideoToPlaylistMessage.VideoUrl})
		if err != nil {
			fmt.Printf("error: %v", err)
			return
		}

		wsRoom.emitAddedToPlaylist(playlistItem)
	default:
		fmt.Println(message)
	}
}

func (wsRoom *wsRoom) handleChatMessage(message ChatMessage) error {
	chatMessage := &models.ChatMessage{
		Author:  message.Author,
		Content: message.Content,
	}

	_, err := wsRoom.roomService.CreateChatMessage(wsRoom.Id, chatMessage)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err
	}

	return nil
}
