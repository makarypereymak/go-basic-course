package main

import (
	"fmt"
)

func main() {
	usdtToEur := 0.97515
	usdtToRub := 101.640033
	eurToRub := (1 / usdtToEur) * usdtToRub
	fmt.Print(eurToRub)
}


func getUserInputs () (string, string, float64) {
	var initialCurrency string
	var targetCurrency string
	var initialNumber float64
	fmt.Println("Enter initial currency: ")
	fmt.Scan(&initialCurrency)	
	fmt.Println("Enter target currency: ")
	fmt.Scan(&targetCurrency)
	fmt.Println("Enter initial number: ")	
	fmt.Scan(&initialNumber)

	return initialCurrency, targetCurrency, initialNumber
}

func convertCurrencies (initialCurrency string,  targetCurrency string, initialNumber float64) string {
	return ""
}
