package main

import (
	"net"
	"log"
	"os"
)

func connHandler(connection net.Conn) {
	remote_ip := connection.RemoteAddr()
	log.Println(remote_ip, " connects.")
}


func main()  {

	ln, err := net.Listen("tcp", ":8080")

	defer ln.Close()

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
		go connHandler(conn)
	}

}