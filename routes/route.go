package routes

import (
	"net/http"

	"test-echo/controllers"
	"test-echo/kreditPlusController"
	mid "test-echo/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	g := e.Group("/api/v1")
	g.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, TEST")
	})
	g.GET("/test", controllers.GetTest, mid.IsAuthenticated)

	//
	g.GET("/gen-token", controllers.GenerateToken)
	g.GET("/check-limits", kreditPlusController.GetLimit, mid.IsAuthenticated)
	g.GET("/get-customer", kreditPlusController.GetCustomer, mid.IsAuthenticated)

	return e
}
