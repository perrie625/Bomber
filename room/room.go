package room

import (
	"Bomber/agent"
	"log"
)

var (
	MainRoom = NewRoom()
)

type Room struct {
	roomId string
	maxLength int32
	entrance chan string
	agentMap map[*agent.Agent] struct{}
}

func (room *Room) AddAgent(agent *agent.Agent){
	log.Println(agent.RemoteAddr, " entries.")
	room.agentMap[agent] = struct{}{}
}

func (room *Room) RemoveAgent(agent *agent.Agent){
	log.Println(agent.RemoteAddr, " exits.")
	delete(room.agentMap, agent)
}

func (room *Room) Destroy(){

}

func (room *Room) BroadCast(message string){
	b := []byte(message)
	for a := range room.agentMap {
		a.Conn.Write(b)
	}
}


func NewRoom() *Room {
	return &Room{
		roomId: "main",
		maxLength: 20,
		agentMap: make(map[*agent.Agent] struct{}),
	}
}



