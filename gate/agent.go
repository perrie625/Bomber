package gate

import (
	"Bomber/protodata"
	"github.com/golang/protobuf/proto"
	"Bomber/gate/handlers"
	"log"
	"Bomber/models"
)


func Agent (session *models.Session){
	defer func(){
		session.Close()
	}()
	for {
		// 待完善
		// 只是单纯实现了proto接收，然后广播字符串
		msgId, err := session.MsgParser.GetMsgId()
		if err != nil {
			log.Println(err.Error())
			return
		}
		msgData, err := session.MsgParser.GetMsgData()
		if err != nil {
			return
		}
		msg := new(protodata.SayMessage)
		_ = proto.Unmarshal(msgData, msg)
		if msgId == 1 {
			handlers.HandleChat(session, msg)
		}
	}
}
