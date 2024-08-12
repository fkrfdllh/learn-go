package main

import (
	"fmt"
	"learn-go/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birth date (YYYY-MM-DD): ")

	var appUser *user.User

	appUser, err := user.NewUser(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println(err)
		return 
	}

	admin := user.NewAdmin("fkrfdllh@gmail.com", "18januari")
	admin.User.OutputUserData()
	admin.User.ClearName()
	admin.User.OutputUserData()

	// outputUserData(&appUser)
	appUser.OutputUserData()
	appUser.ClearName()
	appUser.OutputUserData()
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