package gate

import (
	"Bomber/gate/handlers"
	"Bomber/models"
)


func Agent (session *models.Session){
	defer func(){
		session.Close()
	}()
	for {
		// 待完善
		// 只是单纯实现了proto接收，然后广播字符串
		rawMsg, err := session.MsgParser.ReadMsgPacket()
		if err != nil {
			return
		}
		if rawMsg.Id == 1 {
			handlers.HandleChat(session, rawMsg)
		}
	}
}
