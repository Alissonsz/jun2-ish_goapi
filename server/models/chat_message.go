package models

type ChatMessage struct {
	Id        int64   `db:"chat_message_id" json:"id"`
	Author    string  `db:"author" json:"author"`
	Content   string  `db:"content" json:"content"`
	RoomId    int64   `db:"room_id" json:"room_id"`
	CreatedAt string  `db:"created_at" json:"created_at"`
	UpdatedAt string  `db:"updated_at" json:"updated_at"`
	DeletedAt *string `db:"deleted_at" json:"deleted_at"`
}
