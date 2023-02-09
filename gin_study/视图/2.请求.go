package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 查询参数 localhost:8080/query?id=2&user=茜茜&user=猎镰
func _query(c *gin.Context) {
	fmt.Println(c.Query("user"))
	fmt.Println(c.GetQuery("user"))
	fmt.Println(c.QueryArray("user"))          //拿到多个相同的查询参数
	fmt.Println(c.DefaultQuery("addr", "广东省")) // 如果用户没传，就使用默认值
}

// 动态参数 localhost:8080/param/xxx/2
func _param(c *gin.Context) {
	fmt.Println(c.Param("user_id"))
	fmt.Println(c.Param("book_id"))
}

// 表单参数 localhost:8080/form
func _form(c *gin.Context) {
	fmt.Println(c.PostForm("name"))
	fmt.Println(c.PostFormArray("name"))
	fmt.Println(c.DefaultPostForm("addr", "广东省")) // 如果用户没传，就使用默认值
	forms, err := c.MultipartForm()               // 接收所有的form参数，包括文件
	fmt.Println(forms, err)
}

// 封装一个解析json到结构体上的函数
func bindJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err := json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

// 原始参数 localhost:8080/raw
func _raw(c *gin.Context) {
	type User struct {
		Name string `json:"name1"`
		Age  int    `json:"age"`
		Addr string `json:"addr"`
	}
	var user User
	err := bindJson(c, &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	/*
		body, _ := c.GetRawData()
		contentType := c.GetHeader("Content-Type")
		switch contentType {
		case "application/json":
			type User struct {
				Name string `json:"name1"`
				Age  int    `json:"age"`
			}
			var user User
			err := json.Unmarshal(body, &user)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(user)
		}
	*/
}

func main() {
	router := gin.Default()
	router.GET("/query", _query)
	router.GET("/param/:user_id/", _param)
	router.GET("/param/:user_id/:book_id", _param)
	router.POST("/form", _form)
	router.POST("/raw", _raw)
	router.Run(":8080")
}
