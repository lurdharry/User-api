package main

import (
	"user-api/configs"
	"user-api/routes"

	"github.com/labstack/echo/v4"
)

func main()  {
	e := echo.New()


	md := &configs.MongoDriver{}

	// Create Mongo connection
	md.Connect()
	// Close connection after main
   defer md.Disconnect()
	

   //routes
   routes.UserRole(e) 




	e.GET("/", func(c echo.Context) error {
		return  c.JSON(200, echo.Map{"data": "Hello from Echo & mongoDB"})
	})

	e.Logger.Fatal(e.Start(":6000"))
}