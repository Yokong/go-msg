package login

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go-msg/common/message"
	"net"
)

func Login(username, password string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	var msg message.Message
	msg.Type = message.LoginMsgType

	var loginmsg message.LoginMsg
	loginmsg.UserName = username
	loginmsg.Password = password

	data, err := json.Marshal(loginmsg)
	if err != nil {
		fmt.Println(err)
		return
	}
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	dataLen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], dataLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("*********8")
		fmt.Println(err)
		return
	}
	return
}
