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
	var investmentAmount = 1000
	var expectedReturnRate = 5.5 
	var years = 10

	var futureValue = float64(investmentAmount) * math.Pow(1 + expectedReturnRate / 100, float64(years))

	fmt.Println(futureValue)
}