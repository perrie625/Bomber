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
	agent.ExitRoom()
	err := agent.Conn.Close()
	if err == nil {
		log.Println(agent.RemoteAddr, " disconnect")
	}
}


func (agent *Agent) Run(){
	defer func(){
		agent.Close()
	}()
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
	if agent.room != nil{
		agent.room.RemoveAgent(agent)
		agent.room = nil
	}

}

func NewAgent(conn *net.TCPConn) *Agent {
	remoteAddr := conn.RemoteAddr().String()
	log.Println(remoteAddr, " connected.")
	return &Agent{
		Conn: conn,
		RemoteAddr: remoteAddr,
	}
}