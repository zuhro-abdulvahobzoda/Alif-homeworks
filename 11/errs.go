package main

import (
	err "errors"
	in "fmt"
	m "math"
	str "strconv"
	"unicode"
)

var (
	ErrNegativeAge = err.New("Age cannot be negative") // task 5
	ErrTooOld      = err.New("That seems like you're too old")

	ErrEmptyLogin    = err.New("Login is required") // task 6
	ErrShortPassword = err.New("The password must contain at least 6 characters")

	ErrNotFound = err.New("Item not found") // task 7

	ErrCantParse = err.New("couldn't figure out the age") // task 9

	ErrUserExists = err.New("user already exists") // task 10
)

// Task 1
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, err.New("You can't take the square root of a negative number ")
	}
	return m.Sqrt(x), nil
}

// Task 2
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, in.Errorf("Cannot divide %.1f with %.1f", a, b)
	}
	return a / b, nil
}

// Task 3
func Withdraw(balance, amount float64) (float64, error) {
	if amount <= 0 {
		return 0, in.Errorf("The amount cannot be negative")
	} else if balance < amount {
		return 0, in.Errorf("Not enough recources")
	}

	return balance - amount, nil
}

// Task 4
func step1(n *int) error {
	if *n%2 != 0 {
		return err.New("This number cannot be divided by 2")
	}
	return nil
}
func step2(n *int) error {
	*n += 1
	if *n < 999 {
		return err.New("The number cannot be bigger than 999")
	}
	return nil
}
func step3(n *int) error {
	*n -= 6
	if *n%3 == 0 {
		return err.New("The number can be divided by zero")
	}
	return nil
}
func Pipeline(n *int) error {

	if err := step1(n); err != nil {
		return err
	}

	if err := step2(n); err != nil {
		return err
	}

	if err := step3(n); err != nil {
		return err
	}

	return nil
}

// Task 5
func ValidateAge(age int) error {
	if age < 0 {
		return ErrNegativeAge
	}
	if age > 120 {
		return ErrTooOld
	}
	return nil
}

// Task 6
func CheckCredentials(login, password string) error {
	if login == "" {
		return ErrEmptyLogin
	}

	if len(password) < 6 {
		return ErrShortPassword
	}
	return nil
}

// Task 7
func FindItem(id int) error {
	if id < 0 {
		return ErrNotFound
	}
	return nil
}
func loadItem(id int) error {
	err3 := FindItem(id)

	if err3 != nil {
		in.Errorf("The item %d is not loaded: %w", id, err3)
	}
	return nil
}

// Task 8
func readConfig() error {
	return err.New("config file not found")
}
func startServer() error {
	err1 := readConfig()
	if err1 != nil {
		return in.Errorf("failed to start server: %w", err1)
	}
	return nil
}
func run() error {
	err1 := startServer()
	if err1 != nil {
		return in.Errorf("app execution failed: %w", err1)
	}
	return nil
}

// Task 9
func ParseAge(s string) (int, error) {
	age, e := str.Atoi(s)
	if e != nil {
		return 0, in.Errorf("failed to parse age: %w", e)
	}

	if age < 0 {
		return 0, err.New("age cannot be negative")
	}

	if age > 150 {
		return 0, err.New("age is too high")
	}

	return age, nil
}

// Task 10
func SafeDivide(a, b int) (result int, e error) {
	defer func() {
		if r := recover(); r != nil {
			e = in.Errorf("division by zero caught: %v", r)
		}
	}()

	return a / b, nil
}

// Task 11
func ValidatePassword(password string) []error {
	var errs []error

	if len(password) < 8 {
		errs = append(errs, err.New("password must be at least 8 characters long"))
	}

	hasDigit := false
	hasUpper := false

	for _, char := range password {
		if unicode.IsDigit(char) {
			hasDigit = true
		}
		if unicode.IsUpper(char) {
			hasUpper = true
		}
	}

	if !hasDigit {
		errs = append(errs, err.New("password must contain at least one digit"))
	}

	if !hasUpper {
		errs = append(errs, err.New("password must contain at least one uppercase letter"))
	}

	return errs
}

// Task 12
func Retry(attempts int, f func() error) error {
	var lastErr error

	for i := 0; i < attempts; i++ {
		e := f()
		if e == nil {
			return nil
		}
		lastErr = e
	}

	return in.Errorf("all attempts exhausted: %w", lastErr)
}

// Task 13
func Register(login string, users map[string]bool) error {
	if users[login] {
		return ErrUserExists
	}
	users[login] = true
	return nil
}
func RegisterWithRetry(baseLogin string, users map[string]bool) string {
	candidate := baseLogin
	counter := 2

	for {
		e := Register(candidate, users)
		if e == nil {
			return candidate
		}

		if err.Is(e, ErrUserExists) {
			candidate = in.Sprintf("%s_%d", baseLogin, counter)
			counter++
		} else {
			return candidate
		}
	}
}

// Task 14
func ProcessBatch(items []string, process func(string) error) (successCount int, errs []error) {
	for _, item := range items {
		if e := process(item); e != nil {
			errs = append(errs, e)
		} else {
			successCount++
		}
	}
	return successCount, errs
}

