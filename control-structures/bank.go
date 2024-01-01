package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountNumber = 123456
const accountBalanceFile = "balance.txt"

var accountBalance = getBalanceFromFile()

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func main() {
	for getChoice() != 4 {
		getChoice()
	}
}

func displayIntro() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")

	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdrawl money")
	fmt.Println("4. Exit")
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

	writeBalanceToFile(accountBalance)
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

	writeBalanceToFile(accountBalance)
}

func getBalance(accountNumber int) interface{} {
	if accountNumber == 123456 {
		return 1000.0
	}

	return "Invalid account number"
}
