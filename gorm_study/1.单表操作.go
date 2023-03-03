package main

//type Student struct {
//	ID     uint   `gorm:"size:10"` // 默认使用ID作为主键
//	Name   string `gorm:"size:16"`
//	Age    int    `gorm:"size:3"`
//	Gender bool
//	Email  *string `gorm:"size:128"` // 使用指针是为了存空值
//	//Type  string  `gorm:"column:_type;size:4"`
//	//Date  string  `gorm:"default:2022-12-30"`
//}
//
//func (user *Student) BeforeCreate(tx *gorm.DB) (err error) {
//	email := "test@qq.com"
//	user.Email = &email
//	return nil
//}

func main() {

	//建表
	//DB.Debug().AutoMigrate(&Student{})
	//
	//email := "897491068@qq.com"
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

	////批量插入
	//var studentList []Student
	//for i := 0; i < 10; i++ {
	//	studentList = append(studentList, Student{
	//		Name:   fmt.Sprintf("小张%d", i+1),
	//		Age:    20 + i,
	//		Gender: false,
	//		Email:  &email,
	//	})
	//}
	//err := DB.Create(&studentList).Error
	//fmt.Println(err)

	// 单条记录查询
	//var student Student
	//DB = DB.Session(&gorm.Session{
	//	Logger: mysqlLogger,
	//})
	//DB.Take(&student)
	//fmt.Println(student)
	//student = Student{}
	//DB.First(&student)
	//fmt.Println(student)
	//student = Student{}
	//DB.Last(&student)
	//fmt.Println(student)

	//err := DB.Take(&student, 66).Error
	//fmt.Println(err == gorm.ErrRecordNotFound)

	//DB.Take(&student, "name = ?", "小张9")
	//fmt.Println(student)

	//student.ID = 5
	//count := DB.Take(&student).RowsAffected
	//fmt.Println(student, count)

	//查询多条记录
	//var studentList []Student

	//count := DB.Find(&studentList).RowsAffected
	//fmt.Println(count)
	////for _, student := range studentList {
	////	fmt.Println(student)
	////}
	//data, _ := json.Marshal(studentList)
	//fmt.Println(string(data))

	//DB.Find(&studentList, []int{4, 7, 9})
	//DB.Find(&studentList, "name in (?)", []string{"茜茜", "小张6"})
	//fmt.Println(studentList)

	//save更新
	//var student Student
	//DB.Take(&student, 11)
	//student.Age, student.Name = 0, "我是你爹"
	//DB.Select("Name").Save(&student)

	//var students []Student
	//DB.Find(&students, []int{12, 13, 11}).Update("Gender", true)
	//DB.Find(&students, []int{12, 13, 11}).Updates(Student{Age:    50,Gender: true})
	//email := "123548994"
	//DB.Model(&Student{}).Where("age = ?", 21).Updates(map[string]any{
	//	"email":  &email,
	//	"gender": false,
	//})

	//删除
	//var student Student
	//// student 的 ID 是 `10`
	//DB.Delete(&student, 10)
	//// DELETE from students where id = 10;
	//DB.Delete(&Student{}, []int{1, 2, 3})
	//
	//var studentList []Student
	//// 查询到的切片列表
	//DB.Delete(&studentList)

	//DB.Create(&Student{
	//	Name: "xxxx",
	//	Age:  2,
	//})

}
