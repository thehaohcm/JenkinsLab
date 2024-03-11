package main
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert" // You might need to install testify - go get -u github.com/stretchr/testify
)

func TestHelloWorld(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Register the route we want to test
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	// Create a request object
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Record the response
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// Assert that the response code is OK
	assert.Equal(t, http.StatusOK, rec.Code)

	// Assert that the response body contains the expected message
	expectedBody := "Hello, Docker! <3"
	responseBody := rec.Body.String()
	assert.Contains(t, responseBody, expectedBody)
}

func TestHealthCheck(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Register the route we want to test
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	// Create a request object
	req := httptest.NewRequest(http.MethodGet, "/health", nil)

	// Record the response
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// Assert that the response code is OK
	assert.Equal(t, http.StatusOK, rec.Code)

	// Decode the response body
	var healthResponse map[string]string
	err := json.NewDecoder(rec.Body).Decode(&healthResponse)
	assert.NoError(t, err)

	// Assert that the status is "OK"
	assert.Equal(t, "OK", healthResponse["Status"])
}
