package process

import (
	"encoding/json"
	"fmt"
	"go-msg/common/message"
	"go-msg/utils"
	"net"
)

type UserProcess struct {
	// 暂时无需字段
}

func (this *UserProcess)Login(username, password string) (err error) {
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
	tf := utils.Transer{
		Conn: conn,
	}
	tf.WritePkg(data)
	rspmsg, err := tf.ReadPkg()

	// 反序列化data
	var loginRspMsg message.LoginRsp
	err = json.Unmarshal([]byte(rspmsg.Data), &loginRspMsg)
	if loginRspMsg.Code == 200 {
		fmt.Println("登录成功")
		go serverProcessMsg(conn)
		ShowMenu()
	}
	return
}
