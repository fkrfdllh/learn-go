package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	// Go just have one keyword to loop, that is "for" loop 
	
	fmt.Println("Welcome to GoBank!")
	
	// We can use like "while", for with condition using
	// for someCondition { // TODO }
	
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
	
		var choice uint8
		fmt.Print("Your choice: ")
		fmt.Scanln(&choice)

		switch choice {
			case 1:
				fmt.Printf("Your account balance is $%.2f\n", accountBalance)
			case 2:
				var depositAmount float64
			
				fmt.Print("Your deposit: $")
				fmt.Scanln(&depositAmount)
		
				if depositAmount <= 0 {
					fmt.Println("Invalid amount. Deposit amount must be greater than 0.")
					continue
				}
		
				accountBalance += depositAmount
		
				fmt.Printf("Your updated account balance is $%.2f\n", accountBalance)
			case 3:
				var withdrawAmount float64
	
				fmt.Print("Withdrawal amount: $")
				fmt.Scanln(&withdrawAmount)
		
				if withdrawAmount <= 0 {
					fmt.Println("Invalid amount. Withdrawal amount must be greater than 0.")
					continue
				}
		
				if withdrawAmount > accountBalance {
					fmt.Println("Invalid amount. Withdrawal amount must less than your account balance.")
					continue
				}
				
				accountBalance -= withdrawAmount
		
				fmt.Printf("Your updated account balance is $%.2f\n", accountBalance)
			default:
				fmt.Println("Goodbye!")
				fmt.Println("Thank you for choosing our bank")
				return
		}
	}
}