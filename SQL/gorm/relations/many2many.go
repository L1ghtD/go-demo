package relations

import (
	"fmt"

	"github.com/ra1n6ow/go-demo/SQL/gorm/model"
	"gorm.io/gorm"
)

func CreateMany(db *gorm.DB) {
	// 添加用户并创建角色
	// db.Create(&model.UserModel{
	// 	UserName: "admin",
	// 	Roles: []model.RoleModel{
	// 		{Name: "dev"},
	// 		{Name: "ops"},
	// 	},
	// })

	// 创建用户，添加已有角色
	var roles []model.RoleModel
	db.Find(&roles, "name = ?", "ops")
	db.Create(&model.UserModel{
		UserName: "du1",
		Roles:    roles,
	})
}

func QueryMany(db *gorm.DB) {
	// 查询用户，显示用户的角色
	var user model.UserModel
	db.Preload("Roles").Take(&user, 5)
	fmt.Println(user)

	// 查询角色，显示该角色的用户
	var role model.RoleModel
	db.Preload("Users").Take(&role, 8)
	fmt.Println(role)
}

func UpdateMany(db *gorm.DB) {
	// 方法一
	// 先删除旧的
	// var user model.UserModel
	// db.Preload("Roles").Take(&user, 5)
	// db.Model(&user).Association("Roles").Delete(user.Roles) // DELETE FROM `user_role` WHERE `user_role`.`user_id` = 5 AND `user_role`.`role_id` IN (7,8)
	// // 再添加新的
	// var role model.RoleModel
	// db.Take(&role, 7)
	// db.Model(&user).Association("Roles").Append(&role) // INSERT INTO `user_role` (`user_id`,`role_id`) VALUES (5,7) ON DUPLICATE KEY UPDATE `user_id`=`user_id`

	// 方法二
	var user model.UserModel
	db.Preload("Roles").Take(&user, 6)
	var role model.RoleModel
	db.Take(&role, 7)
	db.Model(&user).Association("Roles").Replace(&role)
}
