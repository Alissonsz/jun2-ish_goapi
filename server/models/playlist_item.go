package models

type PlaylistItem struct {
	Id        int64   `db:"playlist_item_id" json:"id"`
	RoomId    int64   `db:"room_id" json:"room_id"`
	Name      string  `db:"name" json:"name"`
	VideoUrl  string  `db:"video_url" json:"video_url"`
	CreatedAt string  `db:"created_at" json:"created_at"`
	UpdatedAt string  `db:"updated_at" json:"updated_at"`
	DeletedAt *string `db:"deleted_at" json:"deleted_at"`
}
