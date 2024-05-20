package main

import "fmt"

// -- How to define pointer --
// var variable *int
// or
// varPointer := &variable

func main() {
	age := 24

	var agePointer *int

	agePointer = &age

	fmt.Println("Age:", age)

	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
}

func getAdultYears(age int) int {
	return age - 18
}