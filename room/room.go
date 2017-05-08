package room

import (
	agent "Bomber/gate"
	"log"
	"sync"

	"github.com/golang/protobuf/proto"
	"Bomber/network"
)

var (
	MainRoom = NewRoom()
)

type Room struct {
	roomId string
	maxLength int32
	entrance chan string
	agentsRWMutex sync.RWMutex
	agentMap map[*agent.Agent] struct{}
}

func (room *Room) AddAgent(agent *agent.Agent){
	room.agentsRWMutex.Lock()
	room.agentMap[agent] = struct{}{}
	room.agentsRWMutex.Unlock()
	log.Println(agent.RemoteAddr, " entries.")
}

func (room *Room) RemoveAgent(agent *agent.Agent){
	room.agentsRWMutex.Lock()
	delete(room.agentMap, agent)
	room.agentsRWMutex.Unlock()
	log.Println(agent.RemoteAddr, " exits.")
}

func (room *Room) Destroy(){
	for a := range room.agentMap {
		a.Close()
	}
}

func (room *Room) BroadCast(message proto.Message){
	room.agentsRWMutex.RLock()
	for a := range room.agentMap {
		network.WriteMessage(a.Conn, 1, message)
	}
	room.agentsRWMutex.RUnlock()
}


func NewRoom() *Room {
	return &Room{
		roomId: "main",
		maxLength: 20,
		agentMap: make(map[*agent.Agent] struct{}),
	}
}



