package main

import (
	"fmt"
)

const accountBalanceFile = "balance.txt"

func main() {
	// condicionalBankOptions(accountBalance)
	for {
		var accountBalance = getBalanceFromFile()

		optionsMessage()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your account balance is: ", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Withdwa amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Good bye!")
			fmt.Println("Thanks for visit your bank.")
			return
		}

	}
}

/*
func condicionalBankOptions(accountBalance float64) {
	for {

		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Desposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your account balance is: ", accountBalance)
			continue
		}

		if choice == 2 {
			var depositAmount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
		} else if choice == 3 {
			var withdrawAmount float64
			fmt.Print("Withdwa amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
		} else {
			fmt.Println("Good bye!")
			break
		}
	}
	fmt.Println("Thanks for visit your bank.")
}
*/
