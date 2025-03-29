package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/stwrtrio/movie-festival-api/internal/models"
)

func RegisterHandler(c echo.Context) error {
	var user models.UserRequest
	if err := c.Bind(&user); err != nil {
		return err
	}

	return nil
}

func LoginHandler(c echo.Context) error {
	var userInput models.UserRequest
	if err := c.Bind(&userInput); err != nil {
		return err
	}

	return nil
}
