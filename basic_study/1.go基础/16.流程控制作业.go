package main

import "fmt"

func main() {
	var (
		username   string
		password   string
		rePassword string
	)

	for {
		fmt.Println("请输入用户名")
		fmt.Scanln(&username)
		fmt.Println("请输入密码")
		fmt.Scanln(&password)
		fmt.Println("请再次输入密码")
		fmt.Scanln(&rePassword)

		if username == "" || password == "" || rePassword == "" {
			fmt.Println("输入不能为空")
			continue
		}

		if password != rePassword {
			fmt.Println("两次密码不一致")
			continue
		}

		fmt.Println("注册成功！")
		break

	}
}
