package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"test-mnc/util"
)

func AuthenticationHandler(c echo.Context) error {
	tokenString := c.Request().Header.Get("Authorization")
	if tokenString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Missing authorization header"})
	}

	tokenString = tokenString[len("Bearer "):]

	err := util.VerifyToken(tokenString)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Unauthenticated"})
	}
	return nil
}
