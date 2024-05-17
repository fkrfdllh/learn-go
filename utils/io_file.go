package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// To use these functions outside of this package
// The first character of function's name must be uppercase
// It means the function is exported (public function)

func ReadFloatValueFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return 0, errors.New("file does not exists")
	}
	
	valueStr := string(data)
	value, err := strconv.ParseFloat(valueStr, 64)

	if err != nil {
		return 0, errors.New("failed to parse value")
	}

	return value, nil
}

func WriteFloatToFile(filename string, balance float64) {
	valueText := fmt.Sprint(balance)
	os.WriteFile(filename, []byte(valueText), 0644)
}