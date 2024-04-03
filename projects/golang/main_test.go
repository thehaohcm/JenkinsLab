package main

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/labstack/echo/v4"
    "github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
    // Create a test server with the same Echo instance used in main
    e := echo.New()
    // ... Add middleware and routes as in main
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    rec := httptest.NewRecorder()

    // Test the root path (/)
    e.ServeHTTP(rec, req)
    assert.Equal(t, http.StatusOK, rec.Code)
    assert.Equal(t, "Hello, Docker! <3.", rec.Body.String())

    // Test the health path (/health)
    req = httptest.NewRequest(http.MethodGet, "/health", nil)
    rec = httptest.NewRecorder()
    e.ServeHTTP(rec, req)
    assert.Equal(t, http.StatusOK, rec.Code)
    assert.Equal(t, `{"Status":"OK"}`, rec.Body.String())
}

