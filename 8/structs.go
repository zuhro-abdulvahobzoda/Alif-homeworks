package main

import (
	in "fmt"
)

var (
	sum                  float64
	discountedProducts   []ProductFinal
	totalPrice           float64
	totalCompletedAmount float64
	letterGrade          string
)

type (
	Book struct {
		Title  string
		Author string
		Year   int
		Price  float64
	}

	Car struct {
		Brand, Model string
		Year         int
	}

	Student struct {
		Name   string
		Grade  int
		Passed bool
	}

	Point struct {
		X float64
		Y float64
	}

	Laptop struct {
		Brand string
		RAM   int
		SSD   int
		Price float64
	}

	Movie struct {
		Title    string
		Duration int
		Rating   float64
	}

	City struct {
		Name       string
		Population int
		isCapital  bool
	}

	Account struct {
		Owner   string
		Balance float64
	}

	Temperature struct {
		Celsius float64
	}

	Product struct {
		Name     string
		Price    float64
		Quantity int
	}

	Person struct {
		Name string
		Age  int
	}

	Rectangle struct {
		Width, Height float64
	}

	Employee struct {
		Name   string
		Salary float64
	}

	Counter struct {
		Value int
	}

	Address struct {
		City   string
		Street string
	}

	PersonforAddress struct {
		Name    string
		Age     int
		Address Address
	}

	Engine struct {
		HorsePower int
		Fuel       string
	}

	CarEng struct {
		Brand  string
		Engine Engine
	}

	Dimensions struct {
		Width, Height, Depth float64
	}

	Box struct {
		Name       string
		Dimensions Dimensions
	}

	Passport struct {
		Number   int
		issuedBy string
	}

	Citizen struct {
		Name     string
		Passport Passport
	}

	Address1 struct {
		City   string
		Street string
		Zip    int
	}

	Company struct {
		Name    string
		Address Address1
	}

	Employee1 struct {
		Name    string
		Company Company
	}

	Engine1 struct {
		Power int
		Type  string
	}

	CarTask21 struct {
		Brand  string
		Engine Engine1
	}

	Wheel struct {
		Size  int
		Brand string
	}

	CarTask22 struct {
		Model      string
		FrontWheel Wheel
		RearWheel  Wheel
	}

	Book1 struct {
		Title string
		Pages int
		Price float64
	}

	Product1 struct {
		Name  string
		Price float64
	}

	Order struct {
		ID        int
		Amount    float64
		Completed bool
	}

	Student1 struct {
		Name  string
		Grade int
	}

	Product2 struct {
		Name     string
		Price    float64
		Category string
	}

	ProductFinal struct {
		Name       string
		Category   string
		FinalPrice float64
	}

	Day struct {
		Name  string
		TempC float64
	}

	Match struct {
		Home      string
		Away      string
		HomeScore int
		AwayScore int
	}

	Transaction struct {
		Type   string
		Amount float64
	}

	Passenger struct {
		Name string
		Age  int
	}
)

