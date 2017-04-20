package main

import (
	"net"
	"log"
	"os"
	"bufio"
	agentModel "Bomber/agent"
)

var AgentMap map[string] *agentModel.Agent

func agentHandler(agent *agentModel.Agent) {
	defer agent.Close()
	log.Println(agent.RemoteAddr, " connects.")

	reader := bufio.NewReader(agent.Conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		log.Print(agent.RemoteAddr, "says: ", msg)
	}

}


func main()  {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	ln, err := net.ListenTCP("tcp", tcpAddr)
	defer ln.Close()

	AgentMap = make(map[string] *agentModel.Agent)

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
		agent := agentModel.NewAgent(conn)
		AgentMap[agent.RemoteAddr] = agent

		go agentHandler(agent)
	}

}