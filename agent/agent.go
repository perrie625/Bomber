package agent

import (
	"net"
	"log"
	"reflect"
)

type room interface {
	AddAgent(*Agent)
	RemoveAgent(*Agent)
	BroadCast(string)
}

type Agent struct {
	Conn *net.TCPConn
	RemoteAddr string
	room room
}

func (agent *Agent) Close() {
	log.Println(agent.RemoteAddr, " disconnect")
	agent.Conn.Close()
}

func (agent *Agent) EntryRoom(room room) {
	agent.room = room
	log.Println(reflect.TypeOf(room))
	log.Println(reflect.TypeOf(agent.room))
	room.AddAgent(agent)
}

func (agent *Agent) ExitRoom() {
	agent.room.RemoveAgent(agent)
	agent.room = nil
}

func NewAgent(conn *net.TCPConn) *Agent {
	return &Agent{
		Conn: conn,
		RemoteAddr: (*conn).RemoteAddr().String(),
	}
}