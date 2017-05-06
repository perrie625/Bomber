package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"Bomber/protodata"
	"Bomber/network"
)

func main(){
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal("failed")
	}

	go func(conn *net.TCPConn) {
		defer conn.Close()
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				return
			}
			log.Println(message)
		}
	}(conn)

	for {
		var msg string
		fmt.Scanln(&msg)
		if msg == "quit" {
			break
		}
		if msg == ""{
			continue
		}
		pbMessage := &protodata.SayMessage{Words:msg}
		network.WriteMessage(conn, 1, pbMessage)
	}

}