package roman

const (
	LowerLimit = 1
	UpperLimit = 3999

	// Error messages
	ErrInvalidParam         = "only 'numbers' query parameter is allowed"
	ErrMissingNumbersParam  = "the 'numbers' query parameter is required"
	ErrInvalidInput         = "invalid input: please provide valid integers within the supported range (%d-%d)"
	ErrFailedReadBody       = "failed to read request body"
	ErrInvalidJSONPayload   = "invalid JSON payload"
	ErrInvalidRangesPayload = "invalid JSON payload: expected only 'ranges' key with an array value"
	ErrEmptyRanges          = "empty 'ranges': provide valid min and max values for the ranges"
	ErrInvalidRange         = "invalid range: each range must be within %d to %d and min should not be greater than max"
)
