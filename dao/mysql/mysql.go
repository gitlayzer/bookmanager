package mysql

import (
	"bookmanager/model"
	"fmt"

	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义全局是因为在其他的包内会使用
var DB *gorm.DB

func InitMysql() {
	dsn := "root:123456@tcp(10.0.0.17:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("初始化MySQL失败", err)
	}
	DB = db

	if err := DB.AutoMigrate(model.Book{}, model.User{}); err != nil {
		fmt.Println("自动创建表结构失败", err)
	}
}
