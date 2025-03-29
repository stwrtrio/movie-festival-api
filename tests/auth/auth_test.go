package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stwrtrio/movie-festival-api/config"
	"github.com/stwrtrio/movie-festival-api/internal/handlers"
	"github.com/stwrtrio/movie-festival-api/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Load .env file
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Run tests
	code := m.Run()
	os.Exit(code)
}

func TestRegisterFailed(t *testing.T) {
	e := echo.New()
	config.InitDB()

	// Mock request body
	requestData := map[string]string{
		"email":    " ",
		"password": "wrongpassword",
	}
	body, _ := json.Marshal(requestData)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	err := handlers.RegisterHandler(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestRegisterSuccess(t *testing.T) {
	e := echo.New()
	config.InitDB()

	// Mock request body
	requestData := map[string]string{
		"email":    utils.GenerateRandomString(10) + "@email.com",
		"password": "testpassword",
	}
	body, _ := json.Marshal(requestData)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	err := handlers.RegisterHandler(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestLoginFailed(t *testing.T) {
	e := echo.New()
	config.InitDB()

	// Mock request body
	requestData := map[string]string{
		"email":    "test@email.com",
		"password": "",
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

func TestLoginSuccess(t *testing.T) {
	e := echo.New()
	config.InitDB()

	// Mock request body
	requestData := map[string]string{
		"email":    "test@email.com",
		"password": "testpassword",
	}
	body, _ := json.Marshal(requestData)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	err := handlers.LoginHandler(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}
