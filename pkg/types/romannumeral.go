package types

// RomanNumeral struct defines the response model for the converted numbers.
type RomanNumeral struct {
	Decimal uint   `json:"number" example:"100"`
	Roman   string `json:"roman" example:"C"`
}

// RomanNumeralResponse represents a successful response containing Roman numerals.
type RomanNumeralResponse struct {
	Results []RomanNumeral `json:"results"`
}

// ErrorResponse represents an error response with an error message and optional invalid numbers.
type ErrorResponse struct {
	Error          string   `json:"error" example:"[ERR1002] invalid input: please provide valid integers within the supported range (1-3999)"`
	InvalidNumbers []string `json:"invalid_numbers,omitempty" example:"['8888']"`
}

// ErrorResponse represents an error response with an error message and optional invalid numbers.
type JsonErrorResponse struct {
	Error string `json:"error" example:"[ERR1005] invalid JSON: JSON must contain only the 'ranges' key, which should be an array of one or more objects with 'min' and 'max' values. 'min' and 'max' values must be within 1 to 3999, and 'min' should not be greater than 'max'. No other keys are allowed."`
}

// HealthResponse represents the health status of the service.
type HealthResponse struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"Decimal to Roman Numerals Converter"`
}
