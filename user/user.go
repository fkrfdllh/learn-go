package user

import (
	"errors"
	"fmt"
	"time"
)

// // Struct Literal notation
// var appUser user = user{
// 	firstName: firstName,
// 	lastName: lastName,
// 	birthDate: birthDate,
// 	createdAt: time.Now(),
// }

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User User
}

// Since Go isn't OOP programming language
// While creating constructor, it's follows the pattern 
// func new<structName>
// and then returns pointer
func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name, or birth date can not be null")
	}
	
	return &User{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) (*Admin) {
	return &Admin{
		email: email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName: "Admin",
			birthDate: "---",
		},
	}
}

func (user *User) OutputUserData() {
	fmt.Println(user.firstName, user.lastName, user.birthDate)
}

func (user *User) ClearName() {
	user.firstName = ""
	user.lastName = ""
}