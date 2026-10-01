package main

import (
	inout "fmt"
)

var (
	currentLevel          int // for the first task
	currentMonth          int8
	distance              int16
	radiusOfSun           int32
	areaOfMilkyWay        int64
	totalOfSheep          uint
	hoursPerDay           uint8
	portNumber            uint16
	userID                uint32
	fileSize              uint64
	averageBuildingHeight float32
	price                 float64
	isAlive               bool
	letter                rune
	characterCode         byte
	username              string
	complexNumber1        complex64
	complexNumber2        complex128

	number1 = 126 // two integers for the second task
	number2 = 56

	name      string // for the fourth task
	age       uint
	height    float32
	weight    float32
	isWorking bool

	num1 int // for the fifth task
	num2 int

	workerName         string = "Alisher" // for the sixth task
	workerSurname      string = "Alimov"
	workerAge          uint8  = 24
	seniority          string = "Junior"
	salary             int    = 200
	favoriteCodingLang string = "Go"
	commercialDevExp   int    = 1
	doesknowDocker     bool   = true
	nameInitial        rune   = 'A'
	favAscSymbol       byte   = 'K'
	totalProjects      uint   = 13

	firstVar  = "Hello" // for the seventh task
	secondVar = 16
	thirdVar  = 'D'

	userName      string // for the last task
	userSurname   string
	userAge       uint8
	userBirthYear uint
	userCity      string
	userJob       string
)

const (
	pi          = 3.14 // for the third task
	compamyName = "Alif Company"
	currentYear = 2026
	daysInWeek  = 7
)

