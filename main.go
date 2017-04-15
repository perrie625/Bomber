package main

import (
	"net"
	"log"
	"os"
)

var AgentMap map[string] *Agent

type Agent struct {
	conn *net.Conn
	remoteAddr string
}

func (agent *Agent) Close() {
	(*agent.conn).Close()
}

func NewAgent(conn *net.Conn) *Agent {
	return &Agent{
		conn: conn,
		remoteAddr: (*conn).RemoteAddr().String(),
	}
}

func agentHandler(agent *Agent) {
	defer agent.Close()
	log.Println(agent.remoteAddr, " connects.")
}


func main()  {

	ln, err := net.Listen("tcp", ":8080")
	defer ln.Close()

	AgentMap = make(map[string] *Agent)

	if err != nil {
		log.Println("Listen failed.")
		os.Exit(1)
	}
	log.Println("Listen on 8080.")
	for {
		conn, err := ln.Accept()
		if err != nil {
			conn.Close()
			continue
		}
		agent := NewAgent(&conn)
		AgentMap[agent.remoteAddr] = agent

		go agentHandler(agent)
	}

}