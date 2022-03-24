package web

import (
	"video-chat-app/datastore"

	"github.com/gorilla/mux"
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

	return r
}
