package network

import (
	"io"
	"encoding/binary"
	"net"
)

// 负责协议消息的路由和解析
// -------------------------
// | msgId | msgLenNum | data |
// -------------------------

type MsgParser struct {
	msgLenNum    uint16
	msgIdLen     uint16
	littleEndian bool
	reader net.Conn
}


func NewMsgParser(con net.TCPConn) *MsgParser {
	p := &MsgParser{
		msgLenNum:    2,
		msgIdLen:     2,
		littleEndian: false,
		reader:	      con,
	}
	return p
}


func (parser *MsgParser) GetMsgId() (uint32, error) {
	// 用来获取消息的msgId

	// 判断消息长度
	var l [4]byte
	msgIdBuf := l[:parser.msgIdLen]
	if _, err := io.ReadFull(parser.reader, msgIdBuf); err != nil {
		return nil, err
	}

	// 解析msgID
	var msgId uint32
	if parser.littleEndian {
		msgId = uint32(binary.LittleEndian.Uint32(msgIdBuf))
	} else {
		msgId = uint32(binary.BigEndian.Uint32(msgIdBuf))
	}
	return msgId, nil
}


func (parser *MsgParser) GetMsgLen() (uint16, error) {
	// 用来获取消息结构的长度

	// 判断消息长度
	var l [2]byte
	msgIdBuf := l[:parser.msgLenNum]
	if _, err := io.ReadFull(parser.reader, msgIdBuf); err != nil {
		return nil, err
	}

	// 解析msgLen
	var msgLen uint16
	if parser.littleEndian {
		msgLen = binary.LittleEndian.Uint16(msgIdBuf)
	} else {
		msgLen = binary.BigEndian.Uint16(msgIdBuf)
	}
	return msgLen, nil
}

func (parser *MsgParser) ReadProtoData(msgId uint32){
	// 获取传输的消息结构buff
	//msgId, err := parser.GetMsgId()
	//if err != nil {
	//	// todo: error handle
	//}
	//
	//// 获取消息内容
	//msgData, err := parser.GetMsgData()
	//if err != nil {
	//	// todo: error handle
	//}
	//
	//


}


func (parser *MsgParser) GetMsgData() ([]byte, error){

	msgLen, err := parser.GetMsgLen()
	if err != nil {
		// todo: error handle
		return nil, err
	}

	msgData := make([]byte, msgLen)
	if _, err := io.ReadFull(parser.reader, msgData); err != nil {
		return nil, err
	}

	return msgData, nil
}


