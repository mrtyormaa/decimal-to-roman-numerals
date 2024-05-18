package roman

import (
	"fmt"
)

type BasicRomanConverter struct{}

func (c *BasicRomanConverter) Convert(num int) (string, error) {
	if num < LowerLimit || num > UpperLimit {
		return "", fmt.Errorf("input out of bounds, must be between %d and %d", LowerLimit, UpperLimit)
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
