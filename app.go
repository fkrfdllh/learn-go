package main

// To import another package(s) is
// Use the name of the module (see go.mod) then the name of package
// In this project, I use "learn-go/utils"

import (
	"fmt"

	"learn-go/utils"

	"github.com/Pallinder/go-randomdata"
)

const balanceFilePath = "balance.txt"

func main() {
	// Go just have one keyword to loop, that is "for" loop 
	
	fmt.Println("Welcome to GoBank!")
	fmt.Println("Reach us 24/7 on", randomdata.PhoneNumber())
	
	// We can use like "while", for with condition using
	// for someCondition { // TODO }
	
	for {
		accountBalance, err := utils.ReadFloatValueFromFile(balanceFilePath)

		if err != nil {
			fmt.Println("--------------------")
			fmt.Printf("Error: %v\n", err)
			fmt.Println("--------------------")

			panic("Can't continue, sorry")
		}
	
		showOptions()
		
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
				utils.WriteFloatToFile(balanceFilePath, accountBalance)
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
				utils.WriteFloatToFile(balanceFilePath, accountBalance)
			default:
				fmt.Println("Goodbye!")
				fmt.Println("Thank you for choosing our bank")
				return
		}
	}
}