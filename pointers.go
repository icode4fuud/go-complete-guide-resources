package main

import "fmt"

func main() {
	age := 32 //regular variable

	agePointer := &age //pointer variable

	fmt.Println("Value of age is: ", *agePointer)

	getAdultYears(agePointer)
	fmt.Println(age)
}

func getAdultYears(age *int) {
	//return *age - 18
	*age = *age - 18
}
