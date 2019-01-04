package main

import (
	"fmt"
	"go-msg/common/message"
	"go-msg/server/process"
	"go-msg/utils"
	"net"
)

type Process struct {
	Conn net.Conn
}

func (this *Process) ProcessHandle() {
	defer this.Conn.Close()
	for {
		tf := utils.Transer{Conn: this.Conn}
		msg, err := tf.ReadPkg()
		if err != nil {
			return
		}
		fmt.Println("msg: ", msg)
		err = this.ServerProcess(&msg)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (this *Process) ServerProcess(msg *message.Message) (err error) {
	if msg.Type == message.LoginMsgType {
		up := process.UserProcess{
			Conn: this.Conn,
		}
		err = up.LoginProcess(msg)
	}
	return
}
