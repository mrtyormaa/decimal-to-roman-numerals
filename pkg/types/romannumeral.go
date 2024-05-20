package types

// RomanNumeral struct defines the response model for the converted numbers.
type RomanNumeral struct {
	Decimal uint   `json:"number"`
	Roman   string `json:"roman"`
}