func main() {

	book1 := Book{"Harry Potter", "J.K.Rowling", 1997, 10.5} // task 1
	in.Printf("Task 1): Book #1: %+v", book1)

	emptyCar := Car{} // task 2
	in.Printf("\nTask 2): %+v", emptyCar)

	student1 := Student{} //task 3
	student1.Name = "Anna"
	in.Printf("\nTask 3): %+v", student1)

	points := Point{3.5, 6.7} // task 4
	in.Printf("\nTask 4): %+v |||| %v", points, points)

	myLaptop := Laptop{"Macbook Neo", 8, 256, 599.99} //task 5
	in.Printf("\nTask 5): %+v", myLaptop)

	sonicMovie := Movie{"Sonic the hedgehog 3", 110, 6.9}
	marioMovie := Movie{"The Super Mario Galaxy Movie", 96, 6.3}
	projectHM := Movie{"Project Hail Mary", 156, 8.2}

	//////////////////////////////////////////////////////////////////////////

	in.Println("\nTask 6): ") //task 6
	in.Println("\n=== Movie #1 ===")
	in.Println("Name:", sonicMovie.Title)
	in.Println("Duration:", sonicMovie.Duration, "min")
	in.Println("Rating: ", sonicMovie.Rating)

	in.Println("\n=== Movie #2 ===")
	in.Println("Name:", marioMovie.Title)
	in.Println("Duration:", marioMovie.Duration, "min")
	in.Println("Rating: ", marioMovie.Rating)

	in.Println("\n=== Movie #3 ===")
	in.Println("Name:", projectHM.Title)
	in.Println("Duration:", projectHM.Duration, "min")
	in.Println("Rating: ", projectHM.Rating)

	//////////////////////////////////////////////////////////////////////////////

	city1 := City{Name: "Tokyo", Population: 123, isCapital: true} // task 7
	city2 := City{Population: 1, Name: "Munich", isCapital: false}
	in.Printf("\nTask 7): %+v ||||| %+v", city1, city2)

	////////////////////////////////////////////////////////////////////////////

	annaAccount := Account{"Anna", 1000.0} //task 8
	in.Printf("\nTask 8): %+v", annaAccount)
	annaAccount.Balance = 1250.0
	in.Printf(" || after: %v", annaAccount.Balance)

	////////////////////////////////////////////////////////////////////////////

	temp := Temperature{39.5} // task 9
	toFarenheit := temp.Celsius*9/5 + 32
	in.Printf("\nTask 9): %v, %v", temp.Celsius, toFarenheit)

	////////////////////////////////////////////////////////////////////////////

	prod1 := Product{"Apple", 5.5, 5} // task 10
	prod2 := Product{"Kiwi", 10.3, 9}
	prod3 := Product{"Cereal", 6.99, 2}
	prod4 := Product{"Milk", 10.99, 3}

	calculatePrice(prod1, sum)
	calculatePrice(prod2, sum)
	calculatePrice(prod3, sum)
	calculatePrice(prod4, sum)

	in.Printf("\nTask 10): %+v,%+v,%+v,%+v \nTotal: %v", prod1, prod2, prod3, prod4, sum)

	////////////////////////////////////////////////////////////////////////////

	person1 := Person{"Mike", 19} // task 11
	in.Print("\nTask 11): ", isAdult(person1))

	////////////////////////////////////////////////////////////////////////////

	rctg1 := Rectangle{6.7, 8.9} // task 12
	in.Printf("\nTask 12): Perimeter: %.1f ||| Area: %v", 2*(rctg1.Height+rctg1.Width), rctg1.Height*rctg1.Width)

	////////////////////////////////////////////////////////////////////////////

	employee := Employee{"Jane", 1500.5} // task 13
	in.Printf("\nTask 13): %+v", employee)
	employee.Salary += (employee.Salary * 0.25)
	in.Printf(" || after: %v", employee.Salary)

	////////////////////////////////////////////////////////////////////////////

	in.Print("\nTask 14):\n") //task 14
	c := Counter{Value: 0}
	c.IncrementThreeTimes()

	////////////////////////////////////////////////////////////////////////////

	p := PersonforAddress{ //task 15
		Name: "Rebecca",
		Age:  22,
		Address: Address{
			City:   "Washington",
			Street: "Sesame Street 14",
		},
	}
	in.Println("Task 15): ", p.Address.City)

	////////////////////////////////////////////////////////////////////////////

	p1 := PersonforAddress{ //task 16
		Name: "Reina",
		Age:  19,
		Address: Address{
			City:   "Washington",
			Street: "Sesame Street 9",
		},
	}
	in.Printf("Task 16): %+v", p1)
	p1.Address.Street = "Rover Street 1"
	in.Printf("\nafter: %+v\n", p1)

	////////////////////////////////////////////////////////////////////////////

	myCar := CarEng{ //task 17
		Brand: "Audi",
		Engine: Engine{
			HorsePower: 300,
			Fuel:       "Gasoline",
		},
	}
	in.Println("\nTask 17): The power of engine:", myCar.Engine.HorsePower)

	////////////////////////////////////////////////////////////////////////////

	boxDimension := Box{ //task 18
		Name:       "Idk",
		Dimensions: Dimensions{34.2, 44.3, 33.6},
	}
	in.Printf("Task 18): %+v ||| %.2f", boxDimension, boxDimension.Dimensions.Depth*boxDimension.Dimensions.Width*boxDimension.Dimensions.Height)

	////////////////////////////////////////////////////////////////////////////

	citizen1 := Citizen{ //task 19
		Name:     "Alan",
		Passport: Passport{},
	}
	in.Printf("\nTask 19): %+v", citizen1)

	////////////////////////////////////////////////////////////////////////////

	anEmp := Employee1{ // task 20
		Name: "Jin",
		Company: Company{
			Name: "Nintendo",
			Address: Address1{
				City:   "Kyoto",
				Street: "11-1 Kamitoba Hokodatecho",
				Zip:    489351289,
			},
		},
	}
	in.Print("\nTask 20): This worker is in ", anEmp.Company.Address.City)

	//////////////////////////////////////////////////////////////

	in.Print("\nTask 21): ") // task 21
	car1 := CarTask21{
		Brand:  "BMW",
		Engine: Engine1{Power: 300, Type: "V6"},
	}
	car2 := CarTask21{
		Brand:  "Audi",
		Engine: Engine1{Power: 250, Type: "I4"},
	}

	if car1.Engine.Power > car2.Engine.Power {
		in.Printf("The %s has a more powerful engine (%d HP vs %d HP)", car1.Brand, car1.Engine.Power, car2.Engine.Power)
	} else if car2.Engine.Power > car1.Engine.Power {
		in.Printf("The %s has a more powerful engine (%d HP vs %d HP)", car2.Brand, car2.Engine.Power, car1.Engine.Power)
	} else {
		in.Println("Both cars have the same engine power.")
	}
	/////////////////////////////////////////////////////////////

	in.Print("\nTask 22): ") // task 22
	notMyCar := CarTask22{
		Model:      "Porsche 911",
		FrontWheel: Wheel{Size: 19, Brand: "Michelin"},
		RearWheel:  Wheel{Size: 20, Brand: "Michelin"},
	}

	if notMyCar.FrontWheel.Size > notMyCar.RearWheel.Size {
		in.Printf("Front wheel is larger (%d\" vs %d\")", notMyCar.FrontWheel.Size, notMyCar.RearWheel.Size)
	} else if notMyCar.RearWheel.Size > notMyCar.FrontWheel.Size {
		in.Printf("Rear wheel is larger (%d\" vs %d\")", notMyCar.RearWheel.Size, notMyCar.FrontWheel.Size)
	} else {
		in.Println("Both wheels are the same size.")
	}
	//////////////////////////////////////////////////////////////////

	in.Print("\nTask 23): ") // task 23
	books := []Book1{
		{Title: "War and Peace", Pages: 1225, Price: 25.50},
		{Title: "Crime and Punishment", Pages: 672, Price: 15.00},
		{Title: "The Master and Margarita", Pages: 448, Price: 18.00},
		{Title: "1984", Pages: 320, Price: 12.00},
	}

	maxPagesBook := books[0]
	for _, book := range books[1:] {
		if book.Pages > maxPagesBook.Pages {
			maxPagesBook = book
		}
	}
	in.Printf("Book with the most pages: \"%s\" (%d pages)", maxPagesBook.Title, maxPagesBook.Pages)

	/////////////////////////////////////////////////////////////////////

	in.Print("\nTask 24): ") //task 24
	products := []Product1{
		{Name: "Bread", Price: 2.50},
		{Name: "Milk", Price: 3.80},
		{Name: "Cheese", Price: 7.20},
		{Name: "Apples", Price: 4.10},
		{Name: "Coffee", Price: 9.90},
	}

	for _, product := range products {
		totalPrice += product.Price
	}
	avgPrice := totalPrice / float64(len(products))

	in.Printf("Total price of all products: %.2f, ", totalPrice)
	in.Printf("| Average product price: %.2f", avgPrice)

	///////////////////////////////////////////////////////////////////////

	in.Print("\nTask 25): ")
	orders := []Order{
		{ID: 1, Amount: 150.00, Completed: true},
		{ID: 2, Amount: 300.50, Completed: false},
		{ID: 3, Amount: 45.00, Completed: true},
		{ID: 4, Amount: 120.00, Completed: false},
		{ID: 5, Amount: 89.00, Completed: true},
		{ID: 6, Amount: 210.00, Completed: false},
	}

	completedCount := 0

	for _, order := range orders {
		if order.Completed {
			completedCount++
			totalCompletedAmount += order.Amount
		}
	}

	in.Printf("Number of completed orders: %d, | ", completedCount)
	in.Printf("Total amount of completed orders: %.2f\n", totalCompletedAmount)

	///////////////////////////////////////////////////////////////////////////

	in.Print("\nTask 26): ")
	students := []Student1{
		{Name: "Alice", Grade: 95},
		{Name: "Bob", Grade: 82},
		{Name: "Charlie", Grade: 68},
		{Name: "Diana", Grade: 54},
		{Name: "Evan", Grade: 90},
	}

	counts := map[string]int{"A": 0, "B": 0, "C": 0, "F": 0}

	for _, s := range students {
		switch {
		case s.Grade >= 90:
			letterGrade = "A"
		case s.Grade >= 75:
			letterGrade = "B"
		case s.Grade >= 60:
			letterGrade = "C"
		default:
			letterGrade = "F"
		}

		counts[letterGrade]++
		in.Printf("%s: %s | ", s.Name, letterGrade)
	}
	in.Printf("Summary: A: %d, B: %d, C: %d, F: %d", counts["A"], counts["B"], counts["C"], counts["F"])

	/////////////////////////////////////////////////////////////////////

	in.Print("\nTask 27): ")
	products2 := []Product2{
		{Name: "Laptop", Price: 1200.0, Category: "electronics"},
		{Name: "Headphones", Price: 800.0, Category: "electronics"},
		{Name: "Jacket", Price: 150.0, Category: "clothing"},
		{Name: "Book", Price: 25.0, Category: "books"},
	}

	for _, p := range products2 {
		finalPrice := p.Price

		if p.Category == "electronics" && p.Price > 1000 {
			finalPrice *= 0.85
		} else if p.Category == "clothing" {
			finalPrice *= 0.90
		}

		discountedProducts = append(discountedProducts, ProductFinal{
			Name:       p.Name,
			Category:   p.Category,
			FinalPrice: finalPrice,
		})
	}

	for _, dp := range discountedProducts {
		in.Printf("%s (%s): $%.2f | ", dp.Name, dp.Category, dp.FinalPrice)
	}

	/////////////////////////////////////////////////////////////////////

	in.Print("\nTask 28): ")
	days := []Day{
		{Name: "Monday", TempC: 32.5},
		{Name: "Tuesday", TempC: 22.0},
		{Name: "Wednesday", TempC: 14.5},
		{Name: "Thursday", TempC: 18.0},
		{Name: "Friday", TempC: 31.0},
	}

	hot, warm, cold := 0, 0, 0

	for _, d := range days {
		if d.TempC > 30 {
			hot++
		} else if d.TempC >= 15 {
			warm++
		} else {
			cold++
		}
	}

	in.Printf("Hot days: %d | Warm days: %d | Cold days: %d", hot, warm, cold)

	/////////////////////////////////////////////////////////////////////

	in.Print("\nTask 29): ")
	matches := []Match{
		{Home: "Lions", Away: "Tigers", HomeScore: 2, AwayScore: 1},
		{Home: "Bears", Away: "Lions", HomeScore: 0, AwayScore: 0},
		{Home: "Tigers", Away: "Bears", HomeScore: 3, AwayScore: 2},
		{Home: "Lions", Away: "Bears", HomeScore: 1, AwayScore: 2},
	}

	wins := make(map[string]int)

	for _, m := range matches {
		if m.HomeScore > m.AwayScore {
			wins[m.Home]++
		} else if m.AwayScore > m.HomeScore {
			wins[m.Away]++
		}
	}

	in.Printf("Win Table: Lions: %d, Tigers: %d, Bears: %d", wins["Lions"], wins["Tigers"], wins["Bears"])

	/////////////////////////////////////////////////////////////////////

	in.Print("\nTask 30): ")
	transactions := []Transaction{
		{Type: "deposit", Amount: 100.0},
		{Type: "withdraw", Amount: 40.0},
		{Type: "withdraw", Amount: 80.0},
		{Type: "deposit", Amount: 50.0},
		{Type: "withdraw", Amount: 30.0},
	}

	balance := 0.0
	rejectedCount := 0

	for _, t := range transactions {
		if t.Type == "deposit" {
			balance += t.Amount
		} else if t.Type == "withdraw" {
			if balance-t.Amount < 0 {
				rejectedCount++
			} else {
				balance -= t.Amount
			}
		}
	}

	in.Printf("Final balance: %.2f | Rejected transactions: %d", balance, rejectedCount)

	/////////////////////////////////////////////////////////////////////

	in.Print("\nTask 31): ")
	passengers := []Passenger{
		{Name: "Sam", Age: 8},
		{Name: "Alex", Age: 25},
		{Name: "Maria", Age: 70},
		{Name: "John", Age: 12},
		{Name: "Emma", Age: 64},
	}

	totalRevenue := 0

	for _, p := range passengers {
		if p.Age < 12 {
			totalRevenue += 0
		} else if p.Age <= 64 {
			totalRevenue += 50
		} else {
			totalRevenue += 25
		}
	}

	in.Printf("Total ticket revenue: %d somoni", totalRevenue)

	/////////////////////////////////////////////////////////////////////

}

func calculatePrice(prod Product, sum float64) float64 {
	sum += prod.Price * float64(prod.Quantity)
	return sum
}

func isAdult(someone Person) string {
	if someone.Age >= 18 {
		return "This person is an adult"
	} else {
		return "This person is not an adult"
	}
}

func (c *Counter) IncrementThreeTimes() {
	for i := 0; i < 3; i++ {
		c.Value++
		in.Println(c.Value)
	}
}
