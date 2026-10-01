package main

import in "fmt"

func main() {
	if1()
	if2()
	if3()
	if4()
	if5()
	if6()
	if7()
	if8()
	if9()
	if10()
	if11()
	if12()
	if13()
	if14()
	if15()
	if16()
	if17()
	if18()
	if19()

}

func if1() {
	num1 := 4

	if num1 > 0 {
		num1 += 1
	}

	in.Printf("if_1) %v\n", num1)

}

func if2() {
	num1 := 20

	if num1 > 0 {
		num1 += 1
	} else {
		num1 -= 2
	}

	in.Printf("if_2) %v\n", num1)
}

func if3() {
	num1 := 20

	if num1 > 0 {
		num1 += 1
	} else if num1 < 0 {
		num1 -= 2
	} else {
		num1 = 10
	}

	in.Printf("if_3) %v\n", num1)
}

func if4() {
	num1 := 4
	num2 := 6
	num3 := -1
	counter := 0

	if num1 >= 0 {
		counter += 1
	}

	if num2 >= 0 {
		counter += 1
	}

	if num1 >= 0 {
		counter += 1
	}

	if num3 >= 0 {
		counter += 1
	}

	in.Printf("if_4) %v\n", counter)
}

func if5() {
	num1 := -10
	num2 := 6
	num3 := -1
	counterPos := 0
	counterNeg := 0

	if num1 >= 0 {
		counterPos += 1
	} else {
		counterNeg += 1
	}

	if num2 >= 0 {
		counterPos += 1
	} else {
		counterNeg += 1
	}

	if num1 >= 0 {
		counterPos += 1
	} else {
		counterNeg += 1
	}

	if num3 >= 0 {
		counterPos += 1
	} else {
		counterNeg += 1
	}

	in.Printf("if_5)Positive numbers: %v\nNegtaive numbers: %v\n", counterPos, counterNeg)
}

func if6() {
	num1 := 67
	num2 := 33

	if num1 > num2 {
		in.Printf("if_6) %v\n", num1)
	} else {
		in.Printf("if_6) %v\n", num2)
	}
}

func if7() {
	num1 := 453
	num2 := 666

	if num1 < num2 {
		in.Print("if_7): 1")
	} else {
		in.Print("if_7): 2")
	}
}

func if8() {
	a := 44
	b := 990

	if a > b {
		in.Println("if_8): ", a)
		in.Println("if_8): ", b)
	} else {
		in.Println("if_8): ", b)
		in.Println("if_8): ", a)
	}
}

func if9() {
	a := 9
	b := 7

	if a > b {
		a, b = b, a
	}

	in.Println("if_9): ", a)
	in.Println(b)

}

func if10() {
	a := 0
	b := 33

	if a != b {
		sum := a + b
		a = sum
		b = sum
	} else {
		a = 0
		b = 0
	}

	in.Println("if_10): ", a)
	in.Println(b)
}

func if11() {
	a := 5
	b := 5

	if a == b {
		a = 0
		b = 0
	} else if a > b {
		b = a
	} else {
		a = b
	}

	in.Println("if_11) :", a, b)
}

func if12() {
	a := 55
	b := 2
	c := -1
	min := -a

	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	in.Println("if_12): ", min)

}

func if13() {
	a := 203
	b := 77
	c := -111

	if (a >= b && a <= c) || (a >= c && a <= b) {
		in.Println("if_13) : ", a)
	} else if (b >= a && b <= c) || (b >= c && b <= a) {
		in.Println("if_13) : ", b)
	} else {
		in.Println("if_13) : ", c)
	}
}

func if14() {
	a := 2
	b := 73
	c := -11

	min := a
	max := a

	if b < min {
		min = b
	}
	if c < min {
		min = c
	}

	if b > max {
		max = b
	}
	if c > max {
		max = c
	}

	in.Println("if_14) :", min, max)

}

func if15() {
	a := 6
	b := 4
	c := -1

	if a <= b && a <= c {
		in.Println("if_15): ", b+c)
	} else if b <= a && b <= c {
		in.Println("if_15): ", a+c)
	} else {
		in.Println("if_15): ", a+b)
	}
}

func if16() {
	a := 7
	b := -51
	c := 532

	if a < b && b < c {
		a *= 2
		b *= 2
		c *= 2
	} else {
		a = -a
		b = -b
		c = -c
	}
	in.Println("if_16): ", a, b, c)

}

func if17() {
	a := -2
	b := 44
	c := 46

	if (a < b && b < c) || (a > b && b > c) {
		a *= 2
		b *= 2
		c *= 2
	} else {
		a = -a
		b = -b
		c = -c
	}
	in.Println("if_17): ", a, b, c)
}

func if18() {
	a := 3
	b := 4
	c := 6

	if a == b {
		in.Println("if_18): 3")
	} else if a == c {
		in.Println("if_18): 2")
	} else {
		in.Println("if_18): 1")
	}
}

func if19() {
	a := -3
	b := -5
	c := -1
	d := 0

	if a == b && a == c {
		in.Println(4)
	} else if a == b && a == d {
		in.Println(3)
	} else if a == c && a == d {
		in.Println(2)
	} else {
		in.Println(1)
	}
}
