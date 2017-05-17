package models

import (
	"log"
	"sync"

	"github.com/golang/protobuf/proto"
	"Bomber/protodata"
)

var (
	HallRoom *hall
)


type Room interface {
	RemoveSession(s *Session)
	BroadCast(message proto.Message)
	BroadCastOther(*Session, proto.Message)
	AddSession(*Session)
}

type room struct {
	roomId     string
	maxLength  int32
	entrance   chan string
	sRWMutex   sync.RWMutex
	sessionMap map[*Session] struct{}
}

type hall struct {
	// 大厅
	room
	rooms map[*room]struct{}
}

func (room *room) AddSession(session *Session){
	room.sRWMutex.Lock()
	room.sessionMap[session] = struct{}{}
	room.sRWMutex.Unlock()
	log.Println(session.GetAddr(), " entries.")
}

func (room *room) RemoveSession(s *Session){
	room.sRWMutex.Lock()
	delete(room.sessionMap, s)
	room.sRWMutex.Unlock()
	log.Println(s.GetAddr(), " exits.")
}

func (room *room) Destroy(){
	for s := range room.sessionMap {
		s.ExitRoom()
	}
}

func (room *room) BroadCast(message proto.Message){
	room.sRWMutex.RLock()
	for s := range room.sessionMap {
		s.SendProtoMessage(int32(protodata.SaidMessage_ID), message)
	}
	room.sRWMutex.RUnlock()
}


func (room *room) BroadCastOther(sender *Session, message proto.Message){
	// Broadcast message except the sender
	room.sRWMutex.RLock()
	for s := range room.sessionMap {
		if s != sender {
			s.SendProtoMessage(int32(protodata.SaidMessage_ID), message)
		}
	}
	room.sRWMutex.RUnlock()
}


func NewRoom() *room {
	return &room{
		roomId:     "main",
		maxLength:  20,
		sessionMap: make(map[*Session] struct{}),
	}
}


func init() {
	HallRoom = &hall{
		rooms: make(map[*room] struct{}),
		room: room{
			roomId: "main",
			maxLength: 10000,
			sessionMap: make(map[*Session] struct{}),
		},
	}
}


