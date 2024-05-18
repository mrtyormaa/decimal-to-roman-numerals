package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/middleware"
)

func TestSecurityMiddleware(t *testing.T) {
	// Create a new Gin engine
	router := gin.New()

	// Use the Security middleware
	router.Use(middleware.Security())

	// Define a test route
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "Test route")
	})

	// Create a mock request to the test route
	req, err := http.NewRequest("GET", "/test", nil)
	assert.NoError(t, err, "Failed to create request")

	// Create a mock response recorder
	resp := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(resp, req)

	// Check if the response headers are set correctly
	assert.Equal(t, "max-age=315360000; includeSubdomains", resp.Header().Get("Strict-Transport-Security"), "Strict-Transport-Security header not set correctly")
	assert.Equal(t, "DENY", resp.Header().Get("X-Frame-Options"), "X-Frame-Options header not set correctly")
	assert.Equal(t, "nosniff", resp.Header().Get("X-Content-Type-Options"), "X-Content-Type-Options header not set correctly")
	assert.Equal(t, "1; mode=block", resp.Header().Get("X-XSS-Protection"), "X-XSS-Protection header not set correctly")
	assert.Equal(t, "default-src 'self'", resp.Header().Get("Content-Security-Policy"), "Content-Security-Policy header not set correctly")
	assert.Equal(t, "noopen", resp.Header().Get("X-Download-Options"), "X-Download-Options header not set correctly")
	assert.Equal(t, "strict-origin-when-cross-origin", resp.Header().Get("Referrer-Policy"), "Referrer-Policy header not set correctly")

	// Check if the response status code is OK
	assert.Equal(t, http.StatusOK, resp.Code, "Unexpected response status code")
}
