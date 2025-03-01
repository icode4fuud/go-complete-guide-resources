package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxrate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxrate)

	ebt := revenue - expenses
	profit := ebt - (ebt * taxrate / 100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

}
