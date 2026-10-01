package main

import (
	in "fmt"
	str "strings"
	"unicode"
)

var (
	price, percent        float64
	age                   = 18
	a, b, c               = 12, 56, 4
	weight, height, width = 49.8, 161.1, 200.6
	seconds               = 34621974
	firstName, lastName   = "Alex", "Handerson"
	fullName1             = "Mike Tyson"
	password              = "qwertyoftime"
	x1, y1, x2, y2        = 12.3, 18.4, 22.4, 1.2
	nums                  = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	strs                  = []string{"go", "python", "java", "c++", "php", "assembly"}
	prices                = map[string]float64{
		"Apple":  3.5,
		"Milk":   2.5,
		"Banana": 2.9,
		"Bread":  1.5,
	}
)

func main() {
	////////////////////////////////////////////////////////////////////
	printWelcome()
	printDivider()
	showMenu()
	announceGameOver()
	chapterDivider()
	////////////////////////////////////////////////////////////////////
	in.Printf("Total: %.2f", discountPrice(price, percent))
	in.Println("\n", isAdult(age))
	in.Println("Max of three: ", maxOfThree(a, b, c))
	in.Println("BMI: ", bmi(weight, height))
	in.Println("Seconds to minutes: ", secondsToMinutes(seconds))
	in.Println("First and Last name: ", fullName(firstName, lastName))
	chapterDivider()
	///////////////////////////////////////////////////////////////////

	quontent, remainder := divmod(a, b)
	in.Println("Divide two integers: ", quontent, ",", remainder)

	areaRec, perimeterRec := rectangleParams(width, height)
	in.Println("Rectangle parameters: ", areaRec, perimeterRec)

	firstName1, lastName1 := parseFullName(fullName1)
	in.Println("Parsed Full Name", firstName1, lastName1)

	isEnough, message := checkPassword(password)
	in.Println("The password is: ", isEnough, ",", message)

	in.Println(midPoints(x1, y1, x2, y2))
	chapterDivider()

	//////////////////////////////////////////////////////////////////
	printSlice(nums)
	in.Println(sliceLenghts(strs))
	in.Println(makeFruits())
	in.Println(addToSlice(nums, age))
	printPrices(prices)
	in.Print(makeScores())
	chapterDivider()
	///////////////////////////////////////////////////////////////////
	in.Println(applyDiscount([]float64{100, 200, 300}, 10))

	in.Println(wordFrequency([]string{"apple", "banana", "apple", "apple"}))

	in.Println(filterEven([]int{1, 2, 3, 4, 5, 6}))

	name, price := mostExpensive(map[string]float64{
		"Phone":  500,
		"Laptop": 1000,
		"Mouse":  50,
	})
	in.Println(name, price)

	in.Println(averageGrade(map[string]int{
		"Alex":  5,
		"Bob":   4,
		"Sarah": 5,
	}))

	letters, digits := countLettersAndDigits("Hello123")
	in.Println(letters, digits)

}

func printWelcome() {
	in.Println("Welcome to our app!")
}

func printDivider() {
	in.Print("=========================")
}

func showMenu() {
	coffeMenu := []string{"Latte", "Espresso", "Americano", "Cappuchino"}

	for index, value := range coffeMenu {
		in.Printf("\n %d. %s", index+1, value)
	}
}

func announceGameOver() {
	in.Println("\nGAME OVER")
	in.Println("*************************")
}

func chapterDivider() { // function that divides one group of tasks to another
	for i := 0; i <= 3; i++ {
		in.Println("==")
	}

}

func discountPrice(price, percent float64) float64 {

	in.Println("What's the price of the product?")
	in.Scan(&price)

	in.Println("What's the discount percent?")
	in.Scan(&percent)

	return price - (price * percent)
}

func isAdult(age int) bool {

	if age >= 18 {
		return true
	} else {
		return false
	}
}

func maxOfThree(a, b, c int) int {
	slc1 := []int{a, b, c}
	max := a

	for _, value := range slc1 {
		if value > max {
			max = value
		}
	}

	return max
}

func bmi(weightKg, heightM float64) float64 {
	return weightKg / (heightM * heightM)
}

func secondsToMinutes(seconds int) int {
	return seconds / 60
}

func fullName(firstName, lastName string) string {
	return firstName + " " + lastName
}

func divmod(a, b int) (int, int) {
	return (a / b), (a % b)
}

func rectangleParams(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return area, perimeter
}

func parseFullName(full string) (string, string) {
	parts := str.Fields(full)
	return parts[0], parts[1]
}

func checkPassword(pass string) (bool, string) {
	if len(pass) < 8 {
		return false, "the pass is too short"
	} else {
		return true, "OK"
	}
}

func midPoints(x1, y1, x2, y2 float64) (float64, float64) {
	midpointX := (x1 + x2) / 2
	midpointY := (y1 + y2) / 2
	return midpointX, midpointY
}

func printSlice(nums []int) {

	in.Println(nums)

}

func sliceLenghts(strs []string) int {
	return len(strs)
}

func makeFruits() []string {
	fruits := []string{"apple", "banana", "pear", "peach"}
	return fruits
}

func addToSlice(nums []int, value int) []int {
	nums = append(nums, value)
	return nums
}

func printPrices(prices map[string]float64) {
	for product, price := range prices {
		in.Print(product, ":", price, ", ")
	}
}

func makeScores() map[string]float64 {
	scores := map[string]float64{
		"Anna": 92.8,
		"Ben":  67.9,
		"Cole": 88.5,
	}

	return scores
}

func applyDiscount(prices []float64, percent float64) []float64 {
	result := []float64{}

	for _, price := range prices {
		newPrice := price - price*percent/100
		result = append(result, newPrice)
	}

	return result
}

func wordFrequency(words []string) map[string]int {
	result := map[string]int{}

	for _, word := range words {
		result[word]++
	}

	return result
}

func filterEven(nums []int) []int {
	result := []int{}

	for _, num := range nums {
		if num%2 == 0 {
			result = append(result, num)
		}
	}

	return result
}

func mostExpensive(prices map[string]float64) (string, float64) {
	var expensiveName string
	var expensivePrice float64

	for name, price := range prices {
		if price > expensivePrice {
			expensiveName = name
			expensivePrice = price
		}
	}

	return expensiveName, expensivePrice
}

func averageGrade(grades map[string]int) float64 {
	sum := 0

	for _, grade := range grades {
		sum += grade
	}

	return float64(sum) / float64(len(grades))
}

func countLettersAndDigits(s string) (int, int) {
	letters := 0
	digits := 0

	for _, char := range s {
		if unicode.IsLetter(char) {
			letters++
		}

		if unicode.IsDigit(char) {
			digits++
		}
	}

	return letters, digits
}
