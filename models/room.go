package models

import (
	"log"
	"sync"

	"github.com/golang/protobuf/proto"
	"Bomber/protodata"
)

var (
	MainRoom = NewRoom()
)


type Room struct {
	roomId     string
	maxLength  int32
	entrance   chan string
	sRWMutex   sync.RWMutex
	sessionMap map[*Session] struct{}
}

func (room *Room) AddSession(session *Session){
	room.sRWMutex.Lock()
	room.sessionMap[session] = struct{}{}
	room.sRWMutex.Unlock()
	log.Println(session.GetAddr(), " entries.")
}

func (room *Room) RemoveSession(s *Session){
	room.sRWMutex.Lock()
	delete(room.sessionMap, s)
	room.sRWMutex.Unlock()
	log.Println(s.GetAddr(), " exits.")
}

func (room *Room) Destroy(){
	for s := range room.sessionMap {
		s.ExitRoom()
	}
}

func (room *Room) BroadCast(message proto.Message){
	room.sRWMutex.RLock()
	for s := range room.sessionMap {
		s.SendProtoMessage(int32(protodata.SaidMessage_ID), message)
	}
	room.sRWMutex.RUnlock()
}


func NewRoom() *Room {
	return &Room{
		roomId:     "main",
		maxLength:  20,
		sessionMap: make(map[*Session] struct{}),
	}
}



