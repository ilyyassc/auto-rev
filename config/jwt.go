package config

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

const secret = "secret"

func SetJwt(e *echo.Echo) *echo.Group {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	jwtGroup := e.Group("/api")
	jwtGroup.Use(middleware.JWT([]byte(secret)))
	return jwtGroup
}

func GenerateToken(username string) (string, error) {
	defer CatchErrorGeneral()

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //sejam

	// Generate encoded token
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func GetClaims(e echo.Context, datas ...string) map[string]interface{} {
	defer CatchErrorGeneral()

	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	var result = map[string]interface{}{}
	for _, val := range datas {
		result[val] = claims[val]
	}
	return result
}
