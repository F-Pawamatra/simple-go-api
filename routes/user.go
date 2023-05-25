package routes

import (
	"github.com/f-pawamatra/rest-api/controller"
	"github.com/labstack/echo"
)

func UserRoutes(e *echo.Echo) {
	e.GET("/", controller.Hello)
	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users", controller.SaveUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)
}