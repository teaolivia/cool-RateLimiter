package main

import (
	"cool-RateLimiter/controllers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	e.GET("/unlimited", controllers.GetUnlimited)
	e.GET("/limited", controllers.GetLimited)
}
