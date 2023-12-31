package models

type Room struct {
	Id            int64          `db:"room_id" json:"id"`
	Name          string         `db:"name" json:"name"`
	VideoUrl      *string        `db:"video_url" json:"video_url"`
	Messages      []ChatMessage  `json:"messages"`
	Playing       bool           `db:"playing" json:"playing"`
	Progress      float64        `db:"progress" json:"progress"`
	PlaylistItems []PlaylistItem `json:"playlist_items"`
	CreatedAt     string         `db:"created_at" json:"created_at"`
	UpdatedAt     string         `db:"updated_at" json:"updated_at"`
	DeletedAt     *string        `db:"deleted_at" json:"deleted_at"`
}
