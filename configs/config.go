package configs

import (
	"prakerja5/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(){
	dsn := "root:123ABC4d.@tcp(127.0.0.1:3306)/prakerja5?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database error")
	}
	migration()
}

func migration(){
	DB.AutoMigrate(&models.News{})
}