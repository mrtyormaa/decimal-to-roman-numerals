package models

// NumberRange struct defines the model for a range of numbers
type NumberRange struct {
	Min int `json:"min" binding:"required" example:"10"` // The minimum value of the range (inclusive).
	Max int `json:"max" binding:"required" example:"20"` // The maximum value of the range (inclusive).
}

type RangesPayload struct {
	Ranges []NumberRange `json:"ranges" binding:"required"`
}
