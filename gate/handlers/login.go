package handlers

import (
	"Bomber/protodata"
	"Bomber/models"
	"github.com/golang/protobuf/proto"
	"Bomber/network"
	"Bomber/gate"
	"reflect"
)

func HandleLogin(session *models.Session, rawMsg *network.RawMessage) {
	// 一个简单的消息处理
	msg := new(protodata.LoginRequest)
	if err := proto.Unmarshal(rawMsg.Data, msg); err != nil {
		return
	}
	resp := new(protodata.LoginResponse)
	println(msg.Username)
	println(msg.Password)
	resp.Flag = protodata.FlagNum_eOk
	resp.Desc = "haha"
	session.Room.BroadCast(resp)
}


func init() {
	gate.RegisterHandler(
		int32(protodata.LoginRequest_ID),
		HandleLogin,
		reflect.TypeOf(protodata.LoginRequest{}))
}
