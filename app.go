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

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scanln(&investmentAmount)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scanln(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scanln(&years)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)
	
	// fmt.Println("Future Value: ", futureValue)
	// fmt.Println("Future Value (adjusted for inflation): ", futureRealValue)
	// fmt.Printf("Future Value: %.2f\nFuture Value (adjusted for inflation): %.2f\n", futureValue, futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}