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

	var agePointer *int

	agePointer = &age

	// fmt.Println("Age:", age)

	// this line is show where pointer pointing the value memory address
	// fmt.Println("Age:", agePointer)

	// this line is show the value of pointer's pointing
	fmt.Println("Age:", *agePointer)

	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}