package web

import (
	"encoding/json"
	"net/http"
	"video-chat-app/datastore"
	"video-chat-app/utils"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	pg *datastore.PgAccess
}

func NewServer(d *datastore.PgAccess) *Server {
	return &Server{
		pg: d,
	}
}

func NewRouter(s *Server) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/room", s.CreateRoomHandler).Methods("POST")
	r.HandleFunc("/api/v1/join/{room_id}", s.JoinRoomHandler)

	return r
}

func SendResponse(w http.ResponseWriter, err error, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	statusCode, errorMsg := utils.GetStatusCodeByError(err)

	w.WriteHeader(statusCode)
	var resp utils.GeneralResponse
	if statusCode == utils.ErrorCodeOK {
		resp.Success = true
		resp.StatusCode = statusCode
		resp.ErrMsg = errorMsg
		resp.Data = data
	} else {
		resp.Success = false
		resp.ErrMsg = errorMsg
		resp.StatusCode = statusCode
		resp.Data = data
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		logrus.WithError(err).Error("couldn't send response: ", resp)
	}

}
