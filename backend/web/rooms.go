package web

import (
	"net/http"
	"sync"
	"video-chat-app/utils"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

type Participant struct {
	Socket *websocket.Conn
}

type Room struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

func (s *Server) CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if len(name) == 0 || len(name) > 64 {
		err := utils.ErrBadRequest
		SendResponse(w, err, nil)
		return
	}
	roomId, err := s.pg.CreateRoom(r.Context(), name)
	if err != nil {
		err = utils.ErrInternalServerError
		SendResponse(w, err, nil)
		return
	}
	var resp = struct {
		RoomID string `json:"room_id"`
		Name   string `json:"name"`
	}{
		RoomID: roomId,
		Name:   name,
	}
	err = utils.ErrOK
	SendResponse(w, err, resp)
	logrus.Info("Room is created successfully...")
}
