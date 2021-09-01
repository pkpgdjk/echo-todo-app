package router

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/pkpgdjk/echo-todo-app/api"
	"github.com/pkpgdjk/echo-todo-app/config"
	"github.com/pkpgdjk/echo-todo-app/handler"
	"github.com/pkpgdjk/echo-todo-app/middleware"
)

func Init() *echo.Echo {

	e := echo.New()

	// Set Bundle MiddleWare
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Gzip())
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	e.HTTPErrorHandler = handler.JSONHTTPErrorHandler

	// Routes
	v1 := e.Group("/api")
	{
		v1.POST("/register", api.Register, echoMiddleware.RateLimiterWithConfig(config.GetRateLimiterConfig()))

		protected := v1.Group("/todo")
		{
			protected.Use(middleware.Auth)

			protected.GET("/", api.GetTodo)
			protected.GET("/member", api.GetMember)
			protected.POST("/", api.CreateTodo, echoMiddleware.RateLimiterWithConfig(config.GetRateLimiterConfig()))
			protected.PUT("/:id/assign", api.AssignMember)
			protected.PUT("/:id/un-assign", api.UnAssignMember)
			protected.PUT("/:id", api.UpdateTodo)
			protected.PATCH("/:id", api.UpdateTodoStatus)
		}
	}
	return e
}