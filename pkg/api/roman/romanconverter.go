package roman

// Interface for the Roman Convertor
type RomanConverter interface {
	Convert(num int) (string, error)
}
