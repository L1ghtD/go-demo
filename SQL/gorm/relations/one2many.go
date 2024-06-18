package relations

import (
	"fmt"

	"github.com/ra1n6ow/go-demo/SQL/gorm/model"
	"gorm.io/gorm"
)

func Create(db *gorm.DB) {

	// 同时创建用户和文章
	// a1 := model.ArticleM{
	// 	Title: "Python",
	// }
	// a2 := model.ArticleM{
	// 	Title: "Golang",
	// }
	// user := model.UserM{
	// 	Name:     "杜一",
	// 	Articles: []model.ArticleM{a1, a2},
	// }
	// db.Create(&user)

	// 创建文章并关联已有用户
	a3 := model.ArticleM{Title: "Java", UserID: 1}
	db.Create(&a3)

	var a4 model.UserM
	db.Take(&a4, 1)
	db.Create(&model.ArticleM{Title: "C++", User: a4})
}

func ForeignKey(db *gorm.DB) {
	// 给现有的用户绑定文章
	// var user model.UserM
	// db.Take(&user, 1)
	// var article model.ArticleM
	// db.Take(&article, 6)
	// // Articles 是 UserM 的属性(Has many)
	// db.Model(&user).Association("Articles").Append(&article)

	// 给文章绑定用户
	var user model.UserM
	db.Take(&user, 1)
	var article model.ArticleM
	db.Take(&article, 7)
	// User 是 Article 的属性 (belongs to)
	db.Model(&article).Association("User").Append(&user)
}

func Query(db *gorm.DB) {
	// 这样显示不了文章列表
	var user model.UserM
	db.Take(&user, 1)
	fmt.Println(user)

	// 预加载
	user = model.UserM{}
	// Articles 就是外键关联属性名
	// db.Preload("Articles").Take(&user, 1)
	// 带条件的预加载
	db.Preload("Articles", "id = ?", 6).Take(&user, 1)

	fmt.Println(user)

	var article model.ArticleM
	// User 就是 ArticleM 属性名
	db.Preload("User").Take(&article, 6)
	fmt.Println(article)

	// 嵌套预加载
	// 查询指定文章所属用户下面的所有文章
	article = model.ArticleM{}
	db.Preload("User.Articles").Take(&article, 6)
	fmt.Println(user)

	// 自定义预加载
	user = model.UserM{}
	db.Preload("Articles", func(db *gorm.DB) *gorm.DB {
		return db.Where("id in ?", []int{3, 4})
	}).Take(&user, 1)
	fmt.Println(user)
}

func Delete(db *gorm.DB) {
	// 级联删除，与用户关联的文章也会删除
	var user model.UserM
	db.Take(&user, 1)
	db.Select("Articles").Delete(&user)

	// 删除用户，将与之关联的文章的外键设置为 null
	user = model.UserM{}
	db.Preload("Articles").Take(&user, 2)
	db.Model(&user).Association("Articles").Delete(&user.Articles)
}
