package main

import (
	"crud/config"
	"github.com/labstack/echo/v4"
	"crud/routes"
)

func main() {

	
	app := echo.New()
	
	// connect to database
	config.ConnectDB()

	// add routes
	routes.UserRoutes(app)


	// lunch app
	app.Logger.Fatal(app.Start(":8080"))
}
