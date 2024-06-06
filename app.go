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

// Since Go isn't OOP programming language
// While creating constructor, it's follows the pattern 
// func new<structName>
// and then returns pointer
func newUser(firstName, lastName, birthDate string) *user {
	return &user{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
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

	// // Struct Literal notation
	// var appUser user = user{
	// 	firstName: firstName,
	// 	lastName: lastName,
	// 	birthDate: birthDate,
	// 	createdAt: time.Now(),
	// }

	appUser := newUser(firstName, lastName, birthDate)

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