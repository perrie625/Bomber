package gate

import (
	"net"
	"log"
	"bufio"
	"Bomber/network"
)

type room interface {
	AddAgent(*Agent)
	RemoveAgent(*Agent)
	BroadCast(string)
}

type Agent struct {
	Conn       *net.TCPConn
	RemoteAddr string
	msgParser  *network.MsgParser
	Room       room
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
		go agent.Room.BroadCast(msg)
	}
}


func (agent *Agent) EntryRoom(room room) {
	agent.Room = room
	room.AddAgent(agent)
}

func (agent *Agent) ExitRoom() {
	if agent.Room != nil{
		agent.Room.RemoveAgent(agent)
		agent.Room = nil
	}

}

func NewAgent(conn *net.TCPConn) *Agent {
	remoteAddr := conn.RemoteAddr().String()
	log.Println(remoteAddr, " connected.")
	return &Agent{
		Conn: conn,
		RemoteAddr: remoteAddr,
		msgParser: network.NewMsgParser(conn),
	}
}