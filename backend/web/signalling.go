package web

import (
	"net/http"
	"sync"
	"video-chat-app/utils"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var AllRooms RoomsStorage

type RoomsStorage struct {
	Mutex sync.Mutex
	Rooms map[string][]Participant
}

type Participant struct {
	Conn *websocket.Conn
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

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type broadcastMsg struct {
	Message map[string]interface{}
	RoomID  string
	Client  *websocket.Conn
}

func (r *RoomsStorage) AddParticipant(roomId string, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	participant := Participant{
		Conn: conn,
	}
	r.Rooms[roomId] = append(r.Rooms[roomId], participant)
}

var broadcast = make(chan broadcastMsg)

func (s *Server) JoinRoomHandler(w http.ResponseWriter, r *http.Request) {
	roomId := mux.Vars(r)["room_id"]
	_, err := uuid.FromString(roomId)
	if err != nil {
		eMsg := "RoomId is not compatible"
		logrus.WithError(err).Error(eMsg)
		err = utils.ErrBadRequest
		SendResponse(w, err, nil)
		return
	}

	room, err := s.pg.GetRoom(r.Context(), roomId)
	if err != nil {
		err = utils.ErrInternalServerError
		SendResponse(w, err, nil)
		return
	}
	if room == nil {
		eMsg := "Room with this Id is not found"
		logrus.Error(eMsg)
		err = utils.ErrNotFound
		SendResponse(w, err, nil)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Error("Web Socket Connection Error: ", err)
		return
	}

	AllRooms.AddParticipant(roomId, ws)

	go func() {
		for {
			msg := <-broadcast

			for _, client := range AllRooms.Rooms[msg.RoomID] {
				if client.Conn != msg.Client {
					err := client.Conn.WriteJSON(msg.Message)
					if err != nil {
						logrus.WithError(err)
						client.Conn.Close()
						return
					}
				}
			}

		}
	}()

	for {
		var msg broadcastMsg

		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			logrus.WithError(err)
			return
		}

		msg.Client = ws
		msg.RoomID = roomId

		broadcast <- msg
	}

}
