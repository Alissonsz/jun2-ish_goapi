package model

type PlaylistItem struct {
	Id       int64  `db:"playlist_item_id" json:"id"`
	RoomId   int64  `db:"room_id" json:"room_id"`
	Name     string `db:"name" json:"name"`
	VideoUrl string `db:"video_url" json:"video_url"`
}
