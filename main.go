package main

import (
	"net"
	"log"
	"os"
	Gate "Bomber/gate"
	"Bomber/models"
	_ "Bomber/handlers"
	"Bomber/tools"
)


func main()  {
	tcpAddr := tools.ServerConfig.GetTcpAddr()
	ln, err := net.ListenTCP("tcp", tcpAddr)
	defer ln.Close()

	if err != nil {
		log.Println("Listen failed.")
		log.Println(err.Error())
		os.Exit(1)
	}
	log.Printf("Listen on %s.", tcpAddr.String())
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			conn.Close()
			continue
		}
		session := models.NewSession(conn)
		// 进入大厅房间
		session.EntryRoom(models.HallRoom)
		go Gate.Agent(session)
	}
}