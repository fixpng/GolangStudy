package main

import (
	"GO20230123/nilzStudy/init_study"
	"fmt"
)

func init() {
	fmt.Println("INIT文件下的init被调用。。。了")
}
func main() {
	fmt.Println(init_study.Name)

}
