package main

import (
	in "fmt"
	"io"
	m "math"
	"os"
	"strings"
)

var (
	areaSum           float64
	totalAttackDmg    int
	distance          = 15.0
	cheapestTransport Transport
	currentPrice              = price
	price             float64 = 134.32
	smstg             string  = "Hello World"
	isWorking         bool    = true
	nums                      = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	totalArea         float64
	amph              Amphibious = Duck{Name: "Donald"}
	acc               Account    = &Wallet{}
	subscribers                  = make(map[string][]Handler)
)

type (
	//=============================== INTERFACES ===========================
	Vehicle interface { // task 1
		Start() string
		Stop() string
	}

	Shape interface { // task 2
		Area() float64
		Perimeter() float64
	}

	Employee interface { // task 3
		Salary() float64
		Position() string
	}

	Instruments interface { // task 4
		Play() string
	}

	Warrior interface { // task 5
		Attack() int
		Description() string
	}

	Transport interface { // task 6
		Price(distanceKm float64) float64
	}

	Discount interface { //task 7
		Apply(price float64) float64
	}

	Flyer interface { // task 14
		Fly() string
	}

	Swimmer interface {
		Swim() string
	}

	Amphibious interface {
		Flyer
		Swimmer
	}

	Depositor interface { // task 15
		Deposit(amount float64) error
	}

	Withdrawer interface {
		Withdraw(amount float64) error
	}

	Account interface {
		Depositor
		Withdrawer
	}

	Handler interface { // task 19
		Handle(event any)
	}

	Command interface { // task 21
		Execute() (string, error)
	}

	//====================== STRUCTS ==========================
	Bike struct{ Name string } // task 1

	Car struct{ Name string }

	Bus struct{ Name string }

	Circle struct{ Radius float64 } // task 2

	Rectangle struct{ Width, Height float64 }

	Triangle struct{ A, B, C float64 } // task 3

	Manager struct {
		Name       string
		BaseSalary float64
	}

	Developer struct {
		Name       string
		BaseSalary float64
		Level      string
	}

	Piano struct{ Name string } // task 4

	Guitar struct{ Name string }

	Drums struct{ Name string }

	Knight struct {
		Name     string
		Strength int
	} // task 5

	Archer struct {
		Name      string
		Precision int
	}

	Mage struct {
		Name      string
		ManaPower int
	}

	Taxi struct { // task 6
		BaseFare   float64
		PricePerKm float64
	}

	Bus1 struct {
		Fare float64
	}

	Bike1 struct {
		PricePerKm float64
	}

	PercentDiscount struct{ Percent float64 } // task 7
	FixedDiscount   struct{ Amount float64 }
	NoDiscount      struct{}

	Point struct{ X, Y int } // task 8

	Box struct{ Value any } // task 9

	Item struct {
		Name  string
		Price float64
	}

	Product struct { // task 13
		Name  string
		Price float64
	}

	Duck struct { // task 14
		Name string
	}

	Wallet struct { // task 15
		balance float64
	}

	Money struct { // task 16
		Amount   float64
		Currency string
	}

	Task struct { // task 17
		Title string
		Done  bool
	}

	PrintHandler struct{} // task 19

	CountHandler struct {
		count int
	}

	TextCommand struct { // task 21
		Text string
	}

	DivideCommand struct {
		A float64
		B float64
	}

	Category struct { // task 22
		Name string
		Sub  []Category
	}

	Invoice struct { // task 23
		Client string
		Amount float64
	}

	ValidationError struct { // task 24
		Field   string
		Message string
	}
)

// Task 1
func (b Bike) Start() string {
	return "Bike " + b.Name + " starts"
}
func (b Bike) Stop() string {
	return "Bike " + b.Name + " stops"

}

func (c Car) Start() string {
	return "Car " + c.Name + " starts"
}
func (c Car) Stop() string {
	return "Car " + c.Name + " stops"

}

func (bus Bus) Start() string {
	return "Bus " + bus.Name + " starts"
}
func (bus Bus) Stop() string {
	return "Bus " + bus.Name + " stops"

}

// Task 2
func (c Circle) Area() float64      { return m.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * m.Pi * c.Radius }

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

func (t Triangle) Perimeter() float64 { return t.A + t.B + t.C }
func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2.0
	area := m.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))

	return area
}

