package repository

import (
	"tiktok/app/config"
	"tiktok/app/constant"
	. "tiktok/app/schema"
)

// GetAllUsers 获取全部 user 对象
func GetAllUsers() []User {
	var userList []User
	config.Db.Table(constant.USER).Find(&userList)
	return userList
}

func GetUserById(userId int64) User {
	var user User
	config.Db.Table(constant.USER).Where("id = ?", userId).First(&user)
	return user
}

func InsertUser(user User) {
	config.Db.Table(constant.USER).Create(&user)
}

func GetUserByUsernameAndPassword(username string, password string) (User, error) {
	var user User
	err := config.Db.Table(constant.USER).Where("username = ? and password = ?", username, password).First(&user)
	if err.Error != nil {
		return User{}, err.Error
	}
	return user, nil
}
