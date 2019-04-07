package controllers

import (
	"helloworld/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetQuotes(c echo.Context) error {
	result := models.GetQuotes()
	return c.JSON(http.StatusOK, result)
}

func AddQuote(c echo.Context) error {
	quote := new(models.Quote)
	if err := c.Bind(quote); err != nil {
		return c.JSON(http.StatusOK, err)
	}
	err := models.AddQuote(quote)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error)
	}
	return c.JSON(http.StatusCreated, quote)
}
