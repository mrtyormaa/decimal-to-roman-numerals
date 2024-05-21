package roman

import "fmt"

// Error codes and messages map
var ErrorMap = map[string]string{
	CodeInvalidParam:              "only 'numbers' query parameter is allowed",
	CodeMissingNumbersParam:       "'numbers' query parameter is required",
	CodeInvalidInput:              fmt.Sprintf("invalid input: please provide valid integers within the supported range (%d-%d)", LowerLimit, UpperLimit),
	CodeOutOfBounds:               fmt.Sprintf("input out of bounds, must be between %d and %d", LowerLimit, UpperLimit),
	CodeFailedReadBody:            "failed to read request body",
	CodeInvalidRangeJSON:          "invalid JSON: expected 'ranges' key with an array value. Array of 'min' and 'max'. ex. {'ranges': [{'min': 1, 'max': 2}]}",
	CodeInvalidJSONDuplicateKeys:  "invalid JSON payload: duplicate `ranges` keys",
	CodeQueryParamInPostRequest:   "invalid request: query parameters not allowed in POST requests",
	CodeInvalidRangeMinMoreMax:    "invalid ranges: 'min' should be less than 'max'",
	CodeInvalidRangeBounds:        fmt.Sprintf("invalid ranges: 'min' and 'max' values must be within %d to %d", LowerLimit, UpperLimit),
	CodeInValidJSON:               "failed to parse JSON",
	CodeInValidRangeMissingMinMax: "invalid format: each range must have 'min' and 'max' integers",
}

// AppError represents a structured error with a code and message
type AppError struct {
	Code    string
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// NewAppError creates a new AppError given an error code
func NewAppError(code string) *AppError {
	message, exists := ErrorMap[code]
	if !exists {
		message = "unknown error"
	}
	return &AppError{Code: code, Message: message}
}
