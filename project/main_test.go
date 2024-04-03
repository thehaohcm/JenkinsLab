package main_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	e        *echo.Echo
	recorder *httptest.ResponseRecorder
)

func TestMain(m *testing.M) {
	e = echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	recorder = httptest.NewRecorder()

	os.Exit(m.Run())
}

func TestRootRoute(t *testing.T) {
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3.")
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	e.ServeHTTP(recorder, req)

	result := recorder.Result()
	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, result.StatusCode)
	}

	expectedBody := "Hello, Docker! <3."
	actualBody, err := io.ReadAll(result.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %v", err)
	}

	if !bytes.Equal([]byte(expectedBody), actualBody) {
		t.Errorf("Expected body '%s' but got '%s'", expectedBody, string(actualBody))
	}
}
