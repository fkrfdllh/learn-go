package main

import "fmt"

// -- How to define pointer --
// var variable *int
// or
// varPointer := &variable

// If pointer pointing null value
// Then it is "nil"

func main() {
	age := 24

	var agePointer *int = &age

	// fmt.Println("Age:", age)

	// this line is show where pointer pointing the value memory address
	// fmt.Println("Age:", agePointer)

	// this line is show the value of pointer's pointing
	fmt.Println("Age:", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}