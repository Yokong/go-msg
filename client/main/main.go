package main

import (
	"fmt"
	"go-msg/client/process"
	"os"
)

func main() {
	initMainView()
}

func initMainView() {
	var key int

	for {
		fmt.Println("------- 欢迎来到多人聊天系统 ------")
		fmt.Println("1 登录聊天系统")
		fmt.Println("2 注册账号")
		fmt.Println("3 退出系统")
		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println(">>> 登录聊天系统")
			var username string
			var password string
			fmt.Print("用户名: ")
			fmt.Scanf("%s\n", &username)
			fmt.Print("密码: ")
			fmt.Scanf("%s\n", &password)
			u := process.UserProcess{}
			u.Login(username, password)
		case 2:
			fmt.Println(">>> 注册账号")
		case 3:
			fmt.Println(">>> 退出系统")
			os.Exit(0)
		default:
			fmt.Println(">>> 输入有误，请重新输入")
		}
	}
}
