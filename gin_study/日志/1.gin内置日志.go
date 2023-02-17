package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {}

func main() {
	gin.SetMode(gin.ReleaseMode) //不显示debug日志

	//file, _ := os.Create("./gin_study/gin.log")
	// file：输出到文件，os.Stdout：控制台输出
	//gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	// 定义路由格式
	//gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	//	log.Printf(
	//		"[ xixi ] %s %s %s %d \n",
	//		httpMethod, absolutePath, handlerName, nuHandlers,
	//	)
	//}

	router := gin.Default()
	router.GET("/index", index, index)
	router.POST("/users", func(c *gin.Context) {})
	router.POST("/articles", func(c *gin.Context) {})

	api := router.Group("api")

	api.DELETE("/articles/:id", func(c *gin.Context) {})

	for _, info := range router.Routes() {
		fmt.Println(info.Path, info.Method, info.Handler)
	}

	router.Run() // 默认":8080"

}
