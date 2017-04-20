package agent

import (
	"net"
	"log"
)

type Agent struct {
	Conn *net.TCPConn
	RemoteAddr string
}

func (agent *Agent) Close() {
	log.Println(agent.RemoteAddr, " disconnect")
	agent.Conn.Close()
}

func NewAgent(conn *net.TCPConn) *Agent {
	return &Agent{
		Conn: conn,
		RemoteAddr: (*conn).RemoteAddr().String(),
	}
}