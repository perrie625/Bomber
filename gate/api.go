package gate

import (
	"reflect"
	"errors"
	"Bomber/models"
	"Bomber/network"
	"log"
)

var (
	// 协议ID和消息结构的映射
	messageMap map[int32] *MsgInfo
	// 消息结构和协议ID的映射
	protoIdMap map[reflect.Type]int32
)

type MsgHandler func(*models.Session, *network.RawMessage)

type MsgInfo struct {
	msgId int32
	Handler MsgHandler
	ProtoMsg reflect.Type
}

func init() {
	messageMap = make(map[int32] *MsgInfo)
	protoIdMap = make(map[reflect.Type] int32)
}



func RegisterHandler(msgId int32, handler MsgHandler, msgType reflect.Type) (r error){
	tmp := &MsgInfo{
		msgId, handler, msgType,
	}
	messageMap[msgId] = tmp
	defer func(){
		if err := recover();err != nil {
			r = errors.New("msgType type error")
		}
	}()
	protoIdMap[msgType] =msgId
	return nil
}


func Route(session *models.Session, rawMsg *network.RawMessage){
	msgInfo := messageMap[rawMsg.Id]
	if msgInfo == nil {
		log.Println("unknown msg id:", rawMsg.Id)
		return
	}
	msgInfo.Handler(session, rawMsg)
}