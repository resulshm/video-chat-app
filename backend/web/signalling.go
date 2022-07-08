package web

import (
	"fmt"
	"net/http"
	"sync"
	"video-chat-app/utils"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	mail "github.com/xhit/go-simple-mail/v2"
)

var AllRooms RoomsStorage

type RoomsStorage struct {
	Mutex sync.RWMutex
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
					AllRooms.Mutex.Lock()
					err := client.Conn.WriteJSON(msg.Message)
					if err != nil {
						logrus.WithError(err).Error("Write JSON...")
					}
					AllRooms.Mutex.Unlock()
				}
			}

		}
	}()

	for {
		var msg broadcastMsg

		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			logrus.WithError(err).Error("Read JSON...")
			return
		}

		msg.Client = ws
		msg.RoomID = roomId
		// fmt.Println(msg.Message)
		broadcast <- msg
	}
}

func (s *Server) SendInvitationHandler(w http.ResponseWriter, r *http.Request) {
	var to string
	if utils.IsEmailValid(r.FormValue("to")) {
		to = r.FormValue("to")
	} else {
		eMsg := "Invalid email address"
		logrus.Error(eMsg)
		err := utils.ErrBadRequest
		SendResponse(w, err, nil)
		return
	}

	from := r.FormValue("from")
	if from == "" {
		eMsg := "From is empty"
		logrus.Error(eMsg)
		err := utils.ErrBadRequest
		SendResponse(w, err, nil)
		return
	}

	code := r.FormValue("code")
	_, err := uuid.FromString(code)
	if err != nil {
		eMsg := "Code is not correct"
		logrus.Error(eMsg)
		err := utils.ErrBadRequest
		SendResponse(w, err, nil)
		return
	}

	subject := fmt.Sprintf("Invitation from %s", from)
	content := fmt.Sprintf("<p>Your friend wants to talk to you...<br> Enter the code in our app to join &#128521; <br> <strong>Code:</strong> %s</p>", code)
	go SendEmail(to, subject, content)
	err = utils.ErrOK
	SendResponse(w, err, nil)
}

func SendEmail(to, subject, content string) {
	srv := mail.NewSMTPClient()
	srv.Host = utils.Conf.Host
	srv.Port = utils.Conf.Port
	srv.Username = utils.Conf.Username
	srv.Password = utils.Conf.Password
	srv.Encryption = mail.EncryptionTLS

	smtpClient, err := srv.Connect()
	if err != nil {
		eMsg := "Couldn't connect to SMTP server"
		logrus.WithError(err).Error(eMsg)
		return
	}

	email := mail.NewMSG()
	from := fmt.Sprintf("RS video-chat app <%s>", utils.Conf.Username)
	email.SetFrom(from)
	email.AddTo(to)
	email.SetSubject(subject)
	email.SetBody(mail.TextHTML, content)

	err = email.Send(smtpClient)
	if err != nil {
		eMsg := "Could't send email"
		logrus.WithError(err).Error(eMsg)
		return
	}

	logrus.Info("Email is sended successfully...")
}
