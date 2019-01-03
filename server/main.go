package main

import (
	"encoding/json"
	"fmt"
	"go-msg/common/message"
	"go-msg/utils"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8889")
	defer listen.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go process(conn)
	}
}

func loginProcess(conn net.Conn, msg *message.Message) (err error) {
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)

	rspMsg := message.Message{
		Type: message.LoginRspType,
	}

	var loginRsp message.LoginRsp
	if loginMsg.UserName == "yoko" && loginMsg.Password == "123" {
		loginRsp.Code = 200
	} else {
		loginRsp.Code = 500
		loginRsp.Error = "登录失败"
	}

	data, err := json.Marshal(loginRsp)
	if err != nil {
		fmt.Println(err)
		return
	}

	rspMsg.Data = string(data)

	data, err = json.Marshal(rspMsg)
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.WritePkg(conn, data)
	return

}

func serverProcess(conn net.Conn, msg *message.Message) (err error) {
	if msg.Type == message.LoginMsgType {
		err = loginProcess(conn, msg)
	}
	return
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		msg, err := utils.ReadPkg(conn)
		if err != nil {
			return
		}
		fmt.Println("msg: ", msg)
		err = serverProcess(conn, &msg)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
