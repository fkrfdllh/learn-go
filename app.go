package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birth date (YYYY-MM-DD): ")

	// TODO

	fmt.Println(firstName, lastName, birthDate)
}

func getUserData(promptText string) string {
	var value string	
	
	fmt.Print(promptText)
	fmt.Scanln(&value)
	
	return value
}