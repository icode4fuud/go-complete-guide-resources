package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 3.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	//outputs to standard output
	//fmt.Println("Future Value:", futureValue)
	//fmt.Printf("Future Value: %.2f\n", futureValue)
	fmt.Printf("Future Value: %.02f\nFuture Value(adjusted for inflation): %.02f", futureValue, futureRealValue)
	//fmt.Printf("Future Value: %v\nFuture Value(adjusted for inflation): %v", futureValue, futureRealValue)
	//fmt.Println("Future Real Value (adjusted for inflation):", futureRealValue)

	//outputing fomatted string
	formattedFB := fmt.Sprintf("Future Value: %.02f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value(adjusted for inflation): %.02f", futureRealValue)
	fmt.Print(formattedFB, formattedFRV)

}
