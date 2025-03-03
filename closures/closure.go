package main

import (
	"fmt"
)

//type transformfn func(int) int

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	triple := createTransformer(3)

	fmt.Println(double(2))
	fmt.Println(triple(2))

	//creating an anonymous function
	// anonymous functions are useful when you want to create a function that is only used once
	// and you don't want to give it a name
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// example of a closure
// a closure is a function that captures the variables from the environment in which it was defined
// in this case, the variable "factor" is captured by the anonymous function
func createTransformer(factor int) func(int) int {
	// return func(number int) int {
	// 	factor := 2
	// 	return number * factor
	// }
	return func(number int) int {
		return number * factor
	}
}
