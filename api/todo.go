package api

import (
	"fmt"
	"github.com/pkpgdjk/echo-todo-app/database"
	"github.com/pkpgdjk/echo-todo-app/dto"
	"github.com/pkpgdjk/echo-todo-app/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetTodo(c echo.Context) (err error)  {
	db := database.New()
	var todos model.Todo

	if result := db.Preload("Owner").Preload("Assigned").Find(&todos); result.Error != nil {
		fmt.Print(result.Error)
		return echo.NewHTTPError(http.StatusInternalServerError, "Can't get todo list.")
	}

	return c.JSON(http.StatusCreated, todos)
}

func CreateTodo(c echo.Context) (err error)  {
	db := database.New()
	user := c.Get("claim").(*model.User)

	var createTodoDto dto.CreateTodoDto
	if err := (&echo.DefaultBinder{}).BindBody(c, &createTodoDto); err != nil {
		return err
	}

	todo := model.NewTodo(createTodoDto, *user)
	if result := db.Create(&todo); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Save todo into database.")
	}

	return c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c echo.Context) (err error)  {
	db := database.New()

	var updateTodoDto dto.UpdateTodoDto
	if err := (&echo.DefaultBinder{}).Bind(&updateTodoDto, c); err != nil {
		return err
	}
	var todo model.Todo
	result := db.Where("ID = ?", updateTodoDto.ID).First(&todo)
	if result.Error != nil || result.RowsAffected == 0  {
		return echo.NewHTTPError(http.StatusBadRequest, "Can't find todo.")
	}

	todo.Title = updateTodoDto.Title
	todo.Description = updateTodoDto.Description

	if result := db.Save(&todo); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Can't save todo into database.")
	}

	return c.JSON(http.StatusCreated, todo)
}


func UpdateTodoStatus(c echo.Context) (err error)  {
	db := database.New()

	var updateTodoStatusDto dto.UpdateTodoStatusDto
	if err := (&echo.DefaultBinder{}).Bind(&updateTodoStatusDto, c); err != nil {
		return err
	}
	var todo model.Todo
	result := db.Where("ID = ?", updateTodoStatusDto.ID).First(&todo)
	if result.Error != nil || result.RowsAffected == 0  {
		return echo.NewHTTPError(http.StatusBadRequest, "Can't find todo.")
	}

	todo.Status = updateTodoStatusDto.Status
	if result := db.Save(&todo); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Can't save todo into database.")
	}

	return c.JSON(http.StatusCreated, todo)
}

func AssignMember(c echo.Context) (err error)  {
	db := database.New()
	var assignTodoDto dto.AssignTodoDto
	if err := (&echo.DefaultBinder{}).Bind(&assignTodoDto, c); err != nil {
		return err
	}
	var todo model.Todo
	result := db.Where("ID = ?", assignTodoDto.ID).First(&todo)
	if result.Error != nil || result.RowsAffected == 0  {
		return echo.NewHTTPError(http.StatusBadRequest, "Can't find todo.")
	}

	var user model.User
	result = db.Where("ID = ?", assignTodoDto.UserId).First(&user)
	if result.Error != nil || result.RowsAffected == 0  {
		return echo.NewHTTPError(http.StatusBadRequest, "Can't find user.")
	}

	if err := db.Model(&todo).Association("Assigned").Append([]model.User{user}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Save todo into database.")
	}

	return c.JSON(http.StatusCreated, todo)
}

func UnAssignMember(c echo.Context) (err error)  {
	db := database.New()
	var assignTodoDto dto.AssignTodoDto
	if err := (&echo.DefaultBinder{}).Bind(&assignTodoDto, c); err != nil {
		return err
	}
	var todo model.Todo
	result := db.Where("ID = ?", assignTodoDto.ID).First(&todo)
	if result.Error != nil || result.RowsAffected == 0  {
		return echo.NewHTTPError(http.StatusBadRequest, "Can't find todo.")
	}

	var user model.User
	result = db.Where("ID = ?", assignTodoDto.UserId).First(&user)
	if result.Error != nil || result.RowsAffected == 0  {
		return echo.NewHTTPError(http.StatusBadRequest, "Can't find user.")
	}

	if err := db.Model(&todo).Association("Assigned").Delete(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Can't save todo into database.")
	}

	return c.JSON(http.StatusCreated, todo)
}
