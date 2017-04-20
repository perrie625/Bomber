package room

import (
	"net"
	"Bomber/agent"
)

var (
	MainRoom = NewRoom()
)

type Room struct {
	roomId string
	maxLength int32
	entrance chan string
	agentMap map[string] *net.TCPConn
}


func (room *Room) Entry(agent agent.Agent){

}

func (room *Room) Exit(agent2 agent.Agent){

}

func (room *Room) Destroy(){

}

func (room *Room) broadcast(){

}


func NewRoom() *Room {
	return &Room{}
}



