package routes

import (
	"user-api/controllers"

	"github.com/labstack/echo/v4"
)

func UserRole(e *echo.Echo) {

	e.POST("/user", controllers.CreateUser)
	e.GET("user/:userId", controllers.GetUser)
	e.PUT("user/:userId", controllers.EditUser)
	e.DELETE("user/:userId", controllers.DeleteUser)
	e.GET("/users", controllers.GetAllUsers)
}
