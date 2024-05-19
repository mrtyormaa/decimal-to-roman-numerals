package test

import (
	"encoding/json"
	"net/http"
	"sync"
	"testing"
)

// Load test for GET /convert endpoint
// Perform total 1000 requests, distributed among the goroutines
// Check the status code and validate the response body
func TestConvertHandlerLoad(t *testing.T) {
	router := SetupRouter()
	numRequests := 1000
	concurrency := 10
	var wg sync.WaitGroup
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numRequests/concurrency; j++ {
				w := performRequest(router, BasePath+"?numbers=123")
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
