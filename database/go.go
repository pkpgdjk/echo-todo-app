package database

import (
	"github.com/pkpgdjk/echo-todo-app/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	dsn := "root:123123@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("storage err: ", err)
	}

	AutoMigrate(db)
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Todo{})
}