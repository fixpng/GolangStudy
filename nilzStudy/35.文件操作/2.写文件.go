package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 如果文件不存在就创建写入
	// os.O_RDWR 覆写
	// os.O_APPEND 追加
	/*
		file, err := os.OpenFile("w.txt", os.O_CREATE|os.O_APPEND, 077223)
		if err != nil {
			fmt.Println(err)
			return
		}

		n, _ := file.WriteString("我是你爹\n")
		n, _ = file.Write([]byte("我是你爹\n"))
		fmt.Println(n)
	*/

	//快速写
	/*
		err := os.WriteFile("w.txt", []byte("我是你爹，我是你爹！"), 06566)
		fmt.Println(err)
	*/

	// 带缓冲写入（大文件写入）

	file, err := os.OpenFile("w.txt", os.O_CREATE|os.O_APPEND, 077223)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		buf.WriteString(fmt.Sprintf("%d %s\n", i+1, "我是你爹"))
		fmt.Printf("%d %s\n", i+1, "我是你爹")
	}

	buf.Flush()

}
