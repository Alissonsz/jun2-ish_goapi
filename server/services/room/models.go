package room

type PostPayload struct {
	Name     string  `json:"name"`
	VideoUrl *string `json:"video_url"`
}
