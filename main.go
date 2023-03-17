package main

import (
	"github.com/labstack/echo/v4"
)

func main() {

	app := echo.New()
	app.GET("/", func(c echo.Context) error {
		return c.JSON(200, &echo.Map{"data": "Hello from Echo & mongoDB"})
	})

	app.Logger.Fatal(app.Start(":8080"))
}
