package main

import "fmt"

func main() {
	// 切片中的每一个元素是map
	/*
			var studentInfo []map[string]string = []map[string]string{
				{"name": "张三",
					"gender":  "男",
					"addr":    "广东广州",
					"classes": "三年2班",
				},
				{"name": "张五",
					"gender":  "男",
					"addr":    "广东广州",
					"classes": "三年2班",
				},
			}


		var studentInfo []map[string]string

		for i := 1; i < 10; i++ {
			studentInfo = append(studentInfo, map[string]string{
				"name":    fmt.Sprintf("%s-%d", "张", i),
				"gender":  "男",
				"addr":    "广东广州",
				"classes": "三年2班",
			})
		}

		//fmt.Println(studentInfo)

		for _, m := range studentInfo {
			fmt.Printf("name:%s, addr:%s\n", m["name"], m["addr"])
		}
	*/

	//map的value是切片
	//likes := "唱;跳;rap"
	var m map[string][]string = map[string][]string{
		"likes": []string{"唱", "跳", "rap"},
	}

	fmt.Println(m)
	fmt.Println(len(m["likes"]))
}
