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
