package main

import (
	"user-api/configs"
	"user-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Create Mongo connection
	configs.Connect()
	// Close connection after main
	//    defer md.Disconnect()

	//routes
	routes.UserRole(e)

	e.Logger.Fatal(e.Start(":6000"))
}
