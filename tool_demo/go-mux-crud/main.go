package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Brand struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}

type HotDryNoodles struct {
	ID    int    `json:"ID"`
	Name  string `json:"Name"`
	Price int    `json:"Price"`
	Brand *Brand `json:"Brand"`
}

var hotDryNoodles []HotDryNoodles

func fakeData() {
	// 标准热干面 5 蔡林记
	// 麻辣热干面 6 蔡林记
	// 热干面加蛋 7 陈记
	//hotDryNoodles = append(hotDryNoodles, HotDryNoodles{})
	hotDryNoodles = append(hotDryNoodles, HotDryNoodles{
		1,
		"标准热干面",
		5,
		&Brand{1, "蔡林记"},
	})
	hotDryNoodles = append(hotDryNoodles, HotDryNoodles{
		2,
		"麻辣热干面",
		6,
		&Brand{1, "蔡林记"},
	})
	hotDryNoodles = append(hotDryNoodles, HotDryNoodles{
		1,
		"热干面加蛋",
		7,
		&Brand{2, "陈记"},
	})
}

func main() {
	fakeData()
	r := mux.NewRouter()
	// api/noodles crud

	r.HandleFunc("/api/noodles", getAllNoodles).Methods("GET")
	r.HandleFunc("/api/noodles/{id}", getAllNoodle).Methods("GET")
	r.HandleFunc("/api/noodles", createNoodle).Methods("POST")
	r.HandleFunc("/api/noodles/{id}", updateNoodle).Methods("PUT")
	r.HandleFunc("/api/noodles/{id}", deleteNoodle).Methods("DELETE")

	log.Println("Listening on port 8080", "open http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func updateNoodle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range hotDryNoodles {
		id, _ := strconv.Atoi(params["id"])
		if item.ID == id {
			hotDryNoodles = append(hotDryNoodles[:index], hotDryNoodles[index+1:]...)

			var hotDryNoodle HotDryNoodles
			_ = json.NewDecoder(request.Body).Decode(&hotDryNoodle)
			hotDryNoodle.ID = id

			hotDryNoodles = append(hotDryNoodles, hotDryNoodle)
			json.NewEncoder(writer).Encode(hotDryNoodles)
			return
		}
	}

}

func deleteNoodle(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("deleteHotDryNoodle")
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range hotDryNoodles {
		id, _ := strconv.Atoi(params["id"])
		if item.ID == id {
			hotDryNoodles = append(hotDryNoodles[:index], hotDryNoodles[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(hotDryNoodles)

}

func createNoodle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var noodle HotDryNoodles
	_ = json.NewDecoder(request.Body).Decode(&noodle)
	hotDryNoodles = append(hotDryNoodles, noodle)
	json.NewEncoder(writer).Encode(noodle)

}

func getAllNoodle(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range hotDryNoodles {
		if item.ID == toInt(params["id"]) {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	json.NewEncoder(writer).Encode(&HotDryNoodles{})
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func getAllNoodles(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(hotDryNoodles)
}
