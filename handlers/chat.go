package handlers

import (
	"Bomber/protodata"
	"Bomber/models"
	"time"
	"github.com/golang/protobuf/proto"
	"Bomber/network"
	"Bomber/gate"
	"reflect"
)

func HandleChat(session *models.Session, rawMsg *network.RawMessage) {
	// 一个简单的消息处理
	msg := new(protodata.SayMessage)
	if err := proto.Unmarshal(rawMsg.Data, msg); err != nil {
		return
	}
	resp := new(protodata.SaidMessage)
	resp.Name = session.GetAddr()
	now := time.Now()
	resp.Time = now.Format("2006-01-02 15:04:05")
	resp.Words = msg.Words

	pkg, _ := session.MsgProxy.MessageToPackage(int32(protodata.SaidMessage_ID), resp)
	session.Room.BroadCast(pkg.Bytes())
}


func init() {
	gate.RegisterHandler(
		int32(protodata.SayMessage_ID),
		HandleChat,
		reflect.TypeOf(protodata.SayMessage{}))
}
