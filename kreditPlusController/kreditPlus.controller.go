package kreditPlusController

import (
	"net/http"
	"test-echo/services/kreditPlusService"

	"github.com/labstack/echo/v4"
)

func GetLimit(c echo.Context) error {
	result, _ := kreditPlusService.GetLimit(c)

	return c.JSON(http.StatusOK, result)
}

func GetCustomer(c echo.Context) error {
	result, _ := kreditPlusService.GetCustomer(c)

	return c.JSON(http.StatusOK, result)
}
