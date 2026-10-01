package main

import (
	in "fmt"
)

var (
	energy, energy1, energy2, energy3, cost int     = 12, 22, 7, 0, 5
	salary, bonus                           float64 = 1500.50, 500.50
	student, student1                       string  = "Anna", "Belle"
	users                                           = map[string]string{
		"alex":  "Dushanbe",
		"john":  "London",
		"maria": "Berlin",
	}
	num, num1, num2, num3, num4, num5 int = 23, 4, 56, 12, 6, 0
	arr                                   = [5]int{1, 2, 3, 4, 5}
	votes                                 = []string{"Joan William", "Anna Fox", "Joan William", "Sarah Rogerfort", "Anna Fox", "Sarah Rogerfort", "Anna Fox"}
	vote                              string
	cand1, cand2, cand3               int
	currentTemp                       float64 = 37.5

	currentLap int  = 0
	totalLaps  int  = 5
	finished   bool = false
	attempts        = []float64{12.5, 11.8, 12.1, 10.9, 11.2, 10.5}
	best            = attempts[0]
)

func main() {
	task1()
	task2()
	task3()

	spendEnergy(&energy, cost) // task 4
	spendEnergy(&energy1, cost)
	spendEnergy(&energy2, cost)
	spendEnergy(&energy3, cost)
	in.Println("Task 4):", energy, energy1, energy2, energy3)

	double(energy) // task 5
	in.Print("Task 5): ", energy)
	doublePtr(&energy)
	in.Print(" ", energy)

	applyBonus(&salary, bonus) // task 6
	in.Println("\nTask 6): ", float64(salary))

	in.Print("Task 7): ", student, " ", student1)
	swapString(&student, &student1) // task7
	in.Println(" | ", student, student1)

	in.Println("Task 8): ")
	city := findCity(users, "alex") //task 8

	if city != nil {
		in.Println("City:", *city)
	} else {
		in.Println("User not found")
	}

	city = findCity(users, "bob")

	if city != nil {
		in.Println("City:", *city)
	} else {
		in.Println("User not found")
	}

	in.Println("Task 9): ", safeDivide(num, num1))
	in.Println(safeDivide(num2, num3))
	in.Println(safeDivide(num4, num5))

	task10()

	in.Print("Task 11): ", arr)
	doubleArray(&arr)
	in.Print("|| after: ", arr)

	in.Println("\nTask 12): ") //task 12
	for _, vote = range votes {
		tallyVote(vote, &cand1, &cand2, &cand3)
	}
	in.Println("Joan William", cand1)
	in.Println("Anna Fox", cand2)
	in.Println("Sarah Rogerfort", cand3)

	in.Println("Task 13): ") //task 13
	in.Println(adjustTemperature(&currentTemp, 34.8, 21.3, 41.4), currentTemp)
	in.Println(adjustTemperature(&currentTemp, 17.5, 19.8, 25.2), currentTemp)
	in.Println(adjustTemperature(&currentTemp, 23.5, 11.8, 29.2), currentTemp)
	in.Println(adjustTemperature(&currentTemp, 41.2, 15.6, 34.6), currentTemp)

	in.Println("Task 14): ") // task 14

	for i := 0; i < 5; i++ {
		advanceLap(&currentLap, totalLaps, &finished)
		in.Println("Current lap:", currentLap, "Finished:", finished)
	}

	advanceLap(&currentLap, totalLaps, &finished)
	advanceLap(&currentLap, totalLaps, &finished)

	in.Println("After finish:", currentLap, "Finished:", finished)

	in.Print("Task 15): ")
	in.Println("\nAttempts:") //task 15

	for i, attempt := range attempts {
		if trackRecord(&best, attempt) {
			in.Println("Attempt", i+1, "new record:", attempt)
		}
	}

	in.Println("Final record:", best)

}

func task1() {
	age := 16
	in.Printf("\nTask 1): %v, %v, %T", age, &age, age)
}

func task2() {
	temperature := 45
	temp_pointer := &temperature

	*temp_pointer -= 5
	in.Println("\nTask 2): ", temperature)
}

func task3() {
	city := "Dushanbe"
	point_city := &city

	*point_city = "Kyoto"
	in.Println("Task 3): ", city)
}

func spendEnergy(energy *int, cost int) {
	*energy -= cost

	if *energy < cost {
		*energy = 0
	}

}

func double(n int) int {
	return n * 2
}

func doublePtr(n *int) {
	*n *= 2
}

func applyBonus(salary *float64, bonus float64) {
	*salary += bonus
}

func swapString(a, b *string) {
	*a, *b = *b, *a
}

func findCity(users map[string]string, login string) *string {
	city, ok := users[login]

	if !ok {
		return nil
	}

	return &city
}

func safeDivide(a, b int) *int {

	if b == 0 {
		return nil
	} else {
		res := a / b
		return &res
	}
}

func task10() {
	p := new(int)
	*p = 100

	in.Println("Task 10): ", p)
}

func doubleArray(arr *[5]int) {
	for i := range arr {
		arr[i] *= 2
	}
}

func tallyVote(candidate string, countA, countB, countC *int) {
	switch candidate {
	case "Joan William":
		*countA++
	case "Anna Fox":
		*countB++
	case "Sarah Rogerfort":
		*countC++
	}

}

func adjustTemperature(current *float64, delta, min, max float64) bool {
	if delta > max {
		*current = max
		return true
	} else if delta < min {
		*current = min
		return true
	} else {
		*current = delta
		return false
	}
}

func advanceLap(currentLap *int, totalLaps int, finished *bool) {
	if *finished {
		return
	}

	*currentLap++

	if *currentLap >= totalLaps {
		*currentLap = totalLaps
		*finished = true
	}
}

func trackRecord(best *float64, attempt float64) bool {
	if attempt < *best {
		*best = attempt
		return true
	}

	return false
}
