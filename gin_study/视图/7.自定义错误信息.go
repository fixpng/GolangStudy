package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// GetValidMsg 返回结构体中的msg参数
func GetValidMsg(err error, obj any) string {

	getObj := reflect.TypeOf(obj)
	// 将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		// 断言成功
		for _, e := range errs {
			// 循环每一个错误信息
			// 根据报错字段名，获取结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}

func main() {
	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
		type User struct {
			Name string `json:"name" binding:"required"  msg:"用户名校验失败"`
			Age  int    `json:"age" binding:"required" msg:"请输入年龄"`
		}
		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			errMsg := GetValidMsg(err, &user)
			c.JSON(200, gin.H{"msg": errMsg})
			fmt.Println(errMsg)
			return
		}
		c.JSON(200, gin.H{"data": user})

	})
	router.Run(":8080")
}
