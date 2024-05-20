package roman

// Converts an integer to its corresponding Roman numeral string.
type BasicRomanConverter struct{}

// Convert converts an integer to its corresponding Roman numeral string.
// It first checks if the input number is within the acceptable range
// (LowerLimit to UpperLimit). If the number is out of bounds, it returns
// an error. The function uses predefined arrays of decimal values and
// their corresponding Roman numeral symbols to iteratively build the
// Roman numeral string. It subtracts values from the input number
// while appending the corresponding symbols to the result string until
// the number is reduced to zero, ensuring an accurate conversion.
func (c *BasicRomanConverter) Convert(num int) (string, error) {
	if num < LowerLimit || num > UpperLimit {
		return "", NewAppError(CodeOutOfBounds)
	}

	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syb := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			roman += syb[i]
			num -= val[i]
		}
	}
	return roman, nil
}
