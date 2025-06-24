package main

import (
	"fmt"
	"os"
	"strconv"
)

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balnceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balnceText), 0644)
}
