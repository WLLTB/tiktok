package repository

import (
	"log"
	"tiktok/app/config"
	. "tiktok/app/schema"
)

// GetTableUserList 获取全部 user 对象
func GetTableUserList() ([]User, error) {
	var userList []User
	if err := config.Db.Table("users").Find(&userList).Error; err != nil {
		log.Println(err.Error())
		return userList, err
	}
	return userList, nil
}
