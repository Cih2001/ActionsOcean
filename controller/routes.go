package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func InitializeRoutes(e *echo.Echo) {

	//Initializing routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World from github actions!")
	})

}
