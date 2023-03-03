package main

import "fmt"

type myMap[K string | int, V any] map[K]V
type _User struct {
	Name string
}

func main() {
	m1 := myMap[string, string]{
		"key": "fengfeng",
		"sad": "4tdf",
	}
	fmt.Println(m1)

	m2 := myMap[int, string]{
		1: "fengfeng",
		2: "user",
	}
	fmt.Println(m2)

	m3 := myMap[int, _User]{
		1: _User{"你半仙"},
		2: _User{"你da仙"},
	}
	fmt.Println(m3)
}
