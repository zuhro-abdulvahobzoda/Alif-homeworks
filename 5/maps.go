package main

import (
	in "fmt"
	"sort"
	"strings"
)

var prices = make(map[string]float64)

func main() {
	map1()
	map2()
	map3()
	map4()
	map5()
	map6()
	map7()
	map8()
	map9()
	map10()
	map11()
	map12()
	map13()

}

func map1() {
	students := map[string]int{
		"Anna":    89,
		"Ben":     45,
		"Mike":    56,
		"William": 41,
	}

	in.Print("Task1): ", students, " | ", students["Anna"])
}

func map2() {

	prices["bread"] = 7.5
	prices["meat"] = 25.77
	prices["gum"] = 1.5
	prices["egg"] = 5.0

	in.Print("\nTask2): ", prices, " | ", len(prices))
}

func map3() {
	if _, ok := prices["cheese"]; ok {
		prices["cheese"] = 8.65
		in.Print("\nTask3): ", prices)
	} else {
		in.Print("\nTask3): Cheese does not exist")
	}
}

func map4() {
	ages := map[string]uint{
		"Tom":    23,
		"Ben":    22,
		"Ginger": 8,
		"Angela": 22,
	}
	in.Print("\nTask4): ", ages)

	ages["Tom"] = 24

	in.Print("Happy birtday!", ages["Tom"])
}

func map5() {
	products := map[string]int{
		"Apple":  5,
		"Banana": 0,
		"Orange": 3,
	}

	delete(products, "Banana")

	_, ok := products["Banana"]

	in.Print("\nTask5): ", products)
	in.Print(" | Banana exists: ", ok)
}

func map6() {
	likes := make(map[string]int)

	likes["Video 1"]++
	likes["Video 1"]++
	likes["Video 1"]++

	likes["Video 2"]++
	likes["Video 2"]++

	in.Print("\nTask6): ", likes)
}

func map7() {
	students := map[string]int{
		"Anna":    89,
		"Ben":     45,
		"Mike":    56,
		"William": 41,
	}

	in.Print("\nTask7): ")

	for name, grade := range students {
		in.Print(name, ": ", grade, " ")
	}

	in.Print(" Only names:")

	for name := range students {
		in.Print(name, " ")
	}
}

func map8() {
	students := map[string]int{
		"Anna":    89,
		"Ben":     45,
		"Mike":    56,
		"William": 41,
	}

	var names []string

	for name := range students {
		names = append(names, name)
	}

	sort.Strings(names)

	in.Print("\nTask8): ")

	for _, name := range names {
		in.Print(name, ": ", students[name])
	}
}

func map9() {
	visitors := []string{
		"Anna",
		"Ben",
		"Anna",
		"Mike",
		"Ben",
		"Anna",
	}

	unique := make(map[string]bool)

	for _, visitor := range visitors {
		unique[visitor] = true
	}

	in.Print("\nTask9): ", unique)
	in.Print(" \"Unique visitors: \"", len(unique))
}

func map10() {
	groups := map[string][]string{
		"Group A": {"Anna", "Ben"},
		"Group B": {"Mike", "William"},
	}

	groups["Group A"] = append(groups["Group A"], "John")

	in.Print("\nTask10): ", groups)
}

func map11() {
	salaries := map[string]int{
		"Anna":    3000,
		"Ben":     4500,
		"Mike":    5000,
		"William": 4000,
	}

	maxName := ""
	maxSalary := 0

	for name, salary := range salaries {
		if salary > maxSalary {
			maxSalary = salary
			maxName = name
		}
	}

	in.Print("\nTask11): ", maxName, " | ", maxSalary)
}

func map12() {
	text := "go is fun and go is simple"

	words := strings.Fields(text)

	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}

	var sortedWords []string

	for word := range counts {
		sortedWords = append(sortedWords, word)
	}

	sort.Strings(sortedWords)

	in.Print("\nTask12): ")
	for _, word := range sortedWords {
		in.Print(word, ": ", counts[word], " ")
	}

	mostWord := ""
	maxCount := 0

	for word, count := range counts {
		if count > maxCount {
			maxCount = count
			mostWord = word
		}
	}

	in.Print("| Most frequent: ", mostWord, " | ", maxCount)
}

func map13() {
	cities := []string{
		"Dushanbe",
		"Moscow",
		"Dushanbe",
		"London",
		"Moscow",
		"Dushanbe",
	}

	counts := make(map[string]int)

	for _, city := range cities {
		counts[city]++
	}

	in.Print("\nTask13): ", counts)
}
