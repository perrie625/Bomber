package models

import (
	"net"
	"log"
	"Bomber/network"
	"github.com/golang/protobuf/proto"
)

type Session struct {
	Conn       *net.TCPConn
	RemoteAddr string
	MsgProxy   *network.MsgProxy
	Room       Room
}

func (s *Session) Close() {
	s.ExitRoom()
	err := s.Conn.Close()
	if err == nil {
		log.Println(s.RemoteAddr, " disconnect")
	}
}



func (s *Session) ReadProtoMessage() (*network.RawMessage, error) {
	return s.MsgProxy.ReadMsgPacket()
}


func (s *Session) SendProtoMessage(msgId int32, data proto.Message) {
	pkg, _ := s.MsgProxy.MessageToPackage(msgId, data)
	s.Conn.Write(pkg.Bytes())
}

func (s *Session) EntryRoom(room Room) {
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
		MsgProxy:   network.NewMsgProxy(conn),
	}
}