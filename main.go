package main

import (
	"net"
	"log"
	"os"
	"bufio"
)

var AgentMap map[string] *Agent

type Agent struct {
	conn *net.TCPConn
	remoteAddr string
}

func (agent *Agent) Close() {
	(*agent.conn).Close()
}

func NewAgent(conn *net.TCPConn) *Agent {
	return &Agent{
		conn: conn,
		remoteAddr: (*conn).RemoteAddr().String(),
	}
}

func agentHandler(agent *Agent) {
	defer agent.Close()
	log.Println(agent.remoteAddr, " connects.")

	reader := bufio.NewReader(*agent.conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		log.Print(agent.remoteAddr, "says: ", msg)
	}

}


func main()  {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	ln, err := net.ListenTCP("tcp", tcpAddr)
	defer ln.Close()

	AgentMap = make(map[string] *Agent)

	if err != nil {
		log.Println("Listen failed.")
		os.Exit(1)
	}
	log.Println("Listen on 8080.")
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			conn.Close()
			continue
		}
		agent := NewAgent(conn)
		AgentMap[agent.remoteAddr] = agent

		go agentHandler(agent)
	}

}