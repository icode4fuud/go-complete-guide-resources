package main

func main() {
	fact := factorial(5)
	println(fact)
}

func factorial(number int) int {
	// result := 1
	// for i := 1; i <= number; i++ {
	// 	result *= i
	// }
	// return result

	//now implement the same using recursion
	if number == 0 {
		return 1
	}

	return number * factorial(number-1)
}

//factorial of 5 => 5 * 4 * 3 * 2 * 1 = 120
