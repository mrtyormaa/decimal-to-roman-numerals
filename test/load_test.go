package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/api"
)

// SetupLoadRouter sets up the Gin router for testing
func SetupLoadRouter() *gin.Engine {
	return api.InitRouter()
}

// Helper function to perform a POST request and return the response recorder
func performLoadTestPostRequest(router *gin.Engine, url string, payload interface{}) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	jsonPayload, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	return w
}

// Helper function to perform a GET request and return the response recorder
func performLoadTestRequest(router *gin.Engine, url string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(w, req)
	return w
}

// Load test for GET /convert endpoint
// Perform total 1000 requests, distributed among the goroutines
// Check the status code and validate the response body
func TestConvertHandlerLoad(t *testing.T) {
	router := SetupLoadRouter()
	numRequests := 1000
	concurrency := 10
	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numRequests/concurrency; j++ {
				w := performLoadTestRequest(router, BasePath+"?numbers=123")
				if w.Code != http.StatusOK {
					t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
				}
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
				expectedResults := []struct {
					Number int
					Roman  string
				}{
					{123, "CXXIII"},
				}
				if len(response.Results) != len(expectedResults) {
					t.Errorf("handler returned unexpected number of results: got %v want %v", len(response.Results), len(expectedResults))
				}
				for k, result := range response.Results {
					if result.Number != expectedResults[k].Number || result.Roman != expectedResults[k].Roman {
						t.Errorf("handler returned unexpected result at index %d: got {number: %d, roman: %s} want {number: %d, roman: %s}", k, result.Number, result.Roman, expectedResults[k].Number, expectedResults[k].Roman)
					}
				}
			}
		}()
	}

	wg.Wait()
}

// Load test for POST /convert endpoint
// Perform total 1000 requests, distributed among the goroutines
// Check the status code and validate the response body
func TestConvertRangesHandlerLoad(t *testing.T) {
	router := SetupLoadRouter()
	numRequests := 1000
	concurrency := 10
	var wg sync.WaitGroup
	wg.Add(concurrency)

	payload := struct {
		Ranges []struct {
			Min int `json:"min"`
			Max int `json:"max"`
		} `json:"ranges"`
	}{
		Ranges: []struct {
			Min int `json:"min"`
			Max int `json:"max"`
		}{
			{Min: 10, Max: 15},
			{Min: 20, Max: 25},
		},
	}

	expectedResults := []struct {
		Number int
		Roman  string
	}{
		{10, "X"},
		{11, "XI"},
		{12, "XII"},
		{13, "XIII"},
		{14, "XIV"},
		{15, "XV"},
		{20, "XX"},
		{21, "XXI"},
		{22, "XXII"},
		{23, "XXIII"},
		{24, "XXIV"},
		{25, "XXV"},
	}

	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numRequests/concurrency; j++ {
				w := performLoadTestPostRequest(router, BasePath, payload)
				if w.Code != http.StatusOK {
					t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
				}
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
				if len(response.Results) != len(expectedResults) {
					t.Errorf("handler returned unexpected number of results: got %v want %v", len(response.Results), len(expectedResults))
				}
				for k, result := range response.Results {
					if result.Number != expectedResults[k].Number || result.Roman != expectedResults[k].Roman {
						t.Errorf("handler returned unexpected result at index %d: got {number: %d, roman: %s} want {number: %d, roman: %s}", k, result.Number, result.Roman, expectedResults[k].Number, expectedResults[k].Roman)
					}
				}
			}
		}()
	}

	wg.Wait()
}