// Task 3
func (m Manager) Salary() float64  { return m.BaseSalary * 1.20 }
func (m Manager) Position() string { return "Manager" }

func (d Developer) Salary() float64 {
	if d.Level == "senior" {
		return d.BaseSalary * 1.30
	}
	return d.BaseSalary
}
func (d Developer) Position() string { return "Developer (" + d.Level + ")" }

func GetHighestPaid(employees []Employee) Employee {
	if len(employees) == 0 {
		return nil
	}

	highest := employees[0]
	for _, emp := range employees[1:] {
		if emp.Salary() > highest.Salary() {
			highest = emp
		}
	}
	return highest
}

// Task 4
func (g Guitar) Play() string { return "Guitar " + g.Name + ": chicka-chicka" }
func (p Piano) Play() string  { return "Piano " + p.Name + ": clunk-click" }
func (d Drums) Play() string  { return "Drums " + d.Name + ": boom-boom" }

// Task 5
func (k Knight) Attack() int         { return k.Strength * 2 }
func (k Knight) Description() string { return k.Name + ":" + " Knight" }

func (a Archer) Attack() int         { return a.Precision + 10 }
func (a Archer) Description() string { return a.Name + ":" + " Archer" }

func (m Mage) Attack() int         { return m.ManaPower * 3 }
func (m Mage) Description() string { return m.Name + ":" + " Mage" }

// Task 6
func (t Taxi) Price(distanceKm float64) float64 {
	return t.BaseFare + (t.PricePerKm * distanceKm)
}
func (b Bus1) Price(distanceKm float64) float64 {
	return b.Fare
}
func (bk Bike1) Price(distanceKm float64) float64 {
	return bk.PricePerKm * distanceKm
}

// Task 7
func (pd PercentDiscount) Apply(price float64) float64 {
	return price * (1.0 - pd.Percent/100.0)
}
func (fd FixedDiscount) Apply(price float64) float64 {
	newPrice := price - fd.Amount
	if newPrice < 0 {
		return 0
	}
	return newPrice
}
func (nd NoDiscount) Apply(price float64) float64 {
	return price
}

// Task 8
func Describe(v any) {
	in.Printf("%v : %T\n", v, v)
}

// / Task 10
func IsString(val any) (string, bool) {
	str, ok := val.(string)
	if ok {
		return str, true
	}
	return "", false
}

// / Task 12
func Classify(v any) string {
	switch x := v.(type) {
	case int:
		return in.Sprintf("Integer: %d", x)
	case float64:
		return in.Sprintf("Float number: %.1f", x)
	case string:
		return in.Sprintf("String with the length of: %d", len(x))
	case bool:
		return in.Sprintf("Logic expression: %v", x)
	case nil:
		return "Empty"
	default:
		return in.Sprintf("Unknown type: %T", x)
	}
}

// Task 13
func Process(v any) {
	switch x := v.(type) {
	case Product:
		in.Printf("\nName: %v | Price: %.1f", x.Name, x.Price)
	case Shape:
		in.Printf("\nThe area of the shape: %.1f", x.Area())
	default:
		in.Printf("\nIdk what to do with that type srry")
	}
}

// / Task 14
func (d Duck) Fly() string  { return in.Sprintf("%s is flying high", d.Name) }
func (d Duck) Swim() string { return in.Sprintf("%s is swimming smoothly", d.Name) }

// / Task 15
func (w *Wallet) Deposit(amount float64) error {
	if amount <= 0 {
		return in.Errorf("deposit amount must be positive, got %.2f", amount)
	}
	w.balance += amount
	return nil
}
func (w *Wallet) Withdraw(amount float64) error {
	if amount <= 0 {
		return in.Errorf("withdrawal amount must be positive, got %.2f", amount)
	}
	if amount > w.balance {
		return in.Errorf("insufficient funds: trying to withdraw %.2f, available balance %.2f", amount, w.balance)
	}
	w.balance -= amount
	return nil
}

// / Task 16
func (mon Money) String() string {
	roundedAmount := m.Round(mon.Amount*100) / 100
	return in.Sprintf("%.2f %s", roundedAmount, mon.Currency)
}

// Task 17
func (t Task) String() string {
	if t.Done {
		return in.Sprintf("[x] %s", t.Title)
	}
	return in.Sprintf("[ ] %s", t.Title)
}

