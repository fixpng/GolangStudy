package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// AppendFile 写文件
func AppendFile(filename string, appendDate string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND, 077223)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := bufio.NewWriter(file)
	buf.WriteString(appendDate)
	buf.Flush()
}

// ReadFile2 数据量较大的文档，输入，输出，行包含的字符串
func ReadFile2(inpath string, outpath string, find string) error {
	fileHanle, err := os.OpenFile(inpath, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer fileHanle.Close()
	reader := bufio.NewReader(fileHanle)

	// 按行处理txt
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if !strings.Contains(string(line), find) {
			continue
		} else {
			fmt.Println(string(line))
			AppendFile(outpath, fmt.Sprintf("%s\n", string(line)))
		}
	}
	return nil
}

func main() {
	// 输入，输出，行包含的字符串
	err := ReadFile2("D:/data/temp/catalina.out-20230217", "./big_txt.txt", "errormsg")
	if err != nil {
		print(err)
	}
}
