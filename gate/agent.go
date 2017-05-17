package gate

import (
	"Bomber/models"
	"github.com/name5566/leaf/log"
)


func Agent (session *models.Session){
	defer func(){
		session.Close()
	}()
	for {
		// 待完善
		// 只是单纯实现了proto接收，然后广播字符串
		rawMsg, err := session.ReadProtoMessage()
		if err != nil {
			log.Error(err.Error())
			log.Error("read proto message error!")
			return
		}
		Route(session, rawMsg)
	}
}
