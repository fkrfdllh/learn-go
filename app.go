package main

import "fmt"

func main() {
	revenue := userPrompt("Revenue: ")
	expenses := userPrompt("Expenses: ")
	taxRate := userPrompt("Tax Rate: ")

	ebt, profit, ratio := calculateProfitValues(revenue, expenses, taxRate)

	formatOutput("EBT: ", ebt)
	formatOutput("Profit: ", profit)
	formatOutput("Ratio: ", ratio)
}

func userPrompt(prompt string) (userScan float64) {
	fmt.Print(prompt)
	fmt.Scanln(&userScan)

	return
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