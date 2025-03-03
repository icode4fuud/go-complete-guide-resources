package main

import (
	"fmt"
)

//type transformfn func(int) int

func main() {
	numbers := []int{1, 2, 3}

	//creating an anonymous function
	// anonymous functions are useful when you want to create a function that is only used once
	// and you don't want to give it a name
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}
