package gate

import (
	"Bomber/protodata"
	"github.com/golang/protobuf/proto"
	"time"
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
		_, err := session.MsgParser.GetMsgId()
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
		resp := new(protodata.SaidMessage)
		resp.Name = session.RemoteAddr
		now := time.Now()
		resp.Time = now.Format("2006-01-02 15:04:05")
		resp.Words = msg.Words
		session.Room.BroadCast(resp)
	}
}