func main() {
	// ================================== First task =================================
	inout.Println("=== FIRST TASK ===")
	inout.Print("Default outputs:", currentLevel, currentMonth, distance, radiusOfSun, areaOfMilkyWay, totalOfSheep, hoursPerDay, portNumber, userID, fileSize, averageBuildingHeight, price, isAlive, letter, characterCode, username, complexNumber1, complexNumber2)
	// default outputs of variables without assigned values

	currentLevel = 56
	currentMonth = 8
	distance = 3456
	radiusOfSun = 695700
	areaOfMilkyWay = 7850000000
	totalOfSheep = 5
	hoursPerDay = 24
	portNumber = 1345
	userID = 6534667
	fileSize = 15
	averageBuildingHeight = 3.6
	price = 15.56
	isAlive = true
	letter = 'Z'
	characterCode = 'A'
	username = "black_saphirre"
	complexNumber1 = 3 + 4i
	complexNumber2 = -2 + 7i

	inout.Print("\nFinal outputs:\n")
	inout.Printf("\nCurrent Level: %v\n", currentLevel) // Final output to the first task
	inout.Printf("Current Month: %v\n", currentMonth)
	inout.Printf("Distance: %v km\n", distance)
	inout.Printf("Radius of the Sun: %v km\n", radiusOfSun)
	inout.Printf("Area of Milky Way: %v light years\n", areaOfMilkyWay)
	inout.Printf("Total of sheep: %v\n", totalOfSheep)
	inout.Printf("Hours per day: %v hours \n", hoursPerDay)
	inout.Printf("Port Number: %v\n", portNumber)
	inout.Printf("User ID: %v\n", userID)
	inout.Printf("Size of the current file: %v\n", fileSize)
	inout.Printf("Average building height: %v meters\n", averageBuildingHeight)
	inout.Printf("Random price: %v\n", price)
	inout.Printf("Is that cat alive?: %v\n", isAlive)
	inout.Printf("Random letter: %v\n", letter)
	inout.Printf("Random character code: %v\n", characterCode)
	inout.Printf("Think a new username: %v\n", username)
	inout.Printf("First complex number : %v\n", complexNumber1)
	inout.Printf("Second complex number: %v\n", complexNumber2)

	// ================================= SECOND TASK ======================================
	inout.Println("\n=== SECOND TASK ===")
	inout.Println("= Arithmetic operations with two integers =")

	inout.Println(number1 + number2)
	inout.Println(number1 - number2)
	inout.Println(number1 * number2)
	inout.Println(number1 / number2)
	inout.Println(number1 % number2)

	// ======================================= THIRD TASK =================================
	inout.Println("\n=== THIRD TASK ===")

	inout.Println(pi, compamyName, currentYear, daysInWeek)

	// =============================== FOURTH TASK =========================================
	inout.Println("\n=== FOURTH TASK ===")

	inout.Println("What's your name?")
	inout.Scan(&name)

	inout.Println("How old are you?")
	inout.Scan(&age)

	inout.Println("What's your height?")
	inout.Scan(&height)

	inout.Println("Your weight?")
	inout.Scan(&weight)

	inout.Println("Do you work? (true/false)")
	inout.Scan(&isWorking)

	inout.Printf("\n===== User's Survey =====")
	inout.Printf("\n* Name: %v", name)
	inout.Printf("\n* Age: %v", age)
	inout.Printf("\n* Height: %v", height)
	inout.Printf("\n* Width: %v", weight)
	inout.Printf("\n* Has a job: %v\n", isWorking)
	inout.Println("====================")

	// =========================== FIFTH TASK =========================================
	inout.Println("\n=== FIFTH TASK ===")
	inout.Println("\n== Mini Calculator ==")

	inout.Printf("\nCome up with the first number: ")
	inout.Scan(&num1)

	inout.Printf("Come up with the second number: ")
	inout.Scan(&num2)

	inout.Println("\nResults:", num1+num2, num1-num2, num1*num2, num1/num2, num1%num2)

	// ============================== SIXTH TASK ====================================
	inout.Println("\n=== SIXTH TASK ===")

	inout.Printf("\n== DEVPASSPORT ==\n")
	inout.Printf("\n* Name: %v", workerName)
	inout.Printf("\n* Surname : %v", workerSurname)
	inout.Printf("\n* Age : %v", workerAge)
	inout.Print("\n*==========*")
	inout.Printf("\n* Seniority : %v", seniority)
	inout.Printf("\n* Salary : %v", salary)
	inout.Printf("\n* Favorite Coding language : %v", favoriteCodingLang)
	inout.Printf("\n* Commercial Developing Experience : %v year(s)", commercialDevExp)
	inout.Printf("\n* Knows Docker : %v", doesknowDocker)
	inout.Print("\n*==========*")
	inout.Printf("\n* Name initial : %c", nameInitial)
	inout.Printf("\n* Favorite ASCII symbol : %c", favAscSymbol)
	inout.Printf("\n* Total projects : %v\n", totalProjects)

	// =============================== SEVENTH TASK =========================================
	inout.Print("\n ======= SEVENTH TASK =======\n")
	inout.Println("-Before changing values:", firstVar, secondVar, thirdVar)

	firstVar = "Good Day"
	secondVar = 88
	thirdVar = 'S'
	inout.Println("--First change: ", firstVar, secondVar, thirdVar)

	firstVar = "Have a great day"
	secondVar = 598
	thirdVar = 'P'
	inout.Println("--Second change: ", firstVar, secondVar, thirdVar)

	firstVar = "Bye"
	secondVar = 554
	thirdVar = 'Y'
	inout.Println("--Last change", firstVar, secondVar, thirdVar)

	// =========================== EITH TASK ==============================================
	inout.Println("\n ======= LAST TASK =======")
	inout.Print("Name?: ")
	inout.Scan(&userName)

	inout.Print("\nSurname?: ")
	inout.Scan(&userSurname)

	inout.Print("\nAge?: ")
	inout.Scan(&userAge)

	inout.Print("\nBirth Year?: ")
	inout.Scan(&userBirthYear)

	inout.Print("\nWhere do you live?: ")
	inout.Scan(&userCity)

	inout.Print("\nYour profession?: ") // I don't know what else to ask ;-;
	inout.Scan(&userJob)

	inout.Println("\n--*=== My Visit Card ===*--")
	inout.Printf("* Name: %v", userName)
	inout.Printf("\n* Surname : %v", userSurname)
	inout.Printf("\n* Age : %v", userAge)
	inout.Printf("\n* Birth year : %v", userBirthYear)
	inout.Printf("\n* City : %v", userCity)
	inout.Printf("\n* Profession : %v", userJob)
	inout.Print("\n--*======================*--")

}
