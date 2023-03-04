package main

// User 用户表 一个用户拥有多篇文章
//type User struct {
//	ID       uint      `gorm:"size:4"`
//	Name     string    `gorm:"size:8"`
//	Articles []Article `gorm:"foreignKey:UID"` // 用户拥有的文章列表
//}
//
//// Article 文章表 一篇文章属于一个用户
//type Article struct {
//	ID    uint   `gorm:"size:4"`
//	Title string `gorm:"size:16"`
//	UID   uint   `gorm:"size:4"`         // 属于   这里的类型要和引用的外键类型一致，包括大小
//	User  User   `gorm:"foreignKey:UID"` // 属于
//}

func main() {
	//DB.AutoMigrate(&User{}, &Article{})

	//创建用户，带上文章
	//DB.Create(&User{
	//	Name: "谢谢",
	//	Articles: []Article{
	//		{
	//			Title: "Golang",
	//		}, {
	//			Title: "Python",
	//		},
	//	},
	//})

	// 创建文章，关联已有用户
	//DB.Create(&Article{
	//	Title: "欢迎来看python",
	//	UID:   1,
	//})

	//DB.Create(&Article{
	//	Title: "欢迎来看golang",
	//	User: User{
	//		Name: "张三",
	//	},
	//})

	//var user User
	//DB.Take(&user, 2)
	//DB.Create(&Article{
	//	Title: "张三写的golang",
	//	User:  user,
	//})

	// 给已有用户绑定文章
	//var user User
	//DB.Take(&user, 2)
	//var article Article
	//DB.Take(&article, 6)
	//DB.Model(&user).Association("Articles").Append(&article)
	// 给文章绑定用户
	//DB.Model(&article).Association("User").Append(&user)

	//var article Article
	//DB.Preload("User").Take(&article)
	//fmt.Println(article)

	//var user User
	//DB.Preload("Articles.User").Take(&user)
	//fmt.Println(user)

	//var user User
	//DB.Preload("Articles", "id > ?", 2).Take(&user)
	//fmt.Println(user)

	//var user User
	//DB.Preload("Articles").Take(&user, 2)
	//fmt.Println(user.Articles)
	//DB.Model(&user).Association("Articles").Delete(user.Articles)
	//DB.Delete(&user)

	//var user User
	//DB.Take(&user, 1)
	//DB.Select("Articles").Delete(&user)
}