// ============== MAIN ===============
func main() {
	/////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 1 ===")

	vls := []float64{12.0, 34.2, -19.0, 56.4, 3.2, 7.9, -89.0}

	for _, v := range vls {
		res, err := Sqrt(v)
		if err != nil {
			in.Printf("Error: %v", err)
		}
		in.Printf("%.1f\n", res)
	}

	//////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 2 ===")

	testCases := []struct {
		a, b float64
	}{
		{10, 2},
		{10, 0},
		{-15, 3},
		{7.5, 0},
		{0, 5},
	}

	for _, tc := range testCases {
		result, err := Divide(tc.a, tc.b)
		if err != nil {
			in.Printf("Error: %v\n", err)
		} else {
			in.Printf("Result %g / %g = %g\n", tc.a, tc.b, result)
		}
	}

	////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 3 ===")

	operations := []struct {
		balance, amount float64
	}{
		{1000, 233},
		{10, 0},
		{100, 123},
	}

	for _, op := range operations {
		res, err := Withdraw(op.balance, op.amount)

		if err != nil {
			in.Printf("\nError: %v", err)
		} else {
			in.Printf("\nWithdrawal was successfull! Balance: %1.f", res)
		}
	}

	////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 4 ===")

	values := &[]int{1, 7, 3, 5, 2, 10, 67, 34, 22, 555, 78, 1000, 888}

	for _, v := range *values {
		err := Pipeline(&v)
		if err != nil {
			in.Printf("\nAn error occurred: %v", err)
		} else {
			in.Println("\nNot error occurred")
		}

	}

	///////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 5 ===")

	ages := []int{-5, 25, 150, 0, 121}

	for _, age := range ages {
		err1 := ValidateAge(age)

		if err.Is(err1, ErrNegativeAge) {
			in.Printf("Age %d: Negative number\n", age)
		} else if err.Is(err1, ErrTooOld) {
			in.Printf("Age %d: reached limit\n", age)
		} else {
			in.Printf("Age %d: is valid\n", age)
		}
	}

	////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 6 ===")

	tests := []struct {
		login    string
		password string
	}{
		{"", "123456"},
		{"alex", "123"},
		{"admin", "qwerty"},
	}

	for _, t := range tests {
		err2 := CheckCredentials(t.login, t.password)

		switch {
		case err.Is(err2, ErrEmptyLogin):
			in.Println("Authorization error: login field is empty")
		case err.Is(err2, ErrShortPassword):
			in.Println("Authorization error: password is too short")
		case err2 == nil:
			in.Println("Login successful")
		default:
			in.Printf("Unknown error: %v\n", err2)
		}
	}

	////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 7 ===")

	err2 := loadItem(42)

	// Печатаем полный текст ошибки
	in.Println("Error:", err2)

	// 4. Проверяем, содержится ли ErrNotFound внутри цепочки ошибок
	if err.Is(err2, ErrNotFound) {
		in.Println("An error had been found, ouch")
	} else {
		in.Println("None error had occurred, phew!")
	}

	////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 8 ===")

	err1 := run()
	if err1 == nil {
		return
	}

	currErr := err1
	for currErr != nil {
		in.Println(currErr)
		currErr = err.Unwrap(currErr)
	}

	////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 9 ===")

	agesToTest := []string{"25", "abc", "-3", "999"}
	for _, raw := range agesToTest {
		age, e := ParseAge(raw)
		if e != nil {
			in.Printf("Error parsing '%s': %v\n", raw, e)
		} else {
			in.Printf("Successfully parsed '%s': %d\n", raw, age)
		}
	}

	////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 10 ===")

	divTests := [][2]int{
		{10, 2},
		{5, 0},
		{100, 4},
	}
	for _, pair := range divTests {
		res, e := SafeDivide(pair[0], pair[1])
		if e != nil {
			in.Printf("SafeDivide(%d, %d) error: %v\n", pair[0], pair[1], e)
		} else {
			in.Printf("SafeDivide(%d, %d) = %d\n", pair[0], pair[1], res)
		}
	}

	////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 11 ===")

	passwordsToTest := []string{"pass", "password123", "Password", "Valid1234"}
	for _, pwd := range passwordsToTest {
		in.Printf("Validating '%s':\n", pwd)
		errs := ValidatePassword(pwd)
		if len(errs) == 0 {
			in.Println("  Password is valid!")
		} else {
			for _, e := range errs {
				in.Printf("  - %v\n", e)
			}
		}
	}

	////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 12 ===")

	attemptsCount := 0
	flakyFunction := func() error {
		attemptsCount++
		if attemptsCount < 3 {
			return err.New("temporary failure")
		}
		return nil
	}

	e := Retry(5, flakyFunction)
	if e != nil {
		in.Printf("Retry failed: %v\n", e)
	} else {
		in.Printf("Retry succeeded on attempt %d\n", attemptsCount)
	}

	////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 13 ===")

	usersDB := map[string]bool{
		"alex":   true,
		"alex_2": true,
		"alex_3": true,
	}

	finalLogin := RegisterWithRetry("alex", usersDB)
	in.Printf("Registered login: %s\n", finalLogin)

	////////////////////////////////////////////////////////////////////////////

	in.Println("\n=== Task 14 ===")

	batchItems := []string{"hello", "", "this string is way too long for processing", "golang", "valid"}
	validator := func(s string) error {
		if len(s) == 0 {
			return err.New("empty string is not allowed")
		}
		if len(s) > 10 {
			return err.New("string length exceeds 10 characters")
		}
		return nil
	}

	successes, batchErrors := ProcessBatch(batchItems, validator)
	in.Printf("Success count: %d\n", successes)
	in.Println("Errors encountered:")
	for _, e := range batchErrors {
		in.Printf("  - %v\n", e)
	}
}
