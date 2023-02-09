package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func _string(c *gin.Context) {
	c.String(http.StatusOK, "你好啊")
}

// 响应json数据 重点
func _json(c *gin.Context) {
	// json响应结构体
	/*
		type UserInfo struct {
			UserName string `json:"user_name"`
			Age      int    `json:"age"`
			Password string `json:"-"` // 忽略转换为json
		}
		user := UserInfo{"茜茜", 20, "123456"}
		c.JSON(200, user)
	*/

	// json 响应map
	/*
		userMap := map[string]string{
			"user_name": "茜茜2",
			"age":       "23",
		}
		c.JSON(200, userMap)
	*/

	// 直接响应json
	c.JSON(200, gin.H{"username": "茜茜3", "age": 21})
}

// 响应xml
func _xml(c *gin.Context) {
	c.XML(200, gin.H{"user": "茜茜5", "message": "hey", "status": http.StatusOK, "data": gin.H{"username": "茜茜3", "age": 21}})
}

// 响应yaml
func _yaml(c *gin.Context) {
	c.YAML(200, gin.H{"user": "茜茜6", "message": "hey", "status": http.StatusOK, "data": gin.H{"username": "茜茜3", "age": 21}})
}

// 响应html
func _html(c *gin.Context) {
	type UserInfo struct {
		UserName string `json:"user_name"`
		Age      int    `json:"age"`
		Password string `json:"-"` // 忽略转换为json
	}
	user := UserInfo{"茜茜", 20, "123456"}

	c.HTML(200, "index.html", gin.H{"user": user})
}

// 重定向
func _redirect(c *gin.Context) {
	//c.Redirect(302, "/html")
	c.Redirect(302, "http://www.baidu.com")
	//301 Moved Permanently
	//被请求的资源已永久移动到新位置，并且将来任何对此资源的引用都应该使用本响应返回的若干个 URI 之一。如果可能，拥有链接编辑功能的客户端应当自动把请求的地址修改为从服务器反馈回来的地址。除非额外指定，否则这个响应也是可缓存的。

	//302 Found
	//请求的资源现在临时从不同的 URI 响应请求。由于这样的重定向是临时的，客户端应当继续向原有地址发送以后的请求。只有在Cache-Control或Expires中进行了指定的情况下，这个响应才是可缓存的。
}

func main() {
	router := gin.Default()
	// 加载模板目录下所有的模板文件
	router.LoadHTMLGlob("gin_study/templates/*")
	// 在golang中，没有相对文件的路径，他只有相对项目的路径
	// 网页请求这个静态目录的前缀，第二个参数是一个目录（注意前缀不要重复）
	router.StaticFS("static", http.Dir("gin_study/static/static"))
	// 配置单个文件，网页请求路径，文件的路径
	router.StaticFile("xx1.jpg", "gin_study/static/xx1.jpg")

	router.GET("/", _string)
	router.GET("/json", _json)
	router.GET("/xml", _xml)
	router.GET("/yaml", _yaml)
	router.GET("/html", _html)
	router.GET("/baidu", _redirect)

	router.Run(":8080")
}
