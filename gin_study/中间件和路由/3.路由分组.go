package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type ArticleInfo struct {
	Title   string `json:"title"`
	Content int    `json:"content"`
}
type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func UserListView(c *gin.Context) {
	var userList []UserInfo = []UserInfo{
		{"徐徐", 10},
		{"安家", 12},
		{"lingbia", 66},
	}
	c.JSON(200, Response{0, userList, "请求成功·"})

}
func ArticleListView(c *gin.Context) {
	var articleInfo []ArticleInfo = []ArticleInfo{
		{"python入门", 2561},
		{"安烦烦烦烦烦烦", 61512},
	}
	c.JSON(200, Response{0, articleInfo, "请求成功·"})

}

func UserRouterUnit(router *gin.RouterGroup) {
	userManger := router.Group("user_manger")
	{
		userManger.GET("/users", UserListView)
	}
}
func ArticleRouterUnit(router *gin.RouterGroup) {
	articleManger := router.Group("article_manger")
	{
		articleManger.GET("/articles", ArticleListView)

	}
}

func main() {
	router := gin.Default()
	api := router.Group("api")
	UserRouterUnit(api)
	ArticleRouterUnit(api)

	router.Run(":8080")

}
