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

func GetUserById(userId int64) User {
	var user User
	config.Db.Table("users").Where("id = ?", userId).First(&user)
	return user
}

func InsertUser(user User) {
	config.Db.Table("users").Create(&user)
}

func GetUserByUsernameAndPassword(username string, password string) (User, error) {
	var user User
	err := config.Db.Table("users").Where("username = ? and password = ?", username, password).First(&user)
	if err.Error != nil {
		return User{}, err.Error
	}
	return user, nil
}
