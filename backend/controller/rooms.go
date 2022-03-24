package controller

import (
	"log"
	"sync"

	"github.com/gofrs/uuid"
	"golang.org/x/net/websocket"
)

type Participant struct {
	Socket *websocket.Conn
}

type Room struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

func (r *Room) CreateRoom() (roomID string, err error) {
	r.Mutex.RLock()
	defer r.Mutex.Unlock()

	Id, err := uuid.NewV4()
	if err != nil {
		log.Println("Couldn't generate room ID: ", err)
		return "", err
	}
	return Id.String(), nil
}
