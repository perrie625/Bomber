package tools

import (
	"encoding/json"
	"io/ioutil"
	"github.com/name5566/leaf/log"
	"net"
	"fmt"
)

type serverConfig struct {
	DBUrl string	`json:"db_url"`
	MsgLenNum uint16	`json:"msg_len_num"`
	MsgIdLenNum uint16	`json:"msg_id_len_num"`
	Encrypt bool		`json:"encrypt"`
	SecretKey string	`json:"secret_key"`
	Port uint16 		`json:"port"`
	BindIp string		`json:"bind_ip"`
}

func (s *serverConfig) GetTcpAddr() *net.TCPAddr {
	addrString := fmt.Sprintf("%s:%d", s.BindIp, s.Port)
	tcpAddr, _ := net.ResolveTCPAddr("tcp", addrString)
	return tcpAddr
}


var ServerConfig *serverConfig

func initServerConfig() {
	data, err := ioutil.ReadFile("./Server.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	var s serverConfig
	if err := json.Unmarshal(data, &s); err != nil {
		log.Fatal(err.Error())
	}
	ServerConfig = &s
}


func init() {
	initServerConfig()
}