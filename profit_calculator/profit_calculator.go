package main

import "fmt"

func main() {
	revenue := getUserInput("Input Revenue: ")
	expenses := getUserInput("Input expenses: ")
	taxRate := getUserInput("Input Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println("Ebt result: ", ebt)
	fmt.Println("Profit Result", profit)
	fmt.Printf("Ratio Result: %0.2f", ratio)

}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var inputUser float64
	fmt.Print(infoText)
	fmt.Scan(&inputUser)
	return inputUser
}
