package main

import (
	"errors"
	in "fmt"
	"math"
)

var (
	// Task 1
	wallet Wallet

	// Task 2
	playlist Playlist

	// Task 3
	inv Inventory

	// Task 4
	mat Matrix

	// Task 5
	poly Polygon

	// Task 6
	company Company

	// Task 7
	q OrderQueue

	// Task 8
	animals []SoundMaker

	// Task 9
	student9 Student9

	// Task 10
	head10 *Node

	// Task 11
	students11  []Student11
	headStudent Student11

	// Task 12
	acc1, acc2 BankAccount

	// Task 13
	tasks []Task

	// Task 14
	idGen func() int
	users []User

	// Task 15
	rootCategory Category

	// Task 16
	book AddressBook

	// Task 17
	p1, p2 Character

	// Task 18
	bot Robot

	// Task 19
	light TrafficLight

	// Task 20
	stack Stack

	// Task 21
	players []Player

	// Task 22
	canvas Canvas

	// Task 23
	cfg *Config

	// Task 24
	rect    *Rectangle
	rectErr error

	// Task 25
	queue25 *Queue

	// Task 26
	car1, car2 *Car

	// Task 27
	logger1, logger2 *Logger
)

type (
	// Task 1
	Wallet struct {
		balance float64
	}

	// Task 2
	Song struct {
		Title    string
		Duration int
	}

	Playlist struct {
		songs []Song
	}

	// Task 3
	Inventory struct {
		items map[string]int
	}

	// Task 4
	Matrix [][]float64

	// Task 5
	Point struct {
		X float64
		Y float64
	}

	Polygon []Point

	// Task 6
	Employee struct {
		Name        string
		Salary      float64
		YearsWorked int
	}

	Company []Employee

	// Task 7
	Order struct {
		ID     int
		Amount float64
	}

	OrderQueue struct {
		orders []Order
	}

	// Task 8
	SoundMaker interface {
		MakeSound() string
	}

	Animal struct {
		Name  string
		Sound string
	}

	Dog struct {
		Animal
	}

	Cat struct {
		Animal
	}

	// Task 9
	Person struct {
		Name string
		Age  int
	}

	Student9 struct {
		Person
		University string
	}

	// Task 10
	Node struct {
		Value int
		Next  *Node
	}

	// Task 11
	Student11 struct {
		Name   string
		Grades []int
	}

	// Task 12
	BankAccount struct {
		Owner   string
		balance float64
	}

	// Task 13
	Task struct {
		Title    string
		Priority int
		Done     bool
	}

	// Task 14
	User struct {
		ID   int
		Name string
	}

	// Task 15
	Category struct {
		Name string
		Sub  []Category
	}

	// Task 16
	Contact struct {
		Name  string
		Phone string
	}

	AddressBook struct {
		contacts map[string]Contact
	}

	// Task 17
	Character struct {
		Name   string
		HP     int
		Attack int
	}

	// Task 18
	Robot struct {
		x int
		y int
	}

	// Task 19
	TrafficLight struct {
		color string
	}

	// Task 20
	Stack struct {
		data []int
	}

	// Task 21
	Player struct {
		Name  string
		Score int
	}

	// Task 22
	Canvas struct {
		grid [][]rune
	}

	// Task 23
	Config struct {
		Host    string
		Port    int
		Timeout int
	}

	// Task 24
	Rectangle struct {
		width  float64
		height float64
	}

	// Task 25
	Queue struct {
		elements []int
	}

	// Task 26
	Car struct {
		Brand   string
		Model   string
		mileage int
	}

	// Task 27
	Logger struct {
		prefix string
		count  int
	}
)

// Task 1 Methods
func (w *Wallet) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("cannot fill with negative number")
	}
	w.balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount > w.balance {
		return errors.New("Not enough recources")
	}
	w.balance -= amount
	return nil
}

func (w Wallet) Balance() float64 {
	return w.balance
}

// Task 2 Methods
func (p *Playlist) Add(s Song) {
	p.songs = append(p.songs, s)
}

