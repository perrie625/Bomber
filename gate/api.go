package gate

import "github.com/golang/protobuf/proto"

var (
	// 协议id和handler的映射
	handlers map[uint16]interface{}
	// 协议ID和消息结构的映射
	messages map[uint16]*proto.Message
	// 消息结构和协议ID的映射
	protoIds map[string]uint16
)

