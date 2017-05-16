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

type MsgProxy struct {
	msgLenNum    uint16
	msgIdLen     uint16
	littleEndian bool
	conn         *net.TCPConn
}

type RawMessage struct {
	Id uint32
	Data []byte
}


func NewMsgProxy(con *net.TCPConn) *MsgProxy {
	p := &MsgProxy{
		msgLenNum:    4,
		msgIdLen:     4,
		littleEndian: false,
		conn:         con,
	}
	return p
}


func (mp *MsgProxy) GetMsgId() (uint32, error) {
	// 用来获取消息的msgId
	var r uint32
	// 判断消息长度
	msgIdBuf := make([]byte, mp.msgIdLen)
	if _, err := mp.conn.Read(msgIdBuf); err != nil {
		return r, err
	}

	// 解析msgID
	var msgId uint32
	if mp.littleEndian {
		msgId = uint32(binary.LittleEndian.Uint32(msgIdBuf))
	} else {
		msgId = uint32(binary.BigEndian.Uint32(msgIdBuf))
	}
	return msgId, nil
}


func (mp *MsgProxy) GetMsgLen() (uint32, error) {
	// 用来获取消息结构的长度

	var r uint32
	// 判断消息长度
	msgLenBuf := make([]byte, mp.msgLenNum)
	if _, err := io.ReadFull(mp.conn, msgLenBuf); err != nil {
		return r, err
	}

	// 解析msgLen
	var msgLen uint32
	if mp.littleEndian {
		msgLen = binary.LittleEndian.Uint32(msgLenBuf)
	} else {
		msgLen = binary.BigEndian.Uint32(msgLenBuf)
	}
	return msgLen, nil
}

func (mp *MsgProxy) ReadMsgPacket() (*RawMessage, error) {
	// 读取数据包
	// 返回协议内和消息结构bytes
	msgId, err := mp.GetMsgId()
	if err != nil {
		return nil, err
	}
	// 获取消息长度
	msgLen, err := mp.GetMsgLen()
	if err != nil {
		return nil, err
	}
	// 获取消息内容
	msgData := make([]byte, msgLen)
	if _, err := io.ReadFull(mp.conn, msgData); err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	resp := &RawMessage{msgId, msgData}
	return resp, err

}


func(mp *MsgProxy) MessageToPackage(msgId uint32, data proto.Message) (*bytes.Buffer, error){
	// 将消息封装成数据包
	msgBuff, err := proto.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgLen := uint32(len(msgBuff))
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

func (mp *MsgProxy) WriteMessage(msgId uint32, data proto.Message) {
	// 写入数据到连接
	pkg, err := mp.MessageToPackage(msgId, data)
	if err != nil {
		// todo: error handle
	}
	mp.conn.Write(pkg.Bytes())
}


