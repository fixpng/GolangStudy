package main

import (
	"time"
)

type Article struct {
	ID    uint
	Title string
	Tags  []Tag `gorm:"many2many:article_tags"`
}

type Tag struct {
	ID   uint
	Name string
}

type ArticleTag struct {
	ArticleID uint `gorm:"primaryKey"`
	TagID     uint `gorm:"primaryKey"`
	CreatedAt time.Time
}

func main() {
	// 设置Article的Tags表为ArticleTag
	//DB.SetupJoinTable(&Article{}, "Tags", &ArticleTag{})
	// 如果tag要反向应用Article，那么也得加上
	// DB.SetupJoinTable(&Tag{}, "Articles", &ArticleTag{})
	//err := DB.AutoMigrate(&Article{}, &Tag{}, &ArticleTag{})
	//fmt.Println(err)

	//DB.SetupJoinTable(&Article{}, "Tags", &ArticleTag{}) // 要设置这个，才能走到我们自定义的连接表
	//DB.Create(&Article{
	//	Title: "flask零基础入门",
	//	Tags: []Tag{
	//		{Name: "python"},
	//		{Name: "后端"},
	//		{Name: "web"},
	//	},
	//})
	// CreatedAt time.Time 由于我们设置的是CreatedAt，gorm会自动填充当前时间，
	// 如果是其他的字段，需要使用到ArticleTag 的添加钩子 BeforeCreate

	//DB.SetupJoinTable(&Article{}, "Tags", &ArticleTag{})
	//var tags []Tag
	//DB.Find(&tags, "name in ?", []string{"python", "web"})
	//DB.Create(&Article{
	//	Title: "flask请求对象",
	//	Tags:  tags,
	//})

}
