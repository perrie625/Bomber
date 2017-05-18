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
	var msg protodata.LoginRequest
	if err := proto.Unmarshal(rawMsg.Data, msg); err != nil {
		return
	}
	var resp protodata.LoginResponse
	println(msg.Username)
	println(msg.Password)
	resp.Flag = protodata.FlagNum_eOk
	resp.Desc = "haha"
	session.SendProtoMessage(int32(protodata.LoginResponse_ID), resp)
}


func init() {
	gate.RegisterHandler(
		int32(protodata.LoginRequest_ID),
		HandleLogin,
		reflect.TypeOf(protodata.LoginRequest{}))
}
