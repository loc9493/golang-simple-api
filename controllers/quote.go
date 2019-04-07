package controllers

import (
	"helloworld/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetQuotes(c echo.Context) error {
	result := models.GetQuotes()
	println("foo")
	return c.JSON(http.StatusOK, result)
}
