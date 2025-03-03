package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30, 40, 50}
	taxRates := []float64{0, 0.7, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		// result[taxRate] = transformNumbers(&prices, func(price float64) float64 {
		// 	return price * (1 + taxRate)
		// })
		taxIncludedPrices := make([]float64, len(prices))
		for priceIndex, price := range prices {
			taxIncludedPrices[priceIndex] = price * (1 + taxRate)
		}
		result[taxRate] = taxIncludedPrices
	}

	fmt.Println(result)

}
