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
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5 
	var years float64 = 10

	var futureValue = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)

	fmt.Println(futureValue)
}