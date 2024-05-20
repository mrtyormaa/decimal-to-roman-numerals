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
func SetupRouter() *gin.Engine {
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

// Helper function to provide different Edge test cases to be tested for the Ranges Handler
func getRangesValidTestCases() []struct {
	name           string
	payload        interface{}
	expectedStatus int
	expectedResult []struct {
		Number int
		Roman  string
	}
} {
	return []struct {
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
}

// Helper function to provide different Invalid test cases to be tested for the Ranges Handler
func getRangesInvalidTestCases() []struct {
	name           string
	payload        interface{}
	expectedStatus int
} {
	return []struct {
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
}

// Helper function to provide different Edge test cases to be tested for the Ranges Handler
func getRangesEdgeTestCases() []struct {
	name           string
	payload        interface{}
	expectedStatus int
	expectedResult []struct {
		Number int
		Roman  string
	}
} {
	return []struct {
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
}
