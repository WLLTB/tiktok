package repository

import (
	"tiktok/app/config"
	. "tiktok/app/schema"
)

// GetAllUsers 获取全部 user 对象
func GetAllUsers() []User {
	var userList []User
	config.Db.Table("users").Find(&userList)
	return userList
}

func GetUserById(userId int) User {
	var user User
	config.Db.Table("users").Where("id = ?", userId).First(&user)
	return user
}
