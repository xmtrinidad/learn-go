package main

import "fmt"

const accountNumber = 123456

var accountBalance = getBalance(accountNumber)

func main() {
	displayIntro()
	getChoice()

}

func displayIntro() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")

	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

func getChoice() {
	var choice int

	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		checkBalance()
	} else if choice == 2 {
		depositMoney()
	}
}

func checkBalance() {
	switch accountBalance := accountBalance.(type) {
	case float64:
		fmt.Println("Your balance is: ", accountBalance)
	case string:
		fmt.Println("No account found")
	}
}

func depositMoney() {

	fmt.Println("Your deposit: ")
	var depositAmount float64
	fmt.Scan(&depositAmount)

	accountBalance = accountBalance.(float64) + depositAmount

	fmt.Println("Your new balance is: ", accountBalance)
}

func getBalance(accountNumber int) interface{} {
	if accountNumber == 123456 {
		return 1000.0
	}

	return "Invalid account number"
}
