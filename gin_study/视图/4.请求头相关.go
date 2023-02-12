package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	router := gin.Default()

	// 请求头的各种获取方式
	router.GET("/", func(c *gin.Context) {
		// 首字母大小写不区分 单词与单词之间用 - 连接
		// 用于获取一个请求头
		fmt.Println(c.GetHeader("Accept-Encoding"))

		// Header 是一个不同的 map[string][]string
		fmt.Println(c.Request.Header)
		// 如果是使用 Get方法或者是.GetHeader,那么可以不用区分大小写，并且返回第一个value
		fmt.Println(c.Request.Header.Get("Accept-Encoding"))
		// 如果使用map的方式取值，请注意大小写问题
		fmt.Println(c.Request.Header["Accept-eNcoding"])

		c.JSON(200, gin.H{"msg": "成功"})
	})

	// 爬虫和用户的区别对待
	router.GET("/index", func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")
		// 用正则去匹配
		// 字符串的包含匹配
		if strings.Contains(userAgent, "python") {
			//爬虫来了
			c.JSON(0, gin.H{"data": "爬虫来了！！"})
			return
		}

		c.JSON(0, gin.H{"data": "识别为用户"})

	})

	// 设置响应头
	router.GET("/res", func(c *gin.Context) {
		c.Header("Token", "dafhdfhfgh$%^#%Y")
		c.Header("Content-Type", "application/text; charset=utf-8")
		c.JSON(0, gin.H{"data": "看看响应头"})
	})

	router.Run(":8080")

}
