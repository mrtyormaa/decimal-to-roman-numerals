package roman

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mrtyormaa/decimal-to-roman-numerals/pkg/models"
)

const (
	LowerLimit = 1
	UpperLimit = 3999
)

var converter RomanConverter = &BasicRomanConverter{}

// @BasePath /

func Healthcheck(g *gin.Context) {
	response := gin.H{
		"status":  "success",
		"message": "Decimal to Roman Numerals Converter",
	}
	g.JSON(http.StatusOK, response)
}

// ConvertNumbersToRoman handles the API request to convert numbers to Roman numerals.
// @Summary Convert numbers to Roman numerals
// @Description Convert a comma-separated list of numbers to their corresponding Roman numeral representations.
// @ID convertNumbersToRoman
// @Accept json
// @Produce json
// @Param numbers query string true "Comma-separated list of integers to be converted"
// @Success 200 {object} []models.RomanNumeral
// @Router /convert [get]
func ConvertNumbersToRoman(c *gin.Context) {
	// Get all query parameters
	queryParams := c.Request.URL.Query()

	// Check if there are any query parameters other than 'numbers'
	for param := range queryParams {
		if param != "numbers" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Only 'numbers' query parameter is allowed"})
			return
		}
	}

	// Get the numbers parameters from the query string
	numbersParams := c.QueryArray("numbers")

	// Check if the numbers parameter is missing
	if len(numbersParams) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The 'numbers' query parameter is required"})
		return
	}

	// Parse and validate the number list
	numbers, invalidNumbers := ParseNumberList(numbersParams)

	// If there are any invalid numbers, return an error response
	if len(invalidNumbers) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":           fmt.Sprintf("Invalid input. Please provide valid integers within the supported range (%d-%d).", LowerLimit, UpperLimit),
			"invalid_numbers": invalidNumbers,
		})
		return
	}

	// Convert the numbers to Roman numerals
	results := ConvertNumbersToRomanNumerals(numbers)

	// Return the results as a JSON response
	c.JSON(http.StatusOK, gin.H{"results": results})
}

// ParseNumberList parses and validates an array of comma-separated list of numbers
func ParseNumberList(numbersParams []string) ([]int, []string) {
	var numbers []int
	var invalidNumbers []string

	// Iterate over each numbers parameter
	for _, numbersParam := range numbersParams {
		numberStrings := strings.Split(numbersParam, ",")
		for _, numberString := range numberStrings {
			// Trim spaces
			numberString = strings.TrimSpace(numberString)
			if numberString == "" {
				invalidNumbers = append(invalidNumbers, "")
				continue // Skip empty strings
			}
			number, err := strconv.Atoi(numberString)
			if err != nil || number < LowerLimit || number > UpperLimit {
				invalidNumbers = append(invalidNumbers, numberString)
			} else {
				numbers = append(numbers, number)
			}
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

// ConvertRangesToRoman handles the API request to convert ranges of numbers to Roman numerals.
// @Summary Convert ranges of numbers to Roman numerals
// @Description Convert multiple ranges of numbers to their corresponding Roman numeral representations.
// @ID convertRangesToRoman
// @Accept json
// @Produce json
// @Param input body models.RangesPayload true "Array of number ranges"
// @Success 200 {object} []models.RomanNumeral
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
		if r.Min < LowerLimit || r.Max > UpperLimit || r.Min > r.Max {
			return nil, fmt.Errorf("invalid range. each range must be within %d to %d and min should not be greater than max", LowerLimit, UpperLimit)
		}
		for i := r.Min; i <= r.Max; i++ {
			numbers = append(numbers, i)
		}
	}

	return numbers, nil
}
