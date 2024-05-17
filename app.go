package main

import (
	"errors"
	"fmt"
)

func main() {
	revenue, expenses, taxRate, err := userPrompt()

	if err != nil {
		errMessage := fmt.Sprintf("Error: %v\n", err)

		panic(errMessage)
	}

	ebt, profit, ratio := calculateProfitValues(revenue, expenses, taxRate)

	formatOutput("EBT: ", ebt)
	formatOutput("Profit: ", profit)
	formatOutput("Ratio: ", ratio)
}

func userPrompt() (float64, float64, float64, error) {
	var revenue, expenses, taxRate float64
	
	fmt.Print("Revenue: ")
	fmt.Scanln(&revenue)

	if revenue < 1 {
		return 0, 0, 0, errors.New("revenue value cannot be null")
	}

	fmt.Print("Expenses: ")
	fmt.Scanln(&expenses)

	if expenses < 1 {
		return 0, 0, 0, errors.New("expenses value cannot be null")
	}

	fmt.Print("Tax Rate: ")
	fmt.Scanln(&taxRate)

	if taxRate < 1 {
		return 0, 0, 0, errors.New("tax rate value cannot be null")
	}

	return revenue, expenses, taxRate, nil
}

func calculateProfitValues(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate / 100)
	ratio = ebt / profit

	return
}

func formatOutput(prefix string, value float64) {
	fmt.Printf(prefix + "%.2f\n", value)
}