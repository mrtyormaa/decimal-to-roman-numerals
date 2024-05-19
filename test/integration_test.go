package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
)

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

// Helper function to unmarshal the response and check the results for the POST /convert endpoint
func checkPostResponse(t *testing.T, w *httptest.ResponseRecorder, expectedResults []struct {
	Number int
	Roman  string
}) {
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

	for i, result := range response.Results {
		if result.Number != expectedResults[i].Number || result.Roman != expectedResults[i].Roman {
			t.Errorf("handler returned unexpected result at index %d: got {number: %d, roman: %s} want {number: %d, roman: %s}", i, result.Number, result.Roman, expectedResults[i].Number, expectedResults[i].Roman)
		}
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

// Test cases for POST /convert endpoint with valid inputs
func TestConvertRangesHandlerValid(t *testing.T) {
	router := SetupRouter()

	testCases := []struct {
		name           string
		payload        interface{}
		expectedStatus int
		expectedResult []struct {
			Number int
			Roman  string
		}
	}{
		{
			name: "MultipleRanges",
			payload: struct {
				Ranges []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				} `json:"ranges"`
			}{
				Ranges: []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				}{
					{Min: 10, Max: 12},
					{Min: 14, Max: 16},
				},
			},
			expectedStatus: http.StatusOK,
			expectedResult: []struct {
				Number int
				Roman  string
			}{
				{10, "X"},
				{11, "XI"},
				{12, "XII"},
				{14, "XIV"},
				{15, "XV"},
				{16, "XVI"},
			},
		},
		{
			name: "OverlappingRanges",
			payload: struct {
				Ranges []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				} `json:"ranges"`
			}{
				Ranges: []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				}{
					{Min: 5, Max: 10},
					{Min: 8, Max: 12},
				},
			},
			expectedStatus: http.StatusOK,
			expectedResult: []struct {
				Number int
				Roman  string
			}{
				{5, "V"},
				{6, "VI"},
				{7, "VII"},
				{8, "VIII"},
				{9, "IX"},
				{10, "X"},
				{11, "XI"},
				{12, "XII"},
			},
		},
		{
			name: "OutOfOrderRanges",
			payload: struct {
				Ranges []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				} `json:"ranges"`
			}{
				Ranges: []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				}{
					{Min: 15, Max: 18},
					{Min: 10, Max: 12},
				},
			},
			expectedStatus: http.StatusOK,
			expectedResult: []struct {
				Number int
				Roman  string
			}{
				{10, "X"},
				{11, "XI"},
				{12, "XII"},
				{15, "XV"},
				{16, "XVI"},
				{17, "XVII"},
				{18, "XVIII"},
			},
		},
		{
			name: "BoundaryValues",
			payload: struct {
				Ranges []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				} `json:"ranges"`
			}{
				Ranges: []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				}{
					{Min: 1, Max: 1},
					{Min: 3999, Max: 3999},
				},
			},
			expectedStatus: http.StatusOK,
			expectedResult: []struct {
				Number int
				Roman  string
			}{
				{1, "I"},
				{3999, "MMMCMXCIX"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := performPostRequest(router, BasePath, tc.payload)
			checkStatus(t, w, tc.expectedStatus)
			checkPostResponse(t, w, tc.expectedResult)
		})
	}
}

// Test cases for POST /convert endpoint with invalid inputs
func TestConvertRangesHandlerInvalid(t *testing.T) {
	router := SetupRouter()

	testCases := []struct {
		name           string
		payload        interface{}
		expectedStatus int
	}{
		{
			name: "EmptyRanges",
			payload: struct {
				Ranges []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				} `json:"ranges"`
			}{
				Ranges: []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				}{},
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "NonIntegerValues",
			payload: struct {
				Ranges []struct {
					Min string `json:"min"`
					Max string `json:"max"`
				} `json:"ranges"`
			}{
				Ranges: []struct {
					Min string `json:"min"`
					Max string `json:"max"`
				}{
					{Min: "a", Max: "z"},
				},
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "NegativeValues",
			payload: struct {
				Ranges []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				} `json:"ranges"`
			}{
				Ranges: []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				}{
					{Min: -1, Max: 5},
				},
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "MaxLessThanMin",
			payload: struct {
				Ranges []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				} `json:"ranges"`
			}{
				Ranges: []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				}{
					{Min: 10, Max: 5},
				},
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := performPostRequest(router, BasePath, tc.payload)
			checkStatus(t, w, tc.expectedStatus)
		})
	}
}

