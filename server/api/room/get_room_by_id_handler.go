package room

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/alissonsz/jun2-ish_goapi/server/services/room"
	"github.com/go-chi/chi/v5"
)

type GetRoomByIdHandler struct {
	Service room.Service
}

func (h *GetRoomByIdHandler) Handle(w http.ResponseWriter, r *http.Request) error {
	roomIdString := chi.URLParam(r, "id")

	if len(roomIdString) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing room id"))
		return errors.New("missing room id")
	}

	roomId, err := strconv.Atoi(roomIdString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid room id"))
	}

	room, err := h.Service.GetById(int64(roomId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	roomJson, err := json.Marshal(room)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(roomJson))

	return nil
}
