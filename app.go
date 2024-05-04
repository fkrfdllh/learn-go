package main

import (
	"fmt"
	"math"
)

/**
* Falsy condition (Null values)
* int => 0
* float => 0.0
* string => ""
* bool => false
 */

func main() {
	// investmentAmount, expectedReturnRate, years := 1000.0, 5.5, 10.0
	investmentAmount := 1000.0
	expectedReturnRate := 5.5
	years := 10.0

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)

	fmt.Println(futureValue)
}