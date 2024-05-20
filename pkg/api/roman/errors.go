package roman

import "fmt"

// Error codes and messages map
var ErrorMap = map[string]string{
	CodeInvalidParam:             "only 'numbers' query parameter is allowed",
	CodeMissingNumbersParam:      "the 'numbers' query parameter is required",
	CodeInvalidInput:             fmt.Sprintf("invalid input: please provide valid integers within the supported range (%d-%d)", LowerLimit, UpperLimit),
	CodeOutOfBounds:              fmt.Sprintf("input out of bounds, must be between %d and %d", LowerLimit, UpperLimit),
	CodeFailedReadBody:           "failed to read request body",
	CodeInvalidJSONPayload:       "invalid JSON payload",
	CodeInvalidRangesPayload:     "invalid JSON payload: expected only 'ranges' key with an array value",
	CodeEmptyRanges:              "empty 'ranges': provide valid min and max values for the ranges",
	CodeInvalidRange:             fmt.Sprintf("invalid range: each range must be within %d to %d and min should not be greater than max", LowerLimit, UpperLimit),
	CodeInvalidJSONDuplicateKeys: "invalid JSON payload: duplicate `ranges` keys",
	CodeQueryParamInPostRequest:  "invalid request: query parameters not allowed in POST requests",
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
