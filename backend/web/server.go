package web

import (
	"encoding/json"
	"net/http"
	"video-chat-app/utils"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct{}

func NewRouter(s *Server) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/room", s.CreateRoomHandler).Methods("POST")
	r.HandleFunc("/api/v1/join/{room_id}", s.JoinRoomHandler)
	r.HandleFunc("/api/v1/invite", s.SendInvitationHandler).Methods("POST")

	return r
}

func SendResponse(w http.ResponseWriter, err error, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	statusCode, errorMsg := utils.GetStatusCodeByError(err)

	w.WriteHeader(statusCode)

	if statusCode == utils.ErrorCodeOK {
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			logrus.WithError(err).Error("couldn't send response: ", data)
		}
	} else {
		err = json.NewEncoder(w).Encode(errorMsg)
		if err != nil {
			logrus.WithError(err).Error("couldn't send response: ", err)
		}
	}
}
