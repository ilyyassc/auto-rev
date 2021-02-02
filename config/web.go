package config

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitWeb() *echo.Echo {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			if check := validateToken(req.Header.Get("Authorization")); !check {
				return c.JSON(http.StatusUnauthorized, "Invalid Token")
			}
			return next(c)
		}
	})
	return e
}

func validateToken(token string) bool {
	defer CatchErrorGeneral()
	log.Println("========================> token :", token)
	//validate token via grpc
	return true
}
