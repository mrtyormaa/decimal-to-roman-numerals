package roman

// Interface for the Roman Converter
// Converts an integer to its corresponding Roman numeral string.
type RomanConverter interface {
	Convert(num int) (string, error)
}
