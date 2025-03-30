package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stwrtrio/movie-festival-api/config"
	"github.com/stwrtrio/movie-festival-api/internal/models"
	"github.com/stwrtrio/movie-festival-api/pkg/utils"
	"github.com/stwrtrio/movie-festival-api/pkg/validator"
)

func RegisterHandler(c echo.Context) error {
	var user models.User
	var req models.UserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := validator.ValidateStruct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}

	user.ID = utils.GenerateUUID()
	user.Email = req.Email
	user.PasswordHash = hashedPassword

	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}

func LoginHandler(c echo.Context) error {
	var req models.UserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := validator.ValidateStruct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var user models.User
	result := config.DB.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
