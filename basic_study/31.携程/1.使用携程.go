package main

import (
	"fmt"
	"time"
)

func SendCode() {
	t := 3 * time.Second
	fmt.Println("发送验证码开始")
	time.Sleep(t)
	fmt.Println("发送验证码完成")
}

func main() {
	for {
		// 实现用户注册功能
		fmt.Println("用户注册校验完成")
		// 发送验证码
		//SendCode() //会阻塞主线程
		go SendCode() //会阻塞主线程
		fmt.Println("注册验证部分完成")
		var code string
		fmt.Scanln(&code)
	}

}
