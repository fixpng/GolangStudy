package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	MergeFile("my_utils/files/tmp_20230308", "my_utils/files/merge_20230309.txt")
}

func MergeFile(inPath string, outPath string) {
	dir, err := os.ReadDir(inPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(dir)

	for _, entry := range dir {
		name := entry.Name()
		nPath := fmt.Sprintf("%s/%s", inPath, name)
		// 如果是目录就递归
		if entry.IsDir() {
			MergeFile(nPath, outPath)
			continue
		}
		fileData := ReadFile(nPath)
		err = WriteFile(fileData, outPath)
		if err != nil {
			return
		}
		fmt.Println(nPath)
	}
}

func WriteFile(s string, outPath string) error {
	file, err := os.OpenFile(outPath, os.O_CREATE|os.O_APPEND, 077223)
	if err != nil {
		return err
	}
	buf := bufio.NewWriter(file)
	buf.WriteString(s)
	buf.Flush()
	return nil
}

func ReadFile(readPath string) string {
	by, err := os.ReadFile(readPath)
	if err != nil {
		fmt.Printf("目录%s 文件写入失败：%s", readPath, err)
		return ""
	}
	return string(by)
}
