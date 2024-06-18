package model

import "time"

/*
CREATE TABLE `student` (

	`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	`name` varchar(16) DEFAULT NULL,
	`age` tinyint(3) unsigned DEFAULT NULL,
	`gender` tinyint(1) DEFAULT NULL,
	`birthday` timestamp NOT NULL DEFAULT current_timestamp(),
	PRIMARY KEY (`id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
*/
type Student struct {
	ID       int64     `gorm:"primaryKey"`
	Name     string    `gorm:"column:name"`
	Age      int       `gorm:"column:age"`
	Gender   int       `gorm:"column:gender"`
	Email    *string   `gorm:"column:email"`
	Birthday time.Time `gorm:"column:birthday"`
}

func (*Student) TableName() string {
	return "student"
}

/*
CREATE TABLE `article` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `title` longtext DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_user_articles` (`user_id`),
  CONSTRAINT `fk_user_articles` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
*/

// 官方文档把一对多分为了两类, Belongs To 和 Has Many，这里合并在一起测试
type UserM struct {
	ID       uint       `gorm:"primaryKey"`
	Name     string     `gorm:"column:name"`
	Articles []ArticleM `gorm:"foreignKey:UserID"` // 用户拥有的文章列表
}

func (*UserM) TableName() string {
	return "user"
}

type ArticleM struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"column:title"`
	UserID uint   // 属于
	User   UserM  `gorm:"foreignKey:UserID"` // 属于
}

func (*ArticleM) TableName() string {
	return "article"
}

// -----------------------------------------------------------------------

/*
CREATE TABLE `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` longtext DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_user_name` (`user_name`) USING HASH
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `role` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_name` (`name`) USING HASH
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `user_role` (
  `role_id` bigint(20) unsigned NOT NULL,
  `user_id` bigint(20) unsigned NOT NULL,
  PRIMARY KEY (`role_id`,`user_id`),
  KEY `fk_user` (`user_id`),
  CONSTRAINT `fk_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`),
  CONSTRAINT `fk_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
*/

type UserModel struct {
	ID       uint        `gorm:"primaryKey"`
	UserName string      `gorm:"column:user_name"`
	Roles    []RoleModel `gorm:"many2many:user_role;joinForeignKey:UserID;joinReferences:RoleID"`
}

func (*UserModel) TableName() string {
	return "user"
}

type RoleModel struct {
	ID    uint        `gorm:"primaryKey"`
	Name  string      `gorm:"column:name"`
	Users []UserModel `gorm:"many2many:user_role;joinForeignKey:RoleID;joinReferences:UserID"`
}

func (*RoleModel) TableName() string {
	return "role"
}

type UserRoleModel struct {
	UserID uint `gorm:"cloumn:user_id,primaryKey"`
	RoleID uint `gorm:"cloumn:role_id,primaryKey"`
}

func (*UserRoleModel) TableName() string {
	return "user_role"
}
