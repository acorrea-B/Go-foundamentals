package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Input invest Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	// formattedRFV := fmt.Sprintf("nFuture Value Adjusted: %.2f\n", futureRealValue)

	fmt.Printf(`Future Value: %.2f\n
	Future Value Adjusted: %.2f`, futureValue, futureRealValue)

	// fmt.Print(formattedFV, formattedRFV)
}
