package controllers

import (
	"fmt"
	"helloworld/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetQuotes(c echo.Context) error {
	page_str := c.QueryParam("page")
	page, err := strconv.Atoi(page_str)
	if err != nil {
		page = 0
	}
	result := models.GetQuotes(page)
	return c.JSON(http.StatusOK, result)
}

func GetQuote(c echo.Context) error {
	quote := quoteByID(c)
	return c.JSON(http.StatusOK, quote)
}

func AddQuote(c echo.Context) error {
	quote := new(models.Quote)
	if err := c.Bind(quote); err != nil {
		return c.JSON(http.StatusOK, err)
	}
	models.AddQuote(quote)

	return c.JSON(http.StatusCreated, quote)
}

func UpdateQuote(c echo.Context) error {
	quote := new(models.Quote)
	if err := c.Bind(quote); err != nil {
		return c.JSON(http.StatusOK, err)
	}
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	uodated_quote := models.UpdateQuote(quote, id)
	return c.JSON(http.StatusCreated, uodated_quote)
}

func quoteByID(c echo.Context) models.Quote {
	quote_id := c.Param("id")
	id, err := strconv.Atoi(quote_id)
	if err != nil {
		id = 1
	}
	quote := models.GetQuote(id)
	return quote
}
