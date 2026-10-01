package main

import (
	in "fmt"
)

var role string

func main() {

	in.Print("Enter your role (admin/ manager/ user/ guest): ")
	in.Scan(&role)

	switch role {
	case "admin":
		in.Println("Full access")
	case "manager":
		in.Println("User management")
	case "user":
		in.Println("Default access")
	case "guest":
		in.Println("Read-only")
	default:
		in.Println("Unknown role")
	}
}
