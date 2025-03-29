package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stwrtrio/movie-festival-api/config"
	"github.com/stwrtrio/movie-festival-api/internal/handlers"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	e := echo.New()
	config.InitDB()

	// Mock request body
	requestData := map[string]string{
		"email":    "test@email.com",
		"password": "wrongpassword",
	}
	body, _ := json.Marshal(requestData)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	err := handlers.RegisterHandler(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestLoginHandler(t *testing.T) {
	e := echo.New()
	config.InitDB()

	// Mock request body
	requestData := map[string]string{
		"email":    "test@email.com",
		"password": "wrongpassword",
	}
	body, _ := json.Marshal(requestData)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	err := handlers.LoginHandler(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
