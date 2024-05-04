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
	const inflationRate float64 = 6.5
	investmentAmount := 1000.0
	expectedReturnRate := 5.5
	years := 10.0

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}