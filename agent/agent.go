package agent

import "net"

type Agent struct {
	Conn *net.TCPConn
	RemoteAddr string
}

func (agent *Agent) Close() {
	agent.Conn.Close()
}

func NewAgent(conn *net.TCPConn) *Agent {
	return &Agent{
		Conn: conn,
		RemoteAddr: (*conn).RemoteAddr().String(),
	}
}