package main

import in "fmt"

func main() {
	for1()
	for2()
	for3()
	for4()
	for5()
	for6()
	for7()
	for8()
	for9()
	for10()
	for11()
	for12()
	for14()
	for15()
}

func for1() {
	k := 66
	n := 32
	res := 1

	for i := 1; i <= n; i++ {
		res *= k
	}

	in.Println("for1): ", res)
}

func for2() {
	a := 27
	b := 17
	in.Print("for2): ")
	for i := a; i >= b; i-- {
		in.Print(i, " ")
	}
}

func for3() {
	a := 27
	b := 17

	in.Print("\nfor3): ")
	for i := a - 1; i > b; i-- {
		in.Print(i, " ")
	}
}

func for4() {
	price := 6
	sum := 1

	in.Print("\nfor4): ")

	for i := 1; i <= 10; i++ {
		sum = price * i
		in.Print(sum, " ")
	}
}

func for5() {
	price := 6.0
	sum := 1.0

	in.Print("\nfor5): ")

	for i := 0.0; i <= 1.0; i += 0.1 {
		sum = price * i
		in.Printf("%.1f  ", sum)
	}
}

func for6() {
	price := 6.0
	sum := 1.0

	in.Print("\nfor6): ")

	for i := 0.0; i <= 1.0; i += 0.2 {
		sum = price * i
		in.Printf("%.1f  ", sum)
	}
}

func for7() {
	a := 2
	b := 16
	sum := 0

	in.Print("\nfor7): ")

	for i := a; i <= b; i++ {
		sum += i

	}
	in.Print(sum)
}

func for8() {
	a := 8
	b := 16
	total := 1

	in.Print("\nfor8): ")

	for i := a; i <= b; i++ {
		total *= i

	}
	in.Print(total)
}

func for9() {
	a := 1
	b := 10
	sum := 0

	in.Print("\nfor9): ")

	for i := a; i <= b; i++ {
		sum += i * i

	}
	in.Print(sum)

}

func for10() {
	N := 15
	sum := 0.0

	in.Print("\nfor10): ")
	for i := 1; i <= N; i++ {
		sum += 1.0 / float64(i)
	}
	in.Print(sum)
}

func for11() {
	N := 10
	sum := 0

	in.Print("\nfor11): ")
	for i := N; i <= 2*N; i++ {
		sum += i * i
	}
	in.Print(sum)
}

func for12() {
	N := 10
	product := 1.0

	in.Print("\nfor12): ")

	for i := 1; i <= N; i++ {
		product *= 1.0 + float64(i)/10.0
	}
	in.Print(product)
}

func for14() {
	sum := 0
	N := 17

	for i := 1; i <= N; i++ {
		sum += 2*i - 1
	}
}

func for15() {
	a := 4
	N := 10
	product := 1
	in.Print("\nfor15): ")

	for i := 1; i <= N; i++ {
		product *= a

	}
	in.Print(product)
}
