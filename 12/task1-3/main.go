package main

import (
	in "fmt"
	g "task1/geometry"
	s "task1/strutil"
	v "task1/validator"
)

type (
	Rectangle struct {
		Width  float64
		Height float64
	}
)

func main() {
	/////////////////////////////////////////////////////////////////////////////////////

	in.Println("===Task 1===")

	rctngls := []Rectangle{
		{Width: 12.3, Height: 12.0},
		{Width: 56.4, Height: 20.4},
		{Width: 5.4, Height: 2.1},
		{Width: 1.9, Height: 16.3},
	}

	for i, v := range rctngls {
		in.Printf("%d) Area: %.2f", i+1, g.RectangleArea(v.Width, v.Height))
		in.Printf(" || Perimeter: %.2f\n", g.RectanglePerimeter(v.Width, v.Height))
	}

	/////////////////////////////////////////////////////////////////////////////////////////

	in.Println("===Task 2===")

	strs := []string{
		"racecar", "kayak", "man", "city", "apple", "civic", "level",
	}

	for i, v := range strs {
		in.Printf("%d) %s", i+1, s.Reverse(v))
		in.Printf(" || Palindrome: %v\n", s.IsPalindrome(v))
	}

	//////////////////////////////////////////////////////////////////////////////////////////

	in.Println("===Task 3===")

	emails := map[string]bool{
		"user@example.com": true,
		"a@b.c":            true,
		"plainaddress":     false,
		"@example.com":     false,
		"user@":            false,
		"user@domain":      false,
		"user@.com":        false,
		"user@domain.":     false,
		"user@@domain.com": false,
	}

	for email, expected := range emails {
		result := v.IsValidEmail(email)
		status := "FAIL"
		if result == expected {
			status = "OK"
		}
		in.Printf("[%s] IsValidEmail(\"%s\") = %v\n", status, email, result)
	}
}
