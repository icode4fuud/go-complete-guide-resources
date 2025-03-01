package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		//place the code that needs to run if an error is returned to make the code more robust
		return 1000, errors.New("Error reading file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Error parsing float")
	}

	return balance, nil
}

// add new function that can do basic I/O
func writeBalanceToFile(balance float64) {
	//write to file
	//file, err := os.Create("balance.txt")
	balanceText := fmt.Sprint(balance)
	os.WriteFile((accountBalanceFile), []byte(balanceText), 0644) //0644 is an encoding format of file permissions for linux
	//os.WriteFile("balance.txt", []byte(fmt.Sprintf("%f", balance)), 0644)
}

// add new function to read from file
func readBalanceFromFile() {
	os.ReadFile(accountBalanceFile)
}

func main() {
	accountBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

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

	//replace w/ switch
	if wantToCheckBalaaance {
		fmt.Println("Your balance is:", accountBalance)
	} else if choice == 2 {
		fmt.Println("Enter the amount you want to deposit:")
		var depositAmount float64
		fmt.Scanln(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Your new balance is:", accountBalance)
		writeBalanceToFile(accountBalance)
	} else if choice == 3 {
		fmt.Println("Enter the amount you want to withdraw:")
		var withdrawAmount float64
		fmt.Scanln(&withdrawAmount)

	} else {
		fmt.Println("Thank you for using Go Bank!!!")
	}

	//fmt.Println(("Your choice:"), choice)
}
