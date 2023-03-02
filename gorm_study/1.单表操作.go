package main

import "fmt"

type Student struct {
	ID     uint   `gorm:"size:10"` // 默认使用ID作为主键
	Name   string `gorm:"size:16"`
	Age    int    `gorm:"size:3"`
	Gender bool
	Email  *string `gorm:"size:128"` // 使用指针是为了存空值
	//Type  string  `gorm:"column:_type;size:4"`
	//Date  string  `gorm:"default:2022-12-30"`
}

func main() {

	//建表
	//DB.Debug().AutoMigrate(&Student{})
	//
	email := "897491068@qq.com"
	//添加记录
	//s1 := Student{
	//	Name:   "茜茜",
	//	Age:    20,
	//	Gender: false,
	//	Email:  &email,
	//}
	//err := DB.Create(&s1).Error
	//fmt.Println(s1)
	//fmt.Println(err)

	//批量插入
	var studentList []Student
	for i := 0; i < 10; i++ {
		studentList = append(studentList, Student{
			Name:   fmt.Sprintf("小张%d", i+1),
			Age:    20 + i,
			Gender: false,
			Email:  &email,
		})
	}
	err := DB.Create(&studentList).Error
	fmt.Println(err)

}
