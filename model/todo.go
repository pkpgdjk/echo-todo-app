package model

import (
	"github.com/pkpgdjk/echo-todo-app/dto"
	"github.com/jinzhu/gorm"
	"github.com/pkpgdjk/echo-todo-app/enum"
)

type Todo struct {
	gorm.Model
	Title   		string `gorm:"unique_index;not null"`
	Description   	string
	Status   		enum.TodoStatus
	Assigned		[]*User `gorm:"many2many:user_todos;"`
	OwnerId      	uint
	Owner      		*User `gorm:"foreignKey:OwnerId;"`
}

func NewTodo(createTodoDto dto.CreateTodoDto, owner User) *Todo {
	return &Todo{
		Title: createTodoDto.Title,
		Description: createTodoDto.Description,
		OwnerId: owner.ID,
	}
}