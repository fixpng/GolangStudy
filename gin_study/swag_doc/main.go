package main

import (
	_ "GO20230123/gin_study/swag_doc/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

// UserList 用户列表
// @Tags 用户管理
// @Summary 用户列表
// @Description 返回一个用户列表，可根据查询参数指定
// @Param limit query string false "返回多少条"
// @Router /api/users [get]
// @Produce json
// @Success 200 {object} string
func UserList(c *gin.Context) {
	c.JSON(200, gin.H{"mag": "你好"})
}

// @title 倪半仙博客系统后端api文档
// @version 1.0
// @description 这是倪半仙博客系统后端api文档
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	router := gin.Default()
	router.GET("/api/users", UserList)
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
