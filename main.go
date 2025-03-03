package main

import (
	"fmt"
	"sync"
	"time"
)

func add(a, b int) int {
	return a + b
}

// printNumbers() will run as go routine
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

// printNumbers() will run as go routine using WaitGroup
func printNumbersWg(wg *sync.WaitGroup) {
	defer wg.Done() // Notify WaitGroup when done
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

// Begin main()
func main() {

	// Function call
	result := add(3, 5) // Function call
	fmt.Println(result) // Output: 8

	//run printNumbers() as go routine
	go printNumbers()
	fmt.Println("Main function continues (printNumbers())...")
	time.Sleep(3 * time.Second) // Wait for the goroutine to finish

	//run printNumbersWg() as go routine using WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)              // Add 1 to the WaitGroup counter
	go printNumbersWg(&wg) // Goroutine
	fmt.Println("Main function continues (printNumbersWg())...")
	wg.Wait() // Wait for the goroutine to finish
}

// End main()
