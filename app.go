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

func (user *user) outputUserData() {
	fmt.Println(user.firstName, user.lastName, user.birthDate)
}

func (user *user) clearName() {
	user.firstName = ""
	user.lastName = ""
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birth date (YYYY-MM-DD): ")

	// TODO

	// Struct Literal notation
	var appUser user = user{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}

	// outputUserData(&appUser)
	appUser.outputUserData()
	appUser.clearName()
	appUser.outputUserData()
}

func getUserData(promptText string) string {
	var value string	
	
	fmt.Print(promptText)
	fmt.Scanln(&value)
	
	return value
}

// func outputUserData(user *user) {
// 	// // Proper way
// 	// fmt.Println((*user).firstName, (*user).lastName, (*user).birthDate)
	
// 	// Shorter way
// 	fmt.Println(user.firstName, user.lastName, user.birthDate)
// }