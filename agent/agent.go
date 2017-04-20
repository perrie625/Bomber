package agent

import (
	"net"
	"log"
	"bufio"
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
	agent.ExitRoom()
	agent.Conn.Close()
}


func (agent *Agent) Run(){
	defer func(){
		agent.Close()
	}()
	log.Println(agent.RemoteAddr, " connected.")

	reader := bufio.NewReader(agent.Conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		go agent.room.BroadCast(msg)
	}
}


func (agent *Agent) EntryRoom(room room) {
	agent.room = room
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