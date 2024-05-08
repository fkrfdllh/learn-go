package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to GoBank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice uint8
	fmt.Print("Your choice: ")
	fmt.Scanln(&choice)

	if choice == 1 {
		fmt.Printf("Your account balance is $%.2f\n", accountBalance)
	}

	fmt.Println("Your choice is ", choice)
}