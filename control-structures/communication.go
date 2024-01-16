package main

import (
	"fmt"
)

func displayIntro() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")

	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdrawl money")
	fmt.Println("4. Exit")
}
