package main

import (
	"fmt"
	"gorm.io/gorm"
)

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
	//var studentList []Student
	//DB.Find(&studentList).Delete(&studentList)
	//studentList = []Student{
	//	{ID: 1, Name: "李元芳", Age: 32, Email: PtrString("lyf@yf.com"), Gender: true},
	//	{ID: 2, Name: "张武", Age: 18, Email: PtrString("zhangwu@lly.cn"), Gender: true},
	//	{ID: 3, Name: "枫枫", Age: 23, Email: PtrString("ff@yahoo.com"), Gender: true},
	//	{ID: 4, Name: "刘大", Age: 54, Email: PtrString("liuda@qq.com"), Gender: true},
	//	{ID: 5, Name: "李武", Age: 23, Email: PtrString("liwu@lly.cn"), Gender: true},
	//	{ID: 6, Name: "李琦", Age: 14, Email: PtrString("liqi@lly.cn"), Gender: false},
	//	{ID: 7, Name: "晓梅", Age: 25, Email: PtrString("xiaomeo@sl.com"), Gender: false},
	//	{ID: 8, Name: "如燕", Age: 26, Email: PtrString("ruyan@yf.com"), Gender: false},
	//	{ID: 9, Name: "魔灵", Age: 21, Email: PtrString("moling@sl.com"), Gender: true},
	//}
	//DB.Create(&studentList)

	//var users []Student
	// 查询用户名是枫枫的
	//DB.Where("name = ?", "枫枫").Find(&users)
	//fmt.Println(users)
	//// 查询用户名不是枫枫的
	//DB.Where("name <> ?", "枫枫").Find(&users)
	//fmt.Println(users)
	//// 查询用户名包含 如燕，李元芳的
	//DB.Where("name in ?", []string{"如燕", "李元芳"}).Find(&users)
	//fmt.Println(users)
	//// 查询姓李的
	//DB.Where("name like ?", "李%").Find(&users)
	//fmt.Println(users)
	//// 查询年龄大于23，是qq邮箱的
	//DB.Where("age > ? and email like ?", "23", "%@qq.com").Find(&users)
	//fmt.Println(users)
	//// 查询是qq邮箱的，或者是女的
	//DB.Where("gender = ? or email like ?", false, "%@qq.com").Find(&users)
	//fmt.Println(users)
	// 会过滤零值
	//DB.Where(&Student{Name: "李元芳", Age: 0}).Find(&users)
	//DB.Where(map[string]any{"name": "李元芳", "age": 0}).Find(&users)
	// SELECT * FROM `students` WHERE `age` = 0 AND `name` = '李元芳'

	//type User struct {
	//	Name string
	//	Age  int
	//}
	//var userList []User
	////DB.Table("tb_Students").Select("Name", "Gender").Scan(&users)
	//DB.Model(Student{}).Select("Name", "Age").Scan(&userList)
	//fmt.Println(userList)

	//DB.Table("tb_Students").Limit(2).Offset(0).Find(&userList)
	//limit := 3
	//page := 1
	//DB.Table("tb_Students").Limit(limit).Offset((page - 1) * limit).Find(&userList) //分页查询
	//
	//fmt.Println(userList)

	//var ageList []int

	//DB.Model(Student{}).Distinct("Age").Scan(&ageList)
	//DB.Model(&Student{}).Select("distinct Age").Scan(&ageList)
	//fmt.Println(ageList)

	//var countList []int
	//
	//DB.Model(Student{}).Select("count(id)").Group("gender").Scan(&countList)
	//fmt.Println(countList)

	//type Group struct {
	//	Count    int
	//	Gender   string
	//	NameList string
	//}
	//
	//var groupList []Group
	//DB.Model(Student{}).Select("count(id) as count", "Gender").Group("gender").Scan(&groupList)
	//DB.Model(Student{}).
	//	Select(
	//		"group_concat(Name) as NameList",
	//		"count(id) as count",
	//		"Gender",
	//	).
	//	Group("Gender").
	//	Scan(&groupList)

	//DB.Raw("SELECT group_concat(Name) as NameList,count(id) as count,`Gender` FROM `tb_Students` GROUP BY `Gender`").
	//	Scan(&groupList)

	//DB.Raw("select * from tb_students where age >(select avg(age) from tb_students)").Scan(&studentList)
	//DB.Where("age > (?)", DB.Model(Student{}).Select("avg(age)")).Find(&studentList)
	//fmt.Println(studentList)

	//DB.Where("name = ? and age = ?", "枫枫", 23).Find(&studentList)
	//DB.Where("name = @name and age = @age",
	//	sql.Named("name", "枫枫"),
	//	sql.Named("age", 23)).Find(&studentList)
	//DB.Where("name = @name and age = @age",
	//	map[string]any{"name": "枫枫", "age": 23}).Find(&studentList)
	//fmt.Println(studentList)

	var res []map[string]any
	DB.Model(Student{}).Scopes(Age23).Find(&res)
	fmt.Println(res)
}

func Age23(db *gorm.DB) *gorm.DB {
	return db.Where("age > ?", 23)
}

func PtrString(email string) *string {
	return &email
}
