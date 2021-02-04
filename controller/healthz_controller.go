package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetInit(e *echo.Echo) {
	e.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, "Rest API started")
	})
}

func res(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

func resErr(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, err.Error())
}
