package process

import (
	"encoding/json"
	"fmt"
	"go-msg/common/message"
	"go-msg/utils"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess) LoginProcess(msg *message.Message) (err error) {
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

	tf := utils.Transer{
		Conn: this.Conn,
	}
	tf.WritePkg(data)
	return

}
