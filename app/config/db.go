package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func InitGormDb() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	dbs, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Panicln("数据库连接失败")
	}
	log.Print("数据库连接成功")
	return dbs
}
