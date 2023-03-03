package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 读所有
	/*
		by, err := os.ReadFile("./h.txt")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(by))
	*/

	// 读取指定长度
	/*
		file, err := os.OpenFile("./h.txt", os.O_RDONLY, 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		var readStr []byte = make([]byte, 26)
		n, err := file.Read(readStr)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(n)
		fmt.Println(string(readStr))
	*/

	//读取片段
	/*
		file, err := os.OpenFile("./h.txt", os.O_RDONLY, 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		var readStr []byte = make([]byte, 9)

		file.Seek(35, 0)
		n, err := file.Read(readStr) //读取的字节长度
		fmt.Println(n, err)

		fmt.Println(readStr)
		fmt.Println(string(readStr))
	*/

	// 缓存读取
	/*
		file, _ := os.Open("./h.txt")
		defer file.Close()

		reader := bufio.NewReader(file)

		for {
			b, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			fmt.Println(string(b))
		}
	*/

	// 按指定分隔符读取
	file, _ := os.Open("./h.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadSlice(';')
		if err == io.EOF {
			break
		}
		_len := len(line)
		fmt.Println(string(line[0 : _len-1]))
	}

}
