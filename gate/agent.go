package gate

import (
	"Bomber/models"
	"log"
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
			log.Println(err.Error())
			log.Println("read proto message error!")
			return
		}
		Route(session, rawMsg)
	}
}
