package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func _GetValidMsg(err error, obj interface{}) string {
	// obj为结构体指针
	getObj := reflect.TypeOf(obj)
	// 断言为具体的类型，err是一个接口
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				return f.Tag.Get("msg") //错误信息不需要全部返回，当找到第一个错误的信息时，就可以结束
			}
		}
	}
	return err.Error()
}

// 如果用户名不等于fengfeng就校验失败
func signValid(fl validator.FieldLevel) bool {
	name := fl.Field().Interface().(string)
	if name != "fengfeng" {
		return false
	}
	return true
}

func main() {
	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
		type UserInfo struct {
			Name string `json:"name" binding:"sign" msg:"用户名错误"`
			Age  int    `json:"age" binding:""`
		}
		var user UserInfo
		err := c.ShouldBindJSON(&user)
		if err != nil {
			// 显示自定义的错误信息
			msg := _GetValidMsg(err, &user)
			c.JSON(200, gin.H{"msg": msg})
			return
		}
		c.JSON(200, user)
	})
	// 注册
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sign", signValid)
	}
	router.Run(":80")
}