// Task 18
func LogMessage(w io.Writer, msg string) {
	in.Fprintln(w, msg)
}

// Task 19
func (p *PrintHandler) Handle(event any) {
	in.Printf("[PrintHandler] Event received: %v\n", event)
}
func (c *CountHandler) Handle(event any) {
	c.count++
	in.Printf("[CountHandler] Event received: %v (Total handled: %d)\n", event, c.count)
}
func Subscribe(eventName string, h Handler) {
	subscribers[eventName] = append(subscribers[eventName], h)
}
func Publish(eventName string, event any) {
	if handlers, ok := subscribers[eventName]; ok {
		for _, h := range handlers {
			h.Handle(event)
		}
	}
}

// Task 21
func (t TextCommand) Execute() (string, error) {
	return t.Text, nil
}
func (d DivideCommand) Execute() (string, error) {
	if d.B == 0 {
		return "", in.Errorf("division by zero")
	}
	return in.Sprintf("%.2f", d.A/d.B), nil
}

// Task 22
func PrintTree(v any, depth int) {
	indent := strings.Repeat("  ", depth)
	switch x := v.(type) {
	case Category:
		in.Printf("%s%s\n", indent, x.Name)
		PrintTree(x.Sub, depth+1)
	case []Category:
		for _, cat := range x {
			PrintTree(cat, depth)
		}
	}
}

// Task 23
func (inv Invoice) String() string {
	return in.Sprintf("%s: %.2f", inv.Client, inv.Amount)
}
func WriteReport(w io.Writer, invoices []Invoice) float64 {
	var total float64
	for _, inv := range invoices {
		in.Fprintln(w, inv.String())
		total += inv.Amount
	}
	return total
}

// Task 24
func (e *ValidationError) Error() string {
	return in.Sprintf("field %s: %s", e.Field, e.Message)
}

func ValidateAge(age int) error {
	if age < 0 || age > 130 {
		return &ValidationError{
			Field:   "Age",
			Message: "must be between 0 and 130",
		}
	}
	return nil
}

