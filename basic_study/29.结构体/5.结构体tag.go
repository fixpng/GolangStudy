package main

import (
	"encoding/json"
	"fmt"
)

type Article struct {
	Title     string `json:"324"`
	Desc      string `json:"desc"`
	Content   string `json:"content"`
	Username  string `json:"-"` //- 不参与序列化
	LookCount int    `json:"lookCount"`
	Free      bool   `json:"free"`
	password  string //小写字母开头的不参与序列化
}

func main() {
	article := Article{
		Title:     "golang文档",
		Desc:      "golang零基础入门",
		Content:   "golang零基础入门",
		Username:  "Nilz",
		LookCount: 1024,
		password:  "#$RE7RF5GR3",
		Free:      true,
	}

	//结构体转json
	jsonData, err := json.Marshal(article)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonStr := string(jsonData)
	fmt.Println(jsonStr, err)

}
