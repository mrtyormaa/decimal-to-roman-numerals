package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/api"
)

// Constants for API version and base path
const (
	APIVersion = "/api/v1"
	BasePath   = APIVersion + "/convert"
)

// SetupLoadRouter sets up the Gin router for testing
func SetupLoadRouter() *gin.Engine {
	return api.InitRouter()
}

// Helper function to perform a POST request and return the response recorder
func performPostRequest(router *gin.Engine, url string, payload interface{}) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	jsonPayload, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	return w
}

// Helper function to perform a GET request and return the response recorder
func performRequest(router *gin.Engine, url string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)
	return w
}
