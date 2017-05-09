package handlers

import (
	"Bomber/protodata"
	"Bomber/models"
	"time"
	"github.com/golang/protobuf/proto"
	"Bomber/network"
)

func HandleChat(session *models.Session, rawMsg *network.RawMessage) {
	// 一个简单的消息处理
	msg := new(protodata.SayMessage)
	if err := proto.Unmarshal(rawMsg.Data, msg); err != nil {
		return
	}
	resp := new(protodata.SaidMessage)
	resp.Name = session.RemoteAddr
	now := time.Now()
	resp.Time = now.Format("2006-01-02 15:04:05")
	resp.Words = msg.Words
	session.Room.BroadCast(resp)
}
