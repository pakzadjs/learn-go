package main

import "fmt"

func main() {
	age := 32 //regular variable

	// var agePointer *int
	// agePointer = &age

	agePointer := &age

	fmt.Println("Age:", *agePointer)

	editAgeToAdultYears(&age)
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
