package main

import "fmt"

//by added this file to the main package it can be referenced in all files in the main package
func presentOptions() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}
