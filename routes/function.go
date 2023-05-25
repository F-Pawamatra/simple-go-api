package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RoutesInit() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	UserRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}