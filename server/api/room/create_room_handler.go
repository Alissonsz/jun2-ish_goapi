package room

import (
	"encoding/json"
	"net/http"

	"github.com/alissonsz/jun2-ish_goapi/server/services/room"
)

type CreateRoomHandler struct {
	Service room.Service
}

func (h *CreateRoomHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	room := &room.PostPayload{}

	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	createdRoom, err := h.Service.Create(*room)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	createdRoomJson, err := json.Marshal(createdRoom)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(createdRoomJson))

	return nil
}
