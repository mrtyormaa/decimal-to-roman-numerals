package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/api"
)

// Constants for API version and base path
const (
	APIVersion = "/api/v1"
	BasePath   = APIVersion + "/convert"
)

// SetupRouter sets up the Gin router for testing
func SetupRouter() *gin.Engine {
	return api.InitRouter()
}

// Helper function to perform a GET request and return the response recorder
func performRequest(router *gin.Engine, url string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)
	return w
}

// Helper function to check the response status code
func checkStatus(t *testing.T, w *httptest.ResponseRecorder, expectedStatus int) {
	if status := w.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v", status, expectedStatus)
	}
}

// Helper function to unmarshal the response and check the results
func checkResponse(t *testing.T, w *httptest.ResponseRecorder, number int, expected string) {
	var response struct {
		Results []struct {
			Number int    `json:"number"`
			Roman  string `json:"roman"`
		} `json:"results"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}

	if len(response.Results) != 1 || response.Results[0].Number != number || response.Results[0].Roman != expected {
		t.Errorf("handler returned unexpected body: got %v want {number: %d, roman: %s}", w.Body.String(), number, expected)
	}
}

// Test cases for valid inputs for GET /api/v1/convert
func TestConvertHandlerValid(t *testing.T) {
	router := SetupRouter()
	testCases := []struct {
		number   int
		expected string
	}{
		{1, "I"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
		{3999, "MMMCMXCIX"},
	}

	for _, tc := range testCases {
		t.Run("Valid_"+strconv.Itoa(tc.number), func(t *testing.T) {
			w := performRequest(router, BasePath+"?numbers="+strconv.Itoa(tc.number))
			checkStatus(t, w, http.StatusOK)
			checkResponse(t, w, tc.number, tc.expected)
		})
	}
}

// Test cases for invalid inputs for GET /api/v1/convert
func TestConvertHandlerInvalid(t *testing.T) {
	router := SetupRouter()
	testCases := []string{
		"abc",
		"-1",
		"4000",
	}

	for _, tc := range testCases {
		t.Run("Invalid_"+tc, func(t *testing.T) {
			w := performRequest(router, BasePath+"?numbers="+tc)
			checkStatus(t, w, http.StatusBadRequest)
		})
	}
}

// Test cases for edge cases for GET /api/v1/convert
func TestConvertHandlerEdgeCases(t *testing.T) {
	router := SetupRouter()
	testCases := []struct {
		number   int
		expected string
	}{
		{1, "I"},
		{3999, "MMMCMXCIX"},
	}

	for _, tc := range testCases {
		t.Run("EdgeCase_"+strconv.Itoa(tc.number), func(t *testing.T) {
			w := performRequest(router, BasePath+"?numbers="+strconv.Itoa(tc.number))
			checkStatus(t, w, http.StatusOK)
			checkResponse(t, w, tc.number, tc.expected)
		})
	}
}

// Performance test to check the handler under load for GET /api/v1/convert
func TestConvertHandlerPerformance(t *testing.T) {
	router := SetupRouter()
	for i := 0; i < 1000; i++ {
		t.Run("LoadTest_"+strconv.Itoa(i), func(t *testing.T) {
			w := performRequest(router, BasePath+"?numbers=123")
			checkStatus(t, w, http.StatusOK)
			checkResponse(t, w, 123, "CXXIII")
		})
	}
}