func (p Playlist) TotalDuration() int {
	total := 0
	for _, s := range p.songs {
		total += s.Duration
	}
	return total
}

func (p Playlist) Longest() Song {
	if len(p.songs) == 0 {
		return Song{}
	}
	longest := p.songs[0]
	for _, s := range p.songs[1:] {
		if s.Duration > longest.Duration {
			longest = s
		}
	}
	return longest
}

func (p Playlist) String() string {
	res := ""
	for i, s := range p.songs {
		res += in.Sprintf("%d. %s (%d sec)\n", i+1, s.Title, s.Duration)
	}
	return res
}

// Task 3 Methods
func (inv *Inventory) Add(item string, qty int) {
	if inv.items == nil {
		inv.items = make(map[string]int)
	}
	inv.items[item] += qty
}

func (inv *Inventory) Remove(item string, qty int) error {
	if inv.items == nil {
		return errors.New("the item is not in the inventory")
	}
	curr, exists := inv.items[item]
	if !exists || curr < qty {
		return errors.New("the item is not enough or it doesn't exist")
	}
	inv.items[item] -= qty
	if inv.items[item] == 0 {
		delete(inv.items, item)
	}
	return nil
}

func (inv Inventory) Count(item string) int {
	if inv.items == nil {
		return 0
	}
	return inv.items[item]
}

func (inv Inventory) Total() int {
	sum := 0
	for _, qty := range inv.items {
		sum += qty
	}
	return sum
}

// Task 4 Methods
func (m Matrix) Rows() int {
	return len(m)
}

func (m Matrix) Cols() int {
	if len(m) == 0 {
		return 0
	}
	return len(m[0])
}

func (m Matrix) Sum() float64 {
	total := 0.0
	for _, row := range m {
		for _, val := range row {
			total += val
		}
	}
	return total
}

func (m Matrix) Transpose() Matrix {
	rows := m.Rows()
	cols := m.Cols()
	res := make(Matrix, cols)
	for i := 0; i < cols; i++ {
		res[i] = make([]float64, rows)
		for j := 0; j < rows; j++ {
			res[i][j] = m[j][i]
		}
	}
	return res
}

// Task 5 Methods
func (p Point) DistanceTo(other Point) float64 {
	dx := p.X - other.X
	dy := p.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (poly Polygon) Perimeter() float64 {
	if len(poly) < 2 {
		return 0
	}
	perimeter := 0.0
	for i := 0; i < len(poly); i++ {
		nextIndex := (i + 1) % len(poly)
		perimeter += poly[i].DistanceTo(poly[nextIndex])
	}
	return perimeter
}

// Task 6 Methods
func (e Employee) Bonus() float64 {
	periods := e.YearsWorked / 5
	return e.Salary * 0.10 * float64(periods)
}

func (c Company) TotalPayroll() float64 {
	total := 0.0
	for _, emp := range c {
		total += emp.Salary + emp.Bonus()
	}
	return total
}

func (c Company) TopEarner() Employee {
	if len(c) == 0 {
		return Employee{}
	}
	top := c[0]
	topTotal := top.Salary + top.Bonus()
	for _, emp := range c[1:] {
		currTotal := emp.Salary + emp.Bonus()
		if currTotal > topTotal {
			top = emp
			topTotal = currTotal
		}
	}
	return top
}

// Task 7 Methods
func (oq *OrderQueue) Enqueue(o Order) {
	oq.orders = append(oq.orders, o)
}

func (oq *OrderQueue) Dequeue() (Order, error) {
	if len(oq.orders) == 0 {
		return Order{}, errors.New("Queue is empty")
	}
	first := oq.orders[0]
	oq.orders = oq.orders[1:]
	return first, nil
}

func (oq OrderQueue) IsEmpty() bool {
	return len(oq.orders) == 0
}

func (oq OrderQueue) TotalAmount() float64 {
	sum := 0.0
	for _, o := range oq.orders {
		sum += o.Amount
	}
	return sum
}

// Task 8 Methods
func (a Animal) MakeSound() string {
	return in.Sprintf("%s says %s", a.Name, a.Sound)
}

func (d Dog) MakeSound() string {
	return d.Animal.MakeSound() + " (wiggles tail)"
}

func (c Cat) MakeSound() string {
	return c.Animal.MakeSound() + " (purrs)"
}

// Task 9 Methods
func (p Person) Info() string {
	return in.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func (s Student9) Info() string {
	return in.Sprintf("%s, University: %s", s.Person.Info(), s.University)
}

// Task 10 Methods & Functions
func (head *Node) Append(v int) *Node {
	newNode := &Node{Value: v}
	if head == nil {
		return newNode
	}
	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
	return head
}

func PrintList(head *Node) {
	curr := head
	for curr != nil {
		in.Printf("%d", curr.Value)
		if curr.Next != nil {
			in.Printf(" -> ")
		}
		curr = curr.Next
	}
	in.Println()
}

// Task 11 Methods & Functions
func (s Student11) Average() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	sum := 0
	for _, g := range s.Grades {
		sum += g
	}
	return float64(sum) / float64(len(s.Grades))
}

func (s Student11) String() string {
	return in.Sprintf("%s (Average: %.2f)", s.Name, s.Average())
}

// Task 12 Methods & Functions
func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount
}

