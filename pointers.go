package main

import "fmt"

func main() {
	age := 32 //regular variable

	agePointer := &age //pointer variable

	fmt.Println("Value of age is: ", *agePointer)
	//fmt.Println(getAdultYears(age))
}

func getAdultYears(age int) int {
	return age - 18
}
