package login

import (
	"encoding/json"
	"fmt"
	"go-msg/common/message"
	"go-msg/utils"
	"net"
)

func Login(username, password string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	msg := message.Message{
		Type: message.LoginMsgType,
	}

	loginmsg := message.LoginMsg{
		UserName: username,
		Password: password,
	}

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
	utils.WritePkg(conn, data)
	rspmsg, err := utils.ReadPkg(conn)
	fmt.Println(rspmsg)
	return
}
