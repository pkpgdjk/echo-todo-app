package dto

import "github.com/pkpgdjk/echo-todo-app/enum"

type CreateTodoDto struct {
	Title string `form:"title" json:"title" `
	Description string `form:"description" json:"description" `
}

type UpdateTodoDto struct {
	ID string `param:"id"`
	Title string `form:"title" json:"title" `
	Description string `form:"description" json:"description" `
}

type UpdateTodoStatusDto struct {
	ID string `param:"id"`
	Status enum.TodoStatus `form:"status" json:"status" `
}


type AssignTodoDto struct {
	ID string `param:"id"`
	UserId enum.TodoStatus `form:"userId" json:"userId" `
}
