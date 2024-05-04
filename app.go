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
	const inflationRate float64 = 2.5
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Investment Amount: ")
	fmt.Scanln(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scanln(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scanln(&years)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}