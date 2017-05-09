package models

import (
	"net"
	"log"
	"Bomber/network"
	"Bomber/protodata"
	"github.com/golang/protobuf/proto"
	"time"
)

type Session struct {
	Conn       *net.TCPConn
	RemoteAddr string
	MsgParser  *network.MsgParser
	Room       *Room
}

func (s *Session) Close() {
	s.ExitRoom()
	err := s.Conn.Close()
	if err == nil {
		log.Println(s.RemoteAddr, " disconnect")
	}
}


func (s *Session) Run(){
	defer func(){
		s.Close()
	}()
	for {
		// 待完善
		// 只是单纯实现了proto接收，然后广播字符串
		_, err := s.MsgParser.GetMsgId()
		if err != nil {
			return
		}
		msgData, err := s.MsgParser.GetMsgData()
		if err != nil {
			return
		}
		msg := new(protodata.SayMessage)
		_ = proto.Unmarshal(msgData, msg)
		resp := new(protodata.SaidMessage)
		resp.Name = s.RemoteAddr
		now := time.Now()
		resp.Time = now.Format("2006-01-02 15:04:05")
		resp.Words = msg.Words
		s.Room.BroadCast(resp)
	}
}

func (s *Session) ReadMessage() {

}


func (s *Session) EntryRoom(room *Room) {
	s.Room = room
	room.AddSession(s)
}

func (s *Session) ExitRoom() {
	if s.Room != nil{
		s.Room.RemoveSession(s)
		s.Room = nil
	}

}

func(s *Session ) GetAddr() string {
	return s.RemoteAddr
}

func(s *Session) GetCon() *net.TCPConn {
	return s.Conn
}

func NewSession(conn *net.TCPConn) *Session {
	remoteAddr := conn.RemoteAddr().String()
	log.Println(remoteAddr, " connected.")
	return &Session{
		Conn:       conn,
		RemoteAddr: remoteAddr,
		MsgParser:  network.NewMsgParser(conn),
	}
}