// Test cases for edge cases for POST /convert endpoint
func TestConvertRangesHandlerEdgeCases(t *testing.T) {
	router := SetupRouter()

	testCases := []struct {
		name           string
		payload        interface{}
		expectedStatus int
		expectedResult []struct {
			Number int
			Roman  string
		}
	}{
		{
			name: "SingleNumberRange",
			payload: struct {
				Ranges []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				} `json:"ranges"`
			}{
				Ranges: []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				}{
					{Min: 50, Max: 50},
				},
			},
			expectedStatus: http.StatusOK,
			expectedResult: []struct {
				Number int
				Roman  string
			}{
				{50, "L"},
			},
		},
		{
			name: "VerySmallRange",
			payload: struct {
				Ranges []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				} `json:"ranges"`
			}{
				Ranges: []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				}{
					{Min: 101, Max: 102},
				},
			},
			expectedStatus: http.StatusOK,
			expectedResult: []struct {
				Number int
				Roman  string
			}{
				{101, "CI"},
				{102, "CII"},
			},
		},
		{
			name: "MaxValidRange",
			payload: struct {
				Ranges []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				} `json:"ranges"`
			}{
				Ranges: []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				}{
					{Min: 1, Max: 3999},
				},
			},
			expectedStatus: http.StatusOK,
			expectedResult: []struct {
				Number int
				Roman  string
			}{
				{1, "I"},
				{3999, "MMMCMXCIX"},
			},
		},
		{
			name: "OverlappingLargeRange",
			payload: struct {
				Ranges []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				} `json:"ranges"`
			}{
				Ranges: []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				}{
					{Min: 10, Max: 20},
					{Min: 15, Max: 25},
				},
			},
			expectedStatus: http.StatusOK,
			expectedResult: []struct {
				Number int
				Roman  string
			}{
				{10, "X"},
				{11, "XI"},
				{12, "XII"},
				{13, "XIII"},
				{14, "XIV"},
				{15, "XV"},
				{16, "XVI"},
				{17, "XVII"},
				{18, "XVIII"},
				{19, "XIX"},
				{20, "XX"},
				{21, "XXI"},
				{22, "XXII"},
				{23, "XXIII"},
				{24, "XXIV"},
				{25, "XXV"},
			},
		},
		{
			name: "ReverseOrderRanges",
			payload: struct {
				Ranges []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				} `json:"ranges"`
			}{
				Ranges: []struct {
					Min int `json:"min"`
					Max int `json:"max"`
				}{
					{Min: 20, Max: 25},
					{Min: 10, Max: 15},
				},
			},
			expectedStatus: http.StatusOK,
			expectedResult: []struct {
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
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := performPostRequest(router, BasePath, tc.payload)
			checkStatus(t, w, tc.expectedStatus)

			var response struct {
				Results []struct {
					Number int    `json:"number"`
					Roman  string `json:"roman"`
				} `json:"results"`
			}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("failed to unmarshal response: %v", err)
				return
			}

			if len(response.Results) != len(tc.expectedResult) && tc.name != "MaxValidRange" {
				t.Errorf("handler returned unexpected number of results: got %v want %v", len(response.Results), len(tc.expectedResult))
			}

			if tc.name == "MaxValidRange" {
				if len(response.Results) != 3999 {
					t.Errorf("handler returned unexpected number of results: got %v want %v", len(response.Results), 3999)
				}
				if response.Results[0].Number != 1 || response.Results[0].Roman != "I" {
					t.Errorf("handler returned unexpected first result: got {number: %d, roman: %s} want {number: 1, roman: I}", response.Results[0].Number, response.Results[0].Roman)
				}
				if response.Results[3998].Number != 3999 || response.Results[3998].Roman != "MMMCMXCIX" {
					t.Errorf("handler returned unexpected last result: got {number: %d, roman: %s} want {number: 3999, roman: MMMCMXCIX}", response.Results[3998].Number, response.Results[3998].Roman)
				}
			} else {
				for i, result := range response.Results {
					if result.Number != tc.expectedResult[i].Number || result.Roman != tc.expectedResult[i].Roman {
						t.Errorf("handler returned unexpected result at index %d: got {number: %d, roman: %s} want {number: %d, roman: %s}", i, result.Number, result.Roman, tc.expectedResult[i].Number, tc.expectedResult[i].Roman)
					}
				}
			}
		})
	}
}
