package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		data, err := os.ReadFile("nilzStudy/37.http/server/index.html")
		if err != nil {
			fmt.Println(err)
		}
		res.Write(data)
		//res.Write([]byte("hello 茜茜 GET"))
	case "POST":
		// 获取body数据
		data, err := io.ReadAll(req.Body)
		// 拿请求头
		contenType := req.Header.Get("Content-Type")
		fmt.Println(contenType)

		if err != nil {
			fmt.Println(err)
		}
		ma := make(map[string]string)
		json.Unmarshal(data, &ma)
		fmt.Println(ma)
		fmt.Println(ma["username"])
		fmt.Println(string(data))

		// 设置响应头
		header := res.Header()
		header["token"] = []string{"t3g6ted8sw13gfg43tj7hu6u"}

		res.Write([]byte("hello 茜茜 POST"))

	}
}

func main() {
	//1. 绑定回调
	http.HandleFunc("/index", IndexHandler)

	//2. 注册服务
	fmt.Println("listenserver on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
