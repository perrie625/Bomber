package main

import (
	"net"
	"log"
	"os"
	Gate "Bomber/gate"
	"Bomber/room"
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
		agent := Gate.NewAgent(conn)
		agent.EntryRoom(room.MainRoom)

		go agent.Run()

		// destroy room test
		//if count == 1 {
		//	room.MainRoom.Destroy()
		//	break
		//}
	}



}