package roman

import (
	"errors"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/models"
)

var converter RomanConverter = &BasicRomanConverter{}

// @BasePath /

func Healthcheck(g *gin.Context) {
	g.JSON(http.StatusOK, "Decimal to Roman Numeral Converter")
}

// convert godoc
// @Summary Get Roman Numeral
// @Description Get the roman numeral equivalent for a given decimal(s) in ascending order
// @Tags romans
// @Produce json
// @Success 200 {object} models.RomanNumeral "Successfully retrieved Roman Numerals"
// @Router /convert [get]
func ConvertNumbersToRoman(c *gin.Context) {
	// Get the 'numbers' query parameter
	numbersParam := c.Query("numbers")
	if numbersParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The 'numbers' query parameter is required."})
		return
	}

	// Parse and validate the number list
	numbers, invalidNumbers := ParseNumberList(numbersParam)

	// If there are any invalid numbers, return an error response
	if len(invalidNumbers) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":           "Invalid input. Please provide valid integers within the supported range (1-3999).",
			"invalid_numbers": invalidNumbers,
		})
		return
	}

	// Convert the numbers to Roman numerals
	results := ConvertNumbersToRomanNumerals(numbers)

	// Return the results as a JSON response
	c.JSON(http.StatusOK, gin.H{"results": results})
}

// ParseNumberList parses and validates a comma-separated list of numbers
func ParseNumberList(numbersParam string) ([]int, []string) {
	numberStrings := strings.Split(numbersParam, ",")
	var numbers []int
	var invalidNumbers []string

	for _, numberString := range numberStrings {
		number, err := strconv.Atoi(strings.TrimSpace(numberString))
		if err != nil || number < 1 || number > 3999 {
			invalidNumbers = append(invalidNumbers, numberString)
		} else {
			numbers = append(numbers, number)
		}
	}

	return numbers, invalidNumbers
}

// ConvertNumbersToRomanNumerals converts a list of unique numbers to their Roman numeral equivalents
func ConvertNumbersToRomanNumerals(numbers []int) []models.RomanNumeral {
	uniqueNumbers := make(map[int]struct{})
	for _, number := range numbers {
		uniqueNumbers[number] = struct{}{}
	}

	var results []models.RomanNumeral
	for number := range uniqueNumbers {
		roman, _ := converter.Convert(number)
		results = append(results, models.RomanNumeral{
			Decimal: uint(number),
			Roman:   roman,
		})
	}

	// Sort the results by decimal value in ascending order
	sort.Slice(results, func(i, j int) bool {
		return results[i].Decimal < results[j].Decimal
	})

	return results
}

// convert godoc
// @Summary Get Roman Numerals for Ranges of Numbers
// @Description Get the roman numeral equivalent for given ranges in ascending order
// @Tags romans
// @Produce json
// @Success 200 {object} models.RomanNumeral "Successfully retrieved Roman Numerals"
// @Router /convert [post]
func ConvertRangesToRoman(c *gin.Context) {
	var payload models.RangesPayload

	// Bind the JSON payload to the ranges variable
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload."})
		return
	}

	// Process the ranges to generate a list of numbers
	numbers, err := ProcessRanges(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert the numbers to Roman numerals
	results := ConvertNumbersToRomanNumerals(numbers)

	// Return the results as a JSON response
	c.JSON(http.StatusOK, gin.H{"results": results})
}

// ProcessRanges processes the ranges and generates a list of numbers
func ProcessRanges(payload models.RangesPayload) ([]int, error) {
	var numbers []int

	for _, r := range payload.Ranges {
		if r.Min < 1 || r.Max > 3999 || r.Min > r.Max {
			return nil, errors.New("invalid range. each range must be within 1 to 3999 and min should not be greater than max")
		}
		for i := r.Min; i <= r.Max; i++ {
			numbers = append(numbers, i)
		}
	}

	return numbers, nil
}