// /////////////////////////// === MAIN === //////////////////////////////
func main() {

	//////////////////=== TASK 1 ===//////////////////////////////
	vehicles := []Vehicle{
		Car{Name: "Toyota"},
		Bike{Name: "Trek"},
		Bus{Name: "Mercedes-Benz"},
		Car{Name: "BMW"},
	}

	in.Println("=== Task 1 ===")
	for i, v := range vehicles {
		in.Print(i+1, ")", v.Start())
		in.Print(", ", v.Stop(), "\n")
	}
	//////////////////////////////////////////////////////

	shapes := []Shape{
		Circle{Radius: 6.3},
		Rectangle{Width: 43.4, Height: 38.8},
		Triangle{A: 33.0, B: 21.3, C: 35.9},
		Rectangle{Width: 4.7, Height: 3.8},
		Circle{Radius: 9.7},
	}

	for _, v := range shapes {
		areaSum += v.Area()
	}
	in.Println("\n=== Task 2 ===")
	in.Printf("The total area of %v shapes: %.3f\n", len(shapes), areaSum)

	///////////////////////////////////////////////////////

	in.Println("\n=== Task 3 ===")
	employees := []Employee{
		Manager{Name: "Alex", BaseSalary: 100000},
		Developer{Name: "Ivan", BaseSalary: 80000, Level: "junior"},
		Developer{Name: "Elisa", BaseSalary: 110000, Level: "senior"},
		Manager{Name: "Origa", BaseSalary: 115000},
	}

	topEmployee := GetHighestPaid(employees)

	if topEmployee != nil {
		in.Printf("Top Earner:\n")
		in.Printf("Position: %s\n", topEmployee.Position())
		in.Printf("Salary: %.2f\n", topEmployee.Salary())
	}

	///////////////////////////////////////////////////////

	insts := []Instruments{
		Piano{"Yamaha"},
		Guitar{"Fender Telecaster"},
		Drums{"Tama"},
		Piano{"Steinway & Sons"},
		Guitar{"Fender Stratocaster"},
	}

	in.Println("\n=== Task 4 ===")
	for _, v := range insts {
		in.Println(v.Play())
	}

	/////////////////////////////////////////////////////

	wrrs := []Warrior{
		Archer{"Wind Archer Cookie", 900},
		Mage{"Timekeeper Cookie", 590},
		Knight{"Tea Knight Cookie", 147},
		Knight{"Knight Cookie", 65},
		Mage{"Wizard Cookie", 89},
		Mage{"Pure Vanilla Cookie", 788},
		Archer{"Golden Cheese Cookie", 550},
	}

	in.Println("\n=== Task 5 ===")
	for i, v := range wrrs {
		totalAttackDmg += v.Attack()
		in.Println(i+1, ")", v.Description())
	}
	in.Printf("\nTotal damage: %d", totalAttackDmg)

	//////////////////////////////////////////////////////////////////////////////////////////

	transports := []Transport{
		Taxi{BaseFare: 100.0, PricePerKm: 30.0},
		Bus1{Fare: 50.0},
		Bike1{PricePerKm: 15.0},
	}

	minPrice := m.MaxFloat64

	in.Println("\n=== Task 6 ===")
	in.Printf("Calculating prices for a trip of %.1f km:\n", distance)

	for _, t := range transports {
		currentPrice = t.Price(distance)

		in.Printf("- %T: %.2f\n", t, currentPrice)

		if currentPrice < minPrice {
			minPrice = currentPrice
			cheapestTransport = t
		}
	}

	in.Printf("\nCheapest option: %T (Cost: %.2f)\n", cheapestTransport, minPrice)

	/////////////////////////////////////////////////////////////////////////////////////////////
	currentPrice = price
	in.Println("\n=== Task 7 ===")

	discounts := []Discount{
		PercentDiscount{Percent: 10.0},
		FixedDiscount{Amount: 150.0},
		NoDiscount{},
		PercentDiscount{Percent: 5.0},
	}

	in.Printf("Initial price: $%.2f\n\n", price)

	for i, d := range discounts {
		currentPrice = d.Apply(currentPrice)
		in.Printf("Step %d (%T): $%.2f\n", i+1, d, currentPrice)
	}

	in.Printf("\nFinal price: $%.2f\n", currentPrice)

	////////////////////////////////////////////////////////////////////////////////////////////
	in.Println("\n=== Task 8 ===")

	Describe(totalAttackDmg)
	Describe(price)
	Describe(smstg)
	Describe(isWorking)
	Describe(nums)
	Describe(Point{12, 43})

	///////////////////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 9 ===")

	boxes := []Box{
		{Value: 42},
		{Value: "Golang Interfaces"},
		{Value: 3.14159},
		{Value: []int{10, 20, 30}},
		{Value: Item{Name: "Book", Price: 15.99}},
		{Value: true},
	}

	for i, b := range boxes {
		in.Printf("%d) Content: %v | Type: %T\n", i+1, b.Value, b.Value)
	}

	/////////////////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 10 ===")

	in.Println(IsString("Hello"))
	in.Println(IsString(123))
	in.Println(IsString(true))

	/////////////////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 11 ===")

	mixedSlice := []any{
		"Hello World",
		Circle{Radius: 5},
		42,
		Rectangle{Width: 4, Height: 6},
		true,
		Circle{Radius: 2.5},
	}

	for _, item := range mixedSlice {
		if shape, ok := item.(Shape); ok {
			totalArea += shape.Area()

		}
	}

	in.Printf("Total area of shapes: %.2f\n", totalArea)

	/////////////////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 12 ===")

	randvs := []any{12, 56.7, true, "haha", Circle{Radius: 5}}

	for _, v := range randvs {
		in.Println(Classify(v))
	}

	/////////////////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 13 ===")

	anyValues := []any{
		Product{"original Airpods trust", 10.0},
		Circle{Radius: 67},
		12.4,
		23,
		676767,
		"I'm a product",
		Product{"Sonic Plushie", 25.5},
		Rectangle{Width: 4, Height: 6},
	}

	for _, v := range anyValues {
		Process(v)
	}

	/////////////////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 14 ===")

	in.Println(amph.Fly())
	in.Println(amph.Swim())

	/////////////////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 15 ===")

	in.Println("Depositing 150.00...")
	if err := acc.Deposit(150.00); err != nil {
		in.Println("Error:", err)
	}

	in.Println("Withdrawing 50.00...")
	if err := acc.Withdraw(50.00); err != nil {
		in.Println("Error:", err)
	}

	in.Println("Attempting invalid withdrawal (200.00)...")
	if err := acc.Withdraw(200.00); err != nil {
		in.Println("Error:", err)
	}

	/////////////////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 16 ===")

	currencies := []Money{
		{Amount: 150.00, Currency: "USD"},
		{Amount: 99.99, Currency: "EUR"},
		{Amount: 1250.50, Currency: "TJS"},
	}

	for _, mon := range currencies {
		in.Println(mon)
	}

	/////////////////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 17 ===")

	tasks := []Task{
		{Title: "Learn Go Interfaces", Done: true},
		{Title: "Implement Stringer interface", Done: true},
		{Title: "Write unit tests", Done: false},
		{Title: "Refactor legacy code", Done: false},
		{Title: "Submit pull request", Done: false},
	}
	for _, t := range tasks {
		in.Println(t)
	}

	/////////////////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 18 ===")

	LogMessage(os.Stdout, "This message is printed directly to stdout.")

	var builder strings.Builder
	LogMessage(&builder, "This message is captured in strings.Builder.")

	in.Print("Captured output in builder: ", builder.String())

	/////////////////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 19 ===")

	printH := &PrintHandler{}
	countH := &CountHandler{}

	Subscribe("user_logged_in", printH)
	Subscribe("user_logged_in", countH)

	Publish("user_logged_in", "User 'Alice'")
	Publish("user_logged_in", 1001)

	/////////////////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 20 ===")

	shapes1 := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 4, Height: 5},
		Triangle{A: 3, B: 4, C: 5},
		Circle{Radius: 5},
	}

	var maxShape Shape
	var maxArea float64 = -1

	for _, s := range shapes1 {
		area := s.Area()
		if area > maxArea {
			maxArea = area
			maxShape = s
		}
	}

	in.Printf("Max Area: %.2f\nDetails: ", maxArea)
	switch s := maxShape.(type) {
	case Circle:
		in.Printf("Circle with Radius: %.2f\n", s.Radius)
	case Rectangle:
		in.Printf("Rectangle with Width: %.2f, Height: %.2f\n", s.Width, s.Height)
	case Triangle:
		in.Printf("Triangle with Sides: %.2f, %.2f, %.2f\n", s.A, s.B, s.C)
	}

	/////////////////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 21 ===")

	registry := map[string]Command{
		"greet":  TextCommand{Text: "Hello, World!"},
		"div10":  DivideCommand{A: 10, B: 2},
		"divErr": DivideCommand{A: 10, B: 0},
	}

	commandsToRun := []string{"greet", "div10", "divErr", "unknown"}

	for _, name := range commandsToRun {
		if cmd, ok := registry[name]; ok {
			res, err := cmd.Execute()
			if err != nil {
				in.Printf("Command '%s' failed with error: %v\n", name, err)
			} else {
				in.Printf("Command '%s' executed successfully: %s\n", name, res)
			}
		} else {
			in.Printf("Command '%s' not found in registry\n", name)
		}
	}

	/////////////////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 22 ===")

	tree := Category{
		Name: "Electronics",
		Sub: []Category{
			{
				Name: "Computers",
				Sub: []Category{
					{Name: "Laptops"},
					{Name: "Desktops"},
				},
			},
			{
				Name: "Phones",
				Sub: []Category{
					{Name: "Smartphones"},
				},
			},
		},
	}

	PrintTree(tree, 0)

	/////////////////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 23 ===")

	invoices := []Invoice{
		{Client: "Client A", Amount: 1500.00},
		{Client: "Client B", Amount: 2500.50},
	}

	in.Println("--- Writing to os.Stdout ---")
	total1 := WriteReport(os.Stdout, invoices)
	in.Printf("Returned Total: %.2f\n\n", total1)

	total2 := WriteReport(&builder, invoices)
	in.Println("--- Captured in Builder ---")
	in.Print(builder.String())
	in.Printf("Returned Total: %.2f\n\n", total2)

	/////////////////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 24 ===")

	ages := []int{25, -5, 150}

	for _, age := range ages {
		err := ValidateAge(age)
		if err != nil {
			if valErr, ok := err.(*ValidationError); ok {
				in.Printf("Validation failed! Field: '%s' | Error Text: '%s'\n", valErr.Field, valErr.Error())
			} else {
				in.Printf("General error: %v\n", err)
			}
		} else {
			in.Printf("Age %d is valid.\n", age)
		}
	}
}
