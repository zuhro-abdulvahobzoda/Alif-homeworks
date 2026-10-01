package main

import in "fmt"

func main() {
	task1()
	task2()
	task3()
	task4()
	task4()
	task5()
	task6()
	task7()
	task8()
}

func task1() {
	n := 33

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%7 == 0 {
			in.Println(i)
			return
		}
	}

	in.Println("Not found")
}

func task2() {
	n := 20

	sum := 0

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			continue
		}

		if i%7 == 0 {
			break
		}

		sum += i
	}

	in.Println(sum)
}

func task3() {
	secret := 37
	var number int
	in.Println("Guess a number: ")

	for number != secret {
		in.Scan(&number)

		if number < secret {
			in.Println("Guessed number is greater")
		} else if number > secret {
			in.Println("Guessed number is lesser")
		}
	}

	in.Println("Correct!")
}

func task4() {
	password := 1234
	var input int

	for i := 1; i <= 3; i++ {
		in.Println("Type the password:")
		in.Scan(&input)

		if input == password {
			in.Println("Доступ разрешён")
			return
		}
	}

	in.Println("Доступ заблокирован")
}

func task5() {
	n := 23

	for i := 1; i <= n; i++ {
		if i%3 == 0 || i%10 == 5 {
			continue
		}

		in.Print(i, " ")
	}
}

func task6() {
	n := 334571
	sum := 0

	for n > 0 {
		digit := n % 10
		sum += digit
		n /= 10
	}

	in.Println(sum)
}

func task7() {
	n := 38923401

	max := 0

	for n > 0 {
		digit := n % 10

		if digit > max {
			max = digit
		}

		n /= 10
	}

	in.Println(max)
}

func task8() {
	var n int
	in.Println("Type a number:")
	in.Scan(&n)

	reversed := 0

	for n > 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n /= 10
	}

	in.Println(reversed)
}
