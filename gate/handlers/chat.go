package handlers

import (
	"Bomber/protodata"
	"Bomber/gate"
)

func handleChat(agent *gate.Agent, msg *protodata.SayMessage) {
	// 一个简单的消息处理
	agent.Room.BroadCast(msg.Words)
}