func (b *BankAccount) Withdraw(amount float64) error {
	if amount > b.balance {
		return errors.New("not enough recourses")
	}
	b.balance -= amount
	return nil
}

func Transfer(from, to *BankAccount, amount float64) error {
	err := from.Withdraw(amount)
	if err != nil {
		return err
	}
	to.Deposit(amount)
	return nil
}

// Task 13 Methods
func (t *Task) Complete() {
	t.Done = true
}

// Task 14 Functions
func NewIDGenerator() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

// Task 15 Functions
func CountAll(c Category) int {
	total := 1
	for _, sub := range c.Sub {
		total += CountAll(sub)
	}
	return total
}

// Task 16 Methods
func (ab *AddressBook) Add(key string, c Contact) {
	if ab.contacts == nil {
		ab.contacts = make(map[string]Contact)
	}
	ab.contacts[key] = c
}

func (ab AddressBook) Find(key string) (Contact, bool) {
	c, exists := ab.contacts[key]
	return c, exists
}

func (ab *AddressBook) Delete(key string) {
	delete(ab.contacts, key)
}

// Task 17 Methods
func (c Character) IsAlive() bool {
	return c.HP > 0
}

func (c *Character) TakeDamage(dmg int) {
	c.HP -= dmg
	if c.HP < 0 {
		c.HP = 0
	}
}

func (c *Character) AttackTarget(target *Character) {
	target.TakeDamage(c.Attack)
}

// Task 18 Methods
func (r *Robot) MoveUp()    { r.y++ }
func (r *Robot) MoveDown()  { r.y-- }
func (r *Robot) MoveLeft()  { r.x-- }
func (r *Robot) MoveRight() { r.x++ }
func (r Robot) Position() (int, int) {
	return r.x, r.y
}
func (r Robot) String() string {
	return in.Sprintf("Robot @ (%d, %d)", r.x, r.y)
}

// Task 19 Methods
func (t *TrafficLight) Next() {
	switch t.color {
	case "Red":
		t.color = "Yellow"
	case "Yellow":
		t.color = "Yellow"
	case "Green":
		t.color = "Red"
	default:
		t.color = "Red"
	}
}

func (t TrafficLight) Current() string {
	if t.color == "" {
		return "Red"
	}
	return t.color
}

func (t TrafficLight) String() string {
	return in.Sprintf("Traffic light [%s]", t.Current())
}

// Task 20 Methods
func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("Stack is empty")
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val, nil
}

