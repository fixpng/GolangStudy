package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Res struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func _UserListView(c *gin.Context) {
	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var userList = []UserInfo{
		{"徐徐", 10},
		{"安家", 12},
		{"lingbia", 66},
	}
	c.JSON(200, Res{0, userList, "请求成功·"})

}

// Middleware 闭包，返回函数
func Middleware(msg string) gin.HandlerFunc {
	fmt.Println("这里的代码立即执行")
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "1234" {
			c.Next()
			return
		}
		c.JSON(200, Res{1001, nil, msg})
		c.Abort()
	}
}

func _UserRouterUnit(router *gin.RouterGroup) {
	userManger := router.Group("user_manger").Use(Middleware("用户验证失败"))
	{
		userManger.GET("/users", _UserListView)
	}
}

// TimeMiddleware 耗时统计
func TimeMiddleware(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	since := time.Since(startTime)
	// 获取当前请求所对应的函数
	f := c.HandlerName()
	fmt.Printf("函数 %s 耗时 %d\n", f, since)
}

func main() {
	router := gin.New()                                      //不调用中间件，无日志相关信息
	router.Use(gin.Logger(), gin.Recovery(), TimeMiddleware) //加上则与gin.Default() 一致
	//router := gin.Default()

	api := router.Group("api")
	api.GET("/login", func(c *gin.Context) {

		panic("手动报错")
		c.JSON(200, gin.H{"data": "1234"})
	})
	_UserRouterUnit(api)

	router.Run(":8080")

}
