package main

import (
	"fmt"
	"go-msg/server/model"
	"net"
	"time"
)

func main() {
	// 初始化pool, userDao
	initPool(16, 0, "localhost:6379", 300 * time.Second)
	initUserDao()

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
		p := Process{Conn: conn}
		go p.ProcessHandle()
	}
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}
