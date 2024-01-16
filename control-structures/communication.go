package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func displayIntro() {
	fmt.Printf("Welcome %s to Go Bank!", randomdata.SillyName())
	fmt.Println("What do you want to do?")

	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdrawl money")
	fmt.Println("4. Exit")
}
