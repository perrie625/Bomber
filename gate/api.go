package gate

import "github.com/golang/protobuf/proto"

var (
	// 协议id和handler的映射
	handlerMap map[uint16]interface{}
	// 协议ID和消息结构的映射
	messageMap map[uint16]*proto.Message
	// 消息结构和协议ID的映射
	protoIdMap map[string]uint16
)

func init() {
	handlerMap = make(map[uint16]interface{})
	messageMap = make(map[uint16] *proto.Message)
	protoIdMap = make(map[string] uint16)
}

