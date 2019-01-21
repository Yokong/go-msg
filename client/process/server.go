package process

import (
	"fmt"
	"go-msg/utils"
	"net"
)

func ShowMenu() {
	fmt.Println("------ xxx 登录成功 -----")
	for {
		fmt.Println("1. 好友列表")
		fmt.Println("2. 发送信息")
		fmt.Println("3. 信息列表")
		fmt.Println("4. 退出系统")
		var key int
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("Yoko...")
		}
	}
}

func serverProcessMsg(conn net.Conn) {
	tf := utils.Transer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端监听消息...")
		rsp, err := tf.ReadPkg()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(rsp)
	}
}
