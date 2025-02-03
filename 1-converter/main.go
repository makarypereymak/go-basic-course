package main

import (
	"fmt"
	"strings"
)

const (
	USDT = "USDT"
	RUB  = "RUB"
	EUR  = "EUR"
)

var currencies = []string{USDT, RUB, EUR}

func getCurrenciesStringList() string {
	var result strings.Builder

	for index, value := range currencies {
		result.WriteString(value)
		if index == len(currencies)-1 {
			result.WriteString(" or ")
		} else {
			result.WriteString(", ")
		}
	}

	return result.String()
}

func main() {
	initialCurrency, targetCurrency, initialNumber := getUserInputs()
	resultNumber := convertCurrencies(initialCurrency, targetCurrency, initialNumber)
	resultString := fmt.Sprintf("%.2f", resultNumber)
	fmt.Print(resultString + " " + targetCurrency)
}

func validateInputCurrency(currentCurrency string) bool {
	return currentCurrency == USDT || currentCurrency == RUB || currentCurrency == EUR
}

func validateInputNumber(number float64) bool {
	return number > 0
}

func getAndValidateInput[T any](valueForScan *T, validateFunc func(T) bool, errorString string) {
	for {
		fmt.Scan(valueForScan)
		if validateFunc(*valueForScan) {
			break
		}
		fmt.Println(errorString)
	}
}

func getUserInputs() (string, string, float64) {
	var initialCurrency string
	var targetCurrency string
	var initialNumber float64
	currenciesStringsList := getCurrenciesStringList()
	fmt.Println("Enter initial currency (" + currenciesStringsList + "): ")
	getAndValidateInput(&initialCurrency, validateInputCurrency, "Invalid currency, enter currently value: ")
	fmt.Println("Enter target currency (" + currenciesStringsList + "USDT, EUR or RUB): ")
	for {
		getAndValidateInput(&targetCurrency, validateInputCurrency, "Invalid currency, enter currently value: ")
		if targetCurrency == initialCurrency {
			fmt.Println("You can't convert to the same currency, choose target currency again: ")
			continue
		}
		break
	}
	fmt.Println("Enter initial number: ")
	getAndValidateInput(&initialNumber, validateInputNumber, "Invalid number, enter currently value")

	return initialCurrency, targetCurrency, initialNumber
}

func convertCurrencies(initialCurrency string, targetCurrency string, initialNumber float64) float64 {
	usdtToEur := 0.97515
	usdtToRub := 101.640033
	exchangeRate := usdtToEur
	if initialCurrency == USDT {
		if targetCurrency == RUB {
			exchangeRate = usdtToRub
		}
	}

	if initialCurrency == EUR {
		if targetCurrency == USDT {
			exchangeRate = 1 / usdtToEur
		}

		if targetCurrency == RUB {
			exchangeRate = (1 / usdtToEur) * usdtToRub
		}
	}

	if initialCurrency == RUB {
		if targetCurrency == EUR {
			exchangeRate = 1 / ((1 / usdtToEur) * usdtToRub)
		}

		if targetCurrency == USDT {
			exchangeRate = 1 / usdtToRub
		}
	}
	return exchangeRate * initialNumber
}
