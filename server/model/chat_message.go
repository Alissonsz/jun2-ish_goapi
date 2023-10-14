package model

type ChatMessage struct {
	Id      int64  `db:"message_id" json:"id"`
	Author  string `db:"author" json:"author"`
	Content string `db:"content" json:"content"`
	RoomId  int64  `db:"room_id" json:"room_id"`
}
