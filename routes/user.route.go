package routes

import (
	"crud/controller"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	e.POST("/user", controller.CreateUser)
	e.GET("/user/:id", controller.GetAUser)
}
