package api

import (
	"fmt"
	"github.com/pkpgdjk/echo-todo-app/database"
	"github.com/pkpgdjk/echo-todo-app/dto"
	"github.com/pkpgdjk/echo-todo-app/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Register(c echo.Context) (err error)  {
	db := database.New()
	var registerDto dto.RegisterDto
	if err := (&echo.DefaultBinder{}).BindBody(c, &registerDto); err != nil {
		return err
	}

	user := model.NewUser(registerDto.Username)
	if result := db.Create(&user); result.Error != nil {
		fmt.Print(result.Error)
		return echo.NewHTTPError(http.StatusInternalServerError, "Save data into database error.")
	}

	return c.JSON(http.StatusCreated, user)
}
