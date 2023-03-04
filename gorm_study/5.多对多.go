package main

import "time"

type Tag struct {
	ID       uint
	Name     string
	Articles []Article `gorm:"many2many:article_tags"` // 用于反向引用
}

type Article struct {
	ID    uint
	Title string
	Tags  []Tag `gorm:"many2many:article_tags;"`
}

type ArticleTag struct {
	ArticleID uint      `gorm:"primaryKey"`
	TagID     uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	DB.SetupJoinTable(&Article{}, "Tags", &ArticleTag{})
	DB.AutoMigrate(&Tag{}, &Article{})
	//DB.Create(&Article{
	//	Title: "python基础",
	//	Tags: []Tag{
	//		{
	//			Name: "PYTHON",
	//		}, {
	//			Name: "后端",
	//		},
	//	},
	//})

	//var tags []Tag
	//DB.Find(&tags, " name  in (?)", "后端")
	//
	//DB.Create(&Article{
	//	Title: "golang",
	//	Tags:  tags,
	//})

	//var article Article
	//DB.Preload("Tags").Take(&article)
	//fmt.Println(article.Tags)

	// 多对多的更新

	// 先删除原有的
	//var article Article
	//DB.Preload("Tags").Take(&article, 1)
	//DB.Model(&article).Association("Tags").Delete(article.Tags)
	////fmt.Println(article)
	//
	//// 在添加新的
	//var tag Tag
	//DB.Take(&tag, 1)
	//
	//DB.Model(&article).Association("Tags").Append(&tag)

	//
	var article Article
	var tag Tag
	DB.Preload("Tags").Take(&article, 1)
	DB.Take(&tag, 2)
	DB.Model(&article).Association("Tags").Replace(&tag)

}
