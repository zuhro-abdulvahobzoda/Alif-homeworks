package main

import (
	in "fmt"
)

var (
	cardBalance int
	withdrawSum int
)

func main() {
	in.Println("----- ATM -----")

	in.Print("*What's your card balance? :")
	in.Scan(&cardBalance)

	in.Print("\n*How much you want to withdraw? :")
	in.Scan(&withdrawSum)

	if cardBalance < 0 {
		in.Println("\nCard Balance cannot be negative!")
	} else {

		if withdrawSum <= 0 {
			in.Print("\nWithdraw sum cannot be negative or zero!")
		} else if withdrawSum > cardBalance {
			in.Println("\nNot enough recources")
		} else {
			cardBalance -= withdrawSum
			in.Printf("\nCurrent balance: %d", cardBalance)
		}

	}

}
