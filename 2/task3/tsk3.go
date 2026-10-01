package main

import (
	in "fmt"
)

var (
	fuelType   string
	price      float32
	fuelVolume float32
	total      float32
)

func main() {

	in.Print("\nWhat type of fuel do you want to use? (92/ 95/ 98/ diesel): ")
	in.Scan(&fuelType)

	in.Print("\nHow many liters do you want to take?: ")
	in.Scan(&fuelVolume)

	switch fuelType {
	case "92":
		price = 10
		in.Printf("\n%v fuel for 1 liter is: %v", fuelType, price)
	case "95":
		price = 12
		in.Printf("\n%v fuel for 1 liter is: %v", fuelType, price)
	case "98":
		price = 15
		in.Printf("\n%v fuel for 1 liter is: %v", fuelType, price)
	case "diesel":
		price = 11
		in.Printf("\n%v fuel for 1 liter is: %v", fuelType, price)
	default:
		in.Println("Invalid fuel type!")

	}
	in.Printf("\nNumber of liters: %v", fuelVolume)
	total = price * fuelVolume
	in.Printf("\nTotal price: %v\n", total)
}
