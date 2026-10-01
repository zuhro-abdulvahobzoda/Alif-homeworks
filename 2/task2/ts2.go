package main

import in "fmt"

var (
	salary      float32
	tax         float32
	sumAfterTax float32
)

func main() {
	in.Print("Your salary: ")
	in.Scan(&salary)

	if salary < 3000 {
		tax = salary * 0.05
		sumAfterTax = salary - tax
		in.Printf("Your Tax sum: %v", tax)
		in.Printf("\nYour salary after tax: %f\n", sumAfterTax)
	} else if salary >= 3000 && salary <= 10000 {
		tax = salary * 0.10
		sumAfterTax = salary - tax
		in.Printf("Your Tax sum: %f", tax)
		in.Printf("\nYour salary after tax: %f\n", sumAfterTax)
	} else if salary > 10000 {
		tax = salary * 0.15
		sumAfterTax = salary - tax
		in.Printf("Your Tax sum: %f", tax)
		in.Printf("\nYour salary after tax: %f\n", sumAfterTax)
	} else if salary <= 0 {
		in.Println("Salary cannot be zero or negative")
	}

}
