package handlers

import (
	"Bomber/protodata"
	"Bomber/models"
	"time"
)

func HandleChat(session *models.Session, msg *protodata.SayMessage) {
	// 一个简单的消息处理
	resp := new(protodata.SaidMessage)
	resp.Name = session.RemoteAddr
	now := time.Now()
	resp.Time = now.Format("2006-01-02 15:04:05")
	resp.Words = msg.Words
	session.Room.BroadCast(resp)
}
