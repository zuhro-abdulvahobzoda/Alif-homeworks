package main

import (
	in "fmt"
	p "goget/printcheck"
	u "goget/user"
)

func main() {
	names := []string{"Anna_Fox", "The_OnlyCh0sen", "Silverside2", "Angelluxury", "Heartdust", "Olivia6767"}

	for i, v := range names {
		in.Printf("%d) %s: %s\n", i+1, u.NewUser(v).Name, u.NewUser(v).ID)
	}

	p.PrintCheck("\n User authorization test", true)
	p.PrintCheck("Connection to the base servers test", true)
	p.PrintCheck("API timeout test", false)
	p.PrintCheck("Form validation test", true)
	p.PrintCheck("Errors test", false)
}
