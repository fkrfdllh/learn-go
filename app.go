package main

import "fmt"

func main() {
	var investmentAmount = 1000
	var expectedReturnRate = 5.5 
	var years = 10

	var futureValue = float64(investmentAmount) * (1 + expectedReturnRate / 100)
}