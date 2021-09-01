package api

import (
	"github.com/labstack/echo/v4"
	"github.com/pkpgdjk/echo-todo-app/database"
	"github.com/pkpgdjk/echo-todo-app/model"
	"net/http"
)

func GetMember(c echo.Context) (err error)  {
	db := database.New()
	var users []model.User

	if result := db.Preload("AssignedTodos").Find(&users); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Can't get user list.")
	}

	return c.JSON(http.StatusCreated, users)
}
