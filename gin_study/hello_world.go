package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(context *gin.Context) {
	context.String(http.StatusOK, "Hello World")
}

func main() {
	// 创建一个默认的路由
	router := gin.Default()

	// 绑定路由规则和路由函数，访问/index 的路由，将由对应的函数去处理
	router.GET("/index", Index)

	//启动监听,gin会把web服务运行在本机的0.0.0.0:8080 端口上
	//router.Run(":8080") //不写ip则默认0.0.0.0
	//router.Run("127.0.0.1:8080")

	// 用原生http服务方式, router.Run本质就是http.ListenAndServe的进一步封装
	http.ListenAndServe(":8080", router)
}
