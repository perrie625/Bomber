package network

import (
	"io"
	"encoding/binary"
	"net"
	"github.com/golang/protobuf/proto"
	"bytes"
)

// 负责协议消息的路由和解析
// -------------------------
// | msgId | msgLenNum | data |
// -------------------------

type MsgParser struct {
	msgLenNum    uint16
	msgIdLen     uint16
	littleEndian bool
	reader *net.TCPConn
}


func NewMsgParser(con *net.TCPConn) *MsgParser {
	p := &MsgParser{
		msgLenNum:    2,
		msgIdLen:     4,
		littleEndian: false,
		reader:	      con,
	}
	return p
}


func (parser *MsgParser) GetMsgId() (uint32, error) {
	// 用来获取消息的msgId
	var r uint32
	// 判断消息长度
	var l [4]byte
	msgIdBuf := l[:parser.msgIdLen]
	if _, err := parser.reader.Read(msgIdBuf); err != nil {
		return r, err
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

	var r uint16
	// 判断消息长度
	var l [2]byte
	msgIdBuf := l[:parser.msgLenNum]
	if _, err := io.ReadFull(parser.reader, msgIdBuf); err != nil {
		return r, err
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

func MessageToPackage(msgId uint32, data proto.Message) (*bytes.Buffer, error){
	// 将消息封装成数据包
	msgBuff, err := proto.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgLen := uint16(len(msgBuff))
	var pkg *bytes.Buffer = new(bytes.Buffer)
	// 写入协议ID
	err = binary.Write(pkg, binary.BigEndian, msgId)
	if err != nil {
		return nil, err
	}

	// 写入数据长度
	err = binary.Write(pkg, binary.BigEndian, msgLen)
	if err != nil {
		return nil, err
	}
	// 写入消息内容
	err = binary.Write(pkg, binary.BigEndian, msgBuff)
	return pkg, nil
}

func WriteMessage(con *net.TCPConn, msgId uint32, data proto.Message) {
	// 写入数据到连接
	pkg, err := MessageToPackage(msgId, data)
	if err != nil {
		// todo: error handle
	}
	con.Write(pkg.Bytes())
}


