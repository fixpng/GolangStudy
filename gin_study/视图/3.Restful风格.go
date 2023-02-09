package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ArticleModel struct {
	Title   string `json:"title"`
	Context string `json:"context"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// 文章列表
func _getList(c *gin.Context) {
	articleList := []ArticleModel{
		{"Go语言入门", "这边文章是《Go语言入门》"},
		{"python语言入门", "这边文章是《python语言入门》"},
		{"JavaScript语言入门", "这边文章是《JavaScript语言入门》"},
	}
	c.JSON(200, Response{0, articleList, "成功"})
}

func _getDetail(c *gin.Context) {
	// 获取param中的id
	fmt.Println(c.Param("id"))
	article := ArticleModel{
		"Go语言入门", "这边文章是《Go语言入门》",
	}
	c.JSON(200, Response{0, article, "成功"})
}

// 创建文章
func _create(c *gin.Context) {

	c.JSON(200, gin.H{})
}
func _update(c *gin.Context) {
	c.JSON(200, gin.H{})
}
func _delete(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func main() {
	router := gin.Default()
	router.GET("/articles", _getList)       // 文章列表
	router.GET("/articles/:id", _getDetail) // 文章详情
	router.POST("/articles", _create)       // 添加文章
	router.PUT("/articles/:id", _update)    // 编辑文章
	router.DELETE("/articles/:id", _delete) // 删除文章

	router.Run(":8080")

}
