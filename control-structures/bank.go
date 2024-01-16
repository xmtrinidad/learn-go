package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountNumber = 123456
const accountBalanceFile = "balance.txt"

var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

func main() {
	for getChoice() != 4 || err != nil {
		getChoice()
	}
}

func getChoice() int {
	displayIntro()

	var choice int

	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		checkBalance()
	} else if choice == 2 {
		depositMoney()
	} else if choice == 3 {
		withdrawlMoney()
	} else {
		fmt.Println("Thank you for using our service!")
	}

	return choice
}

func withdrawlMoney() {
	var withdrawlAmount float64
	fmt.Print("withdrawl: ")
	fmt.Scan(&withdrawlAmount)

	accountBalance = accountBalance - withdrawlAmount

	fmt.Println("Your new balance is: ", accountBalance)

	fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
}

func checkBalance() {
	fmt.Println("Your balance is: ", accountBalance)
}

func depositMoney() {

	fmt.Println("Your deposit: ")
	var depositAmount float64
	fmt.Scan(&depositAmount)

	accountBalance = accountBalance + depositAmount

	fmt.Println("Your new balance is: ", accountBalance)

	fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
}

func getBalance(accountNumber int) interface{} {
	if accountNumber == 123456 {
		return 1000.0
	}

	return "Invalid account number"
}
