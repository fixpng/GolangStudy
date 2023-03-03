package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServe := http.FileServer(http.Dir("tool_demo/go-easy-http-serve-1/www"))
	http.Handle("/", fileServe)
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Listening on port 8080", "http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loginHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		log.Fatal(err)
		return
	}
	userName := request.FormValue("username")
	fmt.Println(userName)
	fmt.Fprintf(writer, "Hello %s", userName)

}
