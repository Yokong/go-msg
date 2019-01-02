package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8889")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("%s 已连接...", conn)
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024 * 4)
		fmt.Println("读取客户端数据...")
		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println(err)
		}
		fmt.Println("读到的buf: ", buf[:4])
	}
}
