package process

import (
	"encoding/json"
	"fmt"
	"go-msg/common/message"
	"go-msg/server/model"
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

	user, err := model.MyUserDao.Login(loginMsg.UserName, loginMsg.Password)
	fmt.Println(user)

	if err != nil {
		if err == model.ERROR_PWD_FAILED {
			loginRsp.Code = -1000
			loginRsp.Error = err.Error()
		} else if err == model.ERROR_USER_NOTEXISTS {
			loginRsp.Code = -1001
			loginRsp.Error = err.Error()
		} else {
			loginRsp.Code = -999
			loginRsp.Error = "未知错误"
		}

	} else {
		loginRsp.Code = 0
		fmt.Println(user, "登录成功")
	}

	//if loginMsg.UserName == "yoko" && loginMsg.Password == "123" {
	//	loginRsp.Code = 200
	//} else {
	//	loginRsp.Code = 500
	//	loginRsp.Error = "登录失败"
	//}

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
