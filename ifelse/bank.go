package main

import "fmt"

func main() {
	var accountBalance float64 = 1000.0

	fmt.Println("Welcome to Go Bank!!!")
	fmt.Println("What would you like to do today?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scanln(&choice)

	wantToCheckBalaaance := choice == 1

	if wantToCheckBalaaance {
		fmt.Println("Your balance is:", accountBalance)
	} else if choice == 2 {
		fmt.Println("Enter the amount you want to deposit:")
		var depositAmount float64
		fmt.Scanln(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Your new balance is:", accountBalance)
	} else if choice == 3 {
		fmt.Println("Enter the amount you want to withdraw:")
		var withdrawAmount float64
		fmt.Scanln(&withdrawAmount)

	} else {
		fmt.Println("Thank you for using Go Bank!!!")
	}

	//fmt.Println(("Your choice:"), choice)
}