func (s Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// Task 21 Methods & Functions
func (p *Player) AddPoints(pts int) {
	p.Score += pts
}

func TopPlayer(players []Player) Player {
	if len(players) == 0 {
		return Player{}
	}
	top := players[0]
	for _, p := range players[1:] {
		if p.Score > top.Score {
			top = p
		}
	}
	return top
}

// Task 22 Methods
func NewCanvas(width, height int) Canvas {
	g := make([][]rune, height)
	for i := range g {
		g[i] = make([]rune, width)
		for j := range g[i] {
			g[i][j] = ' '
		}
	}
	return Canvas{grid: g}
}

func (c *Canvas) SetPixel(x, y int, char rune) error {
	if y < 0 || y >= len(c.grid) || x < 0 || x >= len(c.grid[0]) {
		return errors.New("the coordinates is out of canvas")
	}
	c.grid[y][x] = char
	return nil
}

func (c Canvas) Print() {
	for _, row := range c.grid {
		for _, char := range row {
			in.Printf("%c", char)
		}
		in.Println()
	}
}

// Task 23 Constructors
func NewConfig(host string, port int) *Config {
	return &Config{
		Host:    host,
		Port:    port,
		Timeout: 30,
	}
}

// Task 24 Constructors & Methods
func NewRectangle(w, h float64) (*Rectangle, error) {
	if w <= 0 || h <= 0 {
		return nil, errors.New("the faces of a rectangle must be positive")
	}
	return &Rectangle{width: w, height: h}, nil
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Task 25 Constructors & Methods
func NewQueue() *Queue {
	return &Queue{elements: make([]int, 0)}
}

func (q *Queue) Enqueue(v int) {
	q.elements = append(q.elements, v)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.elements) == 0 {
		return 0, errors.New("Queue is empty")
	}
	val := q.elements[0]
	q.elements = q.elements[1:]
	return val, nil
}

// Task 26 Constructors & Methods
func NewCar(brand, model string) *Car {
	return &Car{Brand: brand, Model: model, mileage: 0}
}

func NewUsedCar(brand, model string, mileage int) *Car {
	return &Car{Brand: brand, Model: model, mileage: mileage}
}

func (c *Car) Drive(km int) {
	if km > 0 {
		c.mileage += km
	}
}

func (c Car) Mileage() int {
	return c.mileage
}

// Task 27 Constructors & Methods
func NewLogger(prefix string) *Logger {
	return &Logger{prefix: prefix, count: 0}
}

func (l *Logger) Log(msg string) {
	l.count++
	in.Printf("[%s] %s (#%d)\n", l.prefix, msg, l.count)
}

func main() {
	in.Println("=== TASK 1 ===")
	wallet = Wallet{}
	_ = wallet.Deposit(100)
	in.Printf("Balance: %.2f\n", wallet.Balance())
	_ = wallet.Deposit(50)
	in.Printf("Balance: %.2f\n", wallet.Balance())
	if err := wallet.Withdraw(200); err != nil {
		in.Printf("Error: %v\n", err)
	}
	_ = wallet.Withdraw(30)
	in.Printf("Balance: %.2f\n\n", wallet.Balance())

	in.Println("=== TASK 2 ===")
	playlist.Add(Song{"Song 1", 180})
	playlist.Add(Song{"Song 2", 240})
	playlist.Add(Song{"Song 3", 150})
	playlist.Add(Song{"Song 4", 310})
	in.Printf("List of songs:\n%s", playlist)
	in.Printf("Total duration: %d sec\n", playlist.TotalDuration())
	in.Printf("The longest song: %s (%d sec)\n\n", playlist.Longest().Title, playlist.Longest().Duration)

	in.Println("=== TASK 3 ===")
	inv.Add("Apples", 10)
	inv.Add("Pears", 5)
	in.Printf("Apples in storage: %d\n", inv.Count("Apples"))
	in.Printf("In total: %d\n", inv.Total())
	if err := inv.Remove("Apples", 15); err != nil {
		in.Printf("Error: %v\n\n", err)
	}

	in.Println("=== TASK 4 ===")
	mat = Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	in.Printf("Rows: %d, Colums: %d\n", mat.Rows(), mat.Cols())
	in.Printf("Element Summary: %.2f\n", mat.Sum())
	in.Printf("Transposed: %+v\n\n", mat.Transpose())

	in.Println("=== TASK 5 ===")
	poly = Polygon{{X: 0, Y: 0}, {X: 0, Y: 3}, {X: 4, Y: 0}}
	in.Printf("Perimeter of the triangle: %.2f\n\n", poly.Perimeter())

	in.Println("=== TASK 6 ===")
	company = Company{
		{"Alan", 50000, 6},
		{"Jin", 70000, 12},
		{"Kazuya", 60000, 3},
	}
	in.Printf("Total Payroll: %.2f\n", company.TotalPayroll())
	in.Printf("Top earner: %s\n\n", company.TopEarner().Name)

	in.Println("=== TASK 7 ===")
	q.Enqueue(Order{ID: 1, Amount: 150.5})
	q.Enqueue(Order{ID: 2, Amount: 299.5})
	in.Printf("Total orders: %.2f\n", q.TotalAmount())
	for !q.IsEmpty() {
		o, _ := q.Dequeue()
		in.Printf("Made order #%d\n", o.ID)
	}
	_, errQ := q.Dequeue()
	in.Printf("Error: %v\n\n", errQ)

	in.Println("=== TASK 8 ===")
	animals = []SoundMaker{
		Dog{Animal{Name: "Rex", Sound: "Woof"}},
		Cat{Animal{Name: "Liah", Sound: "Meow"}},
	}
	for _, a := range animals {
		in.Println(a.MakeSound())
	}
	in.Println()

	in.Println("=== TASK 9 ===")
	student9 = Student9{
		Person:     Person{Name: "Alexey", Age: 20},
		University: "MSU",
	}
	in.Printf("%s\n\n", student9.Info())

	in.Println("=== TASK 10 ===")
	for _, v := range []int{10, 20, 30, 40, 50} {
		head10 = head10.Append(v)
	}
	in.Print("List: ")
	PrintList(head10)
	in.Println()

	in.Println("=== TASK 11 ===")
	students11 = []Student11{
		{"Ivan", []int{4, 5, 4}},
		{"Olga", []int{5, 5, 5}},
		{"Oleg", []int{3, 4, 3}},
	}
	headStudent = students11[0]
	for _, s := range students11[1:] {
		if s.Average() > headStudent.Average() {
			headStudent = s
		}
	}
	in.Printf("Class monitor: %s\n\n", headStudent)

	in.Println("=== TASK 12 ===")
	acc1 = BankAccount{Owner: "Alice", balance: 100}
	acc2 = BankAccount{Owner: "Bob", balance: 50}
	if err := Transfer(&acc1, &acc2, 120); err != nil {
		in.Printf("Transfer error: %v\n", err)
	}
	_ = Transfer(&acc1, &acc2, 40)
	in.Printf("Transfer successful. Alice balance: %.2f, Bob: %.2f\n\n", acc1.balance, acc2.balance)

	in.Println("=== TASK 13 ===")
	tasks = []Task{
		{"Task A", 3, false},
		{"Task B", 1, false},
		{"Task C", 2, false},
	}
	in.Printf("Before sorting: %+v\n", tasks)
	for i := 0; i < len(tasks); i++ {
		for j := 0; j < len(tasks)-1-i; j++ {
			if tasks[j].Priority > tasks[j+1].Priority {
				tasks[j], tasks[j+1] = tasks[j+1], tasks[j]
			}
		}
	}
	in.Printf("After sorting: %+v\n\n", tasks)

	in.Println("=== TASK 14 ===")
	idGen = NewIDGenerator()
	users = []User{
		{ID: idGen(), Name: "User1"},
		{ID: idGen(), Name: "User2"},
	}
	in.Printf("Users: %+v\n\n", users)

	in.Println("=== TASK 15 ===")
	rootCategory = Category{
		Name: "Electronics",
		Sub: []Category{
			{Name: "Phones", Sub: []Category{{Name: "Smartphones"}}},
			{Name: "Laptops"},
		},
	}
	in.Printf("Total categories in tree: %d\n\n", CountAll(rootCategory))

	in.Println("=== TASK 16 ===")
	book = AddressBook{}
	book.Add("friend", Contact{Name: "Dima", Phone: "12345"})
	if c, found := book.Find("friend"); found {
		in.Printf("Contact found: %+v\n", c)
	}
	book.Delete("friend")
	_, found := book.Find("friend")
	in.Printf("Contact found after deletion? %v\n\n", found)

	in.Println("=== TASK 17 ===")
	p1 = Character{Name: "Knight", HP: 30, Attack: 8}
	p2 = Character{Name: "Goblin", HP: 20, Attack: 5}
	for p1.IsAlive() && p2.IsAlive() {
		p1.AttackTarget(&p2)
		in.Printf("%s hits %s! %s has %d HP remaining\n", p1.Name, p2.Name, p2.Name, p2.HP)
		if p2.IsAlive() {
			p2.AttackTarget(&p1)
			in.Printf("%s hits %s! %s has %d HP remaining\n", p2.Name, p1.Name, p1.Name, p1.HP)
		}
	}
	if p1.IsAlive() {
		in.Printf("Winner: %s\n\n", p1.Name)
	} else {
		in.Printf("Winner: %s\n\n", p2.Name)
	}

	in.Println("=== TASK 18 ===")
	bot = Robot{}
	bot.MoveRight()
	bot.MoveRight()
	bot.MoveUp()
	in.Printf("%s\n\n", bot)

	in.Println("=== TASK 19 ===")
	light = TrafficLight{}
	for i := 0; i < 4; i++ {
		in.Println(light)
		light.Next()
	}
	in.Println()

	in.Println("=== TASK 20 ===")
	stack = Stack{}
	stack.Push(5)
	stack.Push(10)
	val, _ := stack.Pop()
	in.Printf("Popped from stack: %d\n", val)
	_, _ = stack.Pop()
	_, errStack := stack.Pop()
	in.Printf("Error on Pop from empty stack: %v\n\n", errStack)

	in.Println("=== TASK 21 ===")
	players = []Player{{"Player 1", 10}, {"Player 2", 25}, {"Player 3", 15}}
	in.Printf("Winner: %s with score %d\n\n", TopPlayer(players).Name, TopPlayer(players).Score)

	in.Println("=== TASK 22 ===")
	canvas = NewCanvas(8, 4)
	_ = canvas.SetPixel(0, 0, '*')
	_ = canvas.SetPixel(1, 1, '*')
	_ = canvas.SetPixel(2, 2, '*')
	canvas.Print()
	in.Println()

	in.Println("=== TASK 23 ===")
	cfg = NewConfig("localhost", 8080)
	in.Printf("Config: %+v\n\n", cfg)

	in.Println("=== TASK 24 ===")
	_, rectErr = NewRectangle(-5, 10)
	in.Printf("Creation error: %v\n", rectErr)
	rect, _ = NewRectangle(5, 10)
	in.Printf("Area: %.2f, Perimeter: %.2f\n\n", rect.Area(), rect.Perimeter())

	in.Println("=== TASK 25 ===")
	queue25 = NewQueue()
	queue25.Enqueue(100)
	valQ, _ := queue25.Dequeue()
	in.Printf("Dequeued immediately after creation: %d\n\n", valQ)

	in.Println("=== TASK 26 ===")
	car1 = NewCar("Toyota", "Camry")
	car2 = NewUsedCar("BMW", "X5", 50000)
	car1.Drive(150)
	in.Printf("Car1 Mileage: %d km, Car2 Mileage: %d km\n\n", car1.Mileage(), car2.Mileage())

	in.Println("=== TASK 27 ===")
	logger1 = NewLogger("APP")
	logger2 = NewLogger("DB")
	logger1.Log("Application launch")
	logger2.Log("Database connection")
	logger1.Log("Request processing")
}
