package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5

	// value not set because user will input it
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	// Have to use a pointer to the variable to get the value
	fmt.Println("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Println("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	formattedFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.2f\n", futureValue)

	// fmt.Println("Future Value: ", futureValue)
	fmt.Printf("Future Value: %.2f\n", futureValue)
	fmt.Print(futureRealValue)
	fmt.Print(formattedFV)
}

func outputText(text string) {
	fmt.Println(text)
}
