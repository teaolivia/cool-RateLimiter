package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUnlimited(c echo.Context) error {
	return c.String(http.StatusOK, "Unlimited! Let's Go!")
}

func GetLimited(c echo.Context) error {
	return c.String(http.StatusOK, "Limited, don't over use me!")
}
