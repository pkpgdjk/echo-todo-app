package model

import (
	"github.com/pkpgdjk/echo-todo-app/utils"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username   			string `gorm:"unique_index;not null"`
	Token   			string
	Todo 				[]*Todo	`gorm:"foreignKey:OwnerId"`
	AssignedTodos		[]*Todo `gorm:"many2many:user_todos;"`
}

func NewUser(username string) *User {
	return &User{
		Username: username,
		Token: utils.GenerateUserToken(12),
	}
}