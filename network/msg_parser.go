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

type RawMessage struct {
	Id uint32
	Data []byte
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
	msgIdBuf := make([]byte, parser.msgIdLen)
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
	msgLenBuf := make([]byte, parser.msgLenNum)
	if _, err := io.ReadFull(parser.reader, msgLenBuf); err != nil {
		return r, err
	}

	// 解析msgLen
	var msgLen uint16
	if parser.littleEndian {
		msgLen = binary.LittleEndian.Uint16(msgLenBuf)
	} else {
		msgLen = binary.BigEndian.Uint16(msgLenBuf)
	}
	return msgLen, nil
}

func (parser *MsgParser) ReadMsgPacket() (*RawMessage, error) {
	// 读取数据包
	// 返回协议内和消息结构bytes
	msgId, err := parser.GetMsgId()
	if err != nil {
		return nil, err
	}
	// 获取消息长度
	msgLen, err := parser.GetMsgLen()
	if err != nil {
		return nil, err
	}
	// 获取消息内容
	msgData := make([]byte, msgLen)
	if _, err := io.ReadFull(parser.reader, msgData); err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	resp := &RawMessage{msgId, msgData}
	return resp, err

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


