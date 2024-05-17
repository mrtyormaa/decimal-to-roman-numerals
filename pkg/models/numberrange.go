package models

// NumberRange struct defines the model for a range of numbers
type NumberRange struct {
	Min int `json:"min" binding:"required"`
	Max int `json:"max" binding:"required"`
}

type RangesPayload struct {
	Ranges []NumberRange `json:"ranges" binding:"required"`
}
