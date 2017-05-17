package network

import (
	"io"
	"encoding/binary"
	"net"
	"github.com/golang/protobuf/proto"
	"bytes"
	"Bomber/tools"
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
	encrypt      bool
	secret_key   []byte
}

type RawMessage struct {
	Id int32
	Data []byte
}


func NewMsgProxy(con *net.TCPConn) *MsgProxy {
	p := &MsgProxy{
		msgLenNum:    tools.ServerConfig.MsgLenNum,
		msgIdLen:     tools.ServerConfig.MsgIdLenNum,
		littleEndian: true,
		conn:         con,
		// encrypt test
		encrypt:      tools.ServerConfig.Encrypt,
		secret_key:   []byte(tools.ServerConfig.SecretKey),
	}
	return p
}


func (mp *MsgProxy) GetMsgId() (int32, error) {
	// 用来获取消息的msgId
	var r int32
	// 判断消息长度
	msgIdBytes := make([]byte, mp.msgIdLen)
	if _, err := mp.conn.Read(msgIdBytes); err != nil {
		return r, err
	}

	// 解析msgID
	msgIdBuffer := bytes.NewBuffer(msgIdBytes)
	if mp.littleEndian {
		binary.Read(msgIdBuffer, binary.LittleEndian, &r)
	} else {
		binary.Read(msgIdBuffer, binary.BigEndian, &r)
	}
	return r, nil
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

	// 解密
	if mp.encrypt {
		msgData, err = tools.Decrypt(msgData, mp.secret_key)
		if err != nil {
			return nil, err
		}
	}

	resp := &RawMessage{msgId, msgData}
	return resp, err

}


func(mp *MsgProxy) MessageToPackage(msgId int32, data proto.Message) (*bytes.Buffer, error){
	// 将消息封装成数据包
	msgBuff, err := proto.Marshal(data)
	if err != nil {
		return nil, err
	}

	// 判断是否加密
	if mp.encrypt {
		msgBuff, err = tools.Encrypt(msgBuff, mp.secret_key)
		if err != nil {
			return nil, err
		}
	}

	msgLen := uint32(len(msgBuff))
	var pkg *bytes.Buffer = new(bytes.Buffer)

	// 写入协议ID和数据长度
	if mp.littleEndian {
		binary.Write(pkg, binary.LittleEndian, msgId)
		binary.Write(pkg, binary.LittleEndian, msgLen)
	} else {
		binary.Write(pkg, binary.BigEndian, msgId)
		binary.Write(pkg, binary.BigEndian, msgLen)
	}

	// 写入消息内容
	err = binary.Write(pkg, binary.BigEndian, msgBuff)
	return pkg, nil
}

func (mp *MsgProxy) WriteMessage(msgId int32, data proto.Message) {
	// 写入数据到连接
	pkg, err := mp.MessageToPackage(msgId, data)
	if err != nil {
		// todo: error handle
	}
	mp.conn.Write(pkg.Bytes())
}


