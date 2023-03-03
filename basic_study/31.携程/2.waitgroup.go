package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func SendTelCode() {
	fmt.Println("发送验证码开始")
	time.Sleep(3 * time.Second)
	fmt.Println("发送验证码完成！")
	wg.Done()
}

func main() {
	wg.Add(1)
	// 实现用户注册功能
	fmt.Println("用户注册校验完成")
	// 发送验证码
	go SendTelCode()
	wg.Wait()
	fmt.Println("验证码已发送，请注意查收。。。")

}
