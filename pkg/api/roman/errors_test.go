package roman

import (
	"testing"
)

func TestNewAppError(t *testing.T) {
	tests := []struct {
		name         string
		code         string
		expectedCode string
		expectedMsg  string
	}{
		{
			name:         "InvalidParam",
			code:         CodeInvalidParam,
			expectedCode: CodeInvalidParam,
			expectedMsg:  "only 'numbers' query parameter is allowed",
		},
		{
			name:         "MissingNumbersParam",
			code:         CodeMissingNumbersParam,
			expectedCode: CodeMissingNumbersParam,
			expectedMsg:  "the 'numbers' query parameter is required",
		},
		{
			name:         "InvalidInput",
			code:         CodeInvalidInput,
			expectedCode: CodeInvalidInput,
			expectedMsg:  "invalid input: please provide valid integers within the supported range (1-3999)",
		},
		{
			name:         "OutOfBounds",
			code:         CodeOutOfBounds,
			expectedCode: CodeOutOfBounds,
			expectedMsg:  "input out of bounds, must be between 1 and 3999",
		},
		{
			name:         "FailedReadBody",
			code:         CodeFailedReadBody,
			expectedCode: CodeFailedReadBody,
			expectedMsg:  "failed to read request body",
		},
		{
			name:         "InvalidJSONPayload",
			code:         CodeInvalidJSONPayload,
			expectedCode: CodeInvalidJSONPayload,
			expectedMsg:  "invalid JSON payload",
		},
		{
			name:         "InvalidRangesPayload",
			code:         CodeInvalidRangesPayload,
			expectedCode: CodeInvalidRangesPayload,
			expectedMsg:  "invalid JSON payload: expected only 'ranges' key with an array value",
		},
		{
			name:         "EmptyRanges",
			code:         CodeEmptyRanges,
			expectedCode: CodeEmptyRanges,
			expectedMsg:  "empty 'ranges': provide valid min and max values for the ranges",
		},
		{
			name:         "InvalidRange",
			code:         CodeInvalidRange,
			expectedCode: CodeInvalidRange,
			expectedMsg:  "invalid range: each range must be within 1 to 3999 and min should not be greater than max",
		},
		{
			name:         "InvalidJSONDuplicateKeys",
			code:         CodeInvalidJSONDuplicateKeys,
			expectedCode: CodeInvalidJSONDuplicateKeys,
			expectedMsg:  "invalid JSON payload: duplicate `ranges` keys",
		},
		{
			name:         "QueryParamInPostRequest",
			code:         CodeQueryParamInPostRequest,
			expectedCode: CodeQueryParamInPostRequest,
			expectedMsg:  "invalid request: query parameters not allowed in POST requests",
		},
		{
			name:         "UnknownErrorCode",
			code:         "UNKNOWN_CODE",
			expectedCode: "UNKNOWN_CODE",
			expectedMsg:  "unknown error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewAppError(tt.code)
			if err.Code != tt.expectedCode {
				t.Errorf("expected code %s, got %s", tt.expectedCode, err.Code)
			}
			if err.Message != tt.expectedMsg {
				t.Errorf("expected message %s, got %s", tt.expectedMsg, err.Message)
			}
		})
	}
}

func TestAppError_Error(t *testing.T) {
	err := &AppError{Code: "TEST_CODE", Message: "This is a test error message"}
	expected := "[TEST_CODE] This is a test error message"
	if err.Error() != expected {
		t.Errorf("expected %s, got %s", expected, err.Error())
	}
}
