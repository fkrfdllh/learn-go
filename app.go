package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFilePath = "balance.txt"

func readBalance() (float64, error) {
	data, err := os.ReadFile(balanceFilePath)

	if err != nil {
		return 0, errors.New("failed to retrieve balance")
	}
	
	balanceStr := string(data)
	balance, err := strconv.ParseFloat(balanceStr, 64)

	if err != nil {
		return 0, errors.New("failed to parse balance")
	}

	return balance, nil
}

func writeToFile(balance float64) {
	latestBalance := fmt.Sprint(balance)
	os.WriteFile(balanceFilePath, []byte(latestBalance), 0644)
}

func main() {
	// Go just have one keyword to loop, that is "for" loop 
	
	fmt.Println("Welcome to GoBank!")
	
	// We can use like "while", for with condition using
	// for someCondition { // TODO }
	
	for {
		accountBalance, err := readBalance()

		if err != nil {
			fmt.Println("--------------------")
			fmt.Printf("Error: %v\n", err)
			fmt.Println("--------------------")

			panic("Can't continue, sorry")
		}
		
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
				writeToFile(accountBalance)
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
				writeToFile(accountBalance)
			default:
				fmt.Println("Goodbye!")
				fmt.Println("Thank you for choosing our bank")
				return
		}
	}
}