package middleware

import (
	echo "github.com/labstack/echo/v4"
	"github.com/pkpgdjk/echo-todo-app/database"
	"github.com/pkpgdjk/echo-todo-app/model"
	"net/http"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		db := database.New()
		var user model.User
		result := db.Where("token = ?", c.Request().Header.Get("X-Member-Token")).First(&user)
		if result.Error != nil && result.RowsAffected == 0 {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		c.Set("claim", &user)
		if err := next(c); err != nil {
			return err;
		}
		return nil
	}
}
