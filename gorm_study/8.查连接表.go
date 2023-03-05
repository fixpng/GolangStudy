package main

//
//type UserModel struct {
//	ID       uint
//	Name     string
//	Collects []ArticleModel `gorm:"many2many:user_collect_models;joinForeignKey:UserID;JoinReferences:ArticleID"`
//}
//
//type ArticleModel struct {
//	ID    uint
//	Title string
//}
//
//// UserCollectModel 用户收藏文章表
//type UserCollectModel struct {
//	UserID       uint         `gorm:"primaryKey"` // article_id
//	UserModel    UserModel    `gorm:"foreignKey:UserID"`
//	ArticleID    uint         `gorm:"primaryKey"` // tag_id
//	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"`
//	CreatedAt    time.Time
//}

func main() {
	//DB.SetupJoinTable(&UserModel{}, "Collects", &UserCollectModel{})
	//err := DB.AutoMigrate(&UserModel{}, &ArticleModel{}, &UserCollectModel{})
	//fmt.Println(err)

	//var user UserModel
	//DB.Preload("Collects").Take(&user, "name = ?", "枫枫")
	//fmt.Println(user)

	//var collects []UserCollectModel
	//DB.Find(&collects, "userid = ?", 2)
	//fmt.Println(collects)
	//
	//var collects []UserCollectModel
	//
	//var user UserModel
	//DB.Take(&user, "name = ?", "枫枫")
	//// 这里用map的原因是如果没查到，那就会查0值，如果是struct，则会忽略零值，全部查询
	//DB.Debug().Preload("UserModel").Preload("ArticleModel").Where(map[string]any{"user_id": user.ID}).Find(&collects)
	//
	//for _, collect := range collects {
	//	fmt.Println(collect)
	//}

}
