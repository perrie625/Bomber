package gate

import (
	"reflect"
	"errors"
	"Bomber/models"
	"Bomber/network"
)

var (
	// 协议ID和消息结构的映射
	messageMap map[uint32] *MsgInfo
	// 消息结构和协议ID的映射
	protoIdMap map[reflect.Type]uint32
)

type MsgHandler func(*models.Session, *network.RawMessage)

type MsgInfo struct {
	msgId uint32
	Handler MsgHandler
	ProtoMsg reflect.Type
}

func init() {
	messageMap = make(map[uint32] *MsgInfo)
	protoIdMap = make(map[reflect.Type] uint32)
}



func RegisterHandler(msgId uint32, handler MsgHandler, msgType reflect.Type) (r error){
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
		// todo: error handle
		return
	}
	msgInfo.Handler(session, rawMsg)
}