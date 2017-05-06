package gate

import (
	"net"
	"log"
	"Bomber/network"
	"Bomber/protodata"
	"github.com/golang/protobuf/proto"
)

type room interface {
	AddAgent(*Agent)
	RemoveAgent(*Agent)
	BroadCast(string)
}

type Agent struct {
	Conn       *net.TCPConn
	RemoteAddr string
	msgParser  *network.MsgParser
	Room       room
}

func (agent *Agent) Close() {
	agent.ExitRoom()
	err := agent.Conn.Close()
	if err == nil {
		log.Println(agent.RemoteAddr, " disconnect")
	}
}


func (agent *Agent) Run(){
	defer func(){
		agent.Close()
	}()
	for {
		// 待完善
		// 只是单纯实现了proto接收，然后广播字符串
		msgId, err := agent.msgParser.GetMsgId()
		log.Println(msgId)
		if err != nil {
			return
		}
		msgData, err := agent.msgParser.GetMsgData()
		if err != nil {
			return
		}
		msg := new(protodata.SayMessage)
		_ = proto.Unmarshal(msgData, msg)
		agent.Room.BroadCast(msg.Words + "\n")
	}
}

func (agent *Agent) ReadMessage() {

}


func (agent *Agent) EntryRoom(room room) {
	agent.Room = room
	room.AddAgent(agent)
}

func (agent *Agent) ExitRoom() {
	if agent.Room != nil{
		agent.Room.RemoveAgent(agent)
		agent.Room = nil
	}

}

func NewAgent(conn *net.TCPConn) *Agent {
	remoteAddr := conn.RemoteAddr().String()
	log.Println(remoteAddr, " connected.")
	return &Agent{
		Conn: conn,
		RemoteAddr: remoteAddr,
		msgParser: network.NewMsgParser(conn),
	}
}