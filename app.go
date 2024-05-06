package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scanln(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scanln(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scanln(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate / 100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}