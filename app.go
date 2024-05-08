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
	} else if choice == 2 {
		var depositAmount float64
		
		fmt.Print("Your deposit: $")
		fmt.Scanln(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid amount. Deposit amount must be greater than 0.")
			return
		}

		accountBalance += depositAmount

		fmt.Printf("Your updated account balance is $%.2f\n", accountBalance)
	} else if choice == 3 {
		var withdrawAmount float64

		fmt.Print("Withdrawal amount: $")
		fmt.Scanln(&withdrawAmount)

		if withdrawAmount <= 0 {
			fmt.Println("Invalid amount. Withdrawal amount must be greater than 0.")
			return
		}

		if withdrawAmount > accountBalance {
			fmt.Println("Invalid amount. Withdrawal amount must less than your account balance.")
			return
		}
		
		accountBalance -= withdrawAmount

		fmt.Printf("Your updated account balance is $%.2f\n", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}
}