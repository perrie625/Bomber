package main

import (
	"net"
	"log"
	"os"
	Gate "Bomber/gate"
	"Bomber/models"
	_ "Bomber/gate/handlers"
)


func main()  {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	ln, err := net.ListenTCP("tcp", tcpAddr)
	defer ln.Close()

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
		session := models.NewSession(conn)
		// 进入大厅房间
		session.EntryRoom(models.MainRoom)
		go Gate.Agent(session)
	}
}