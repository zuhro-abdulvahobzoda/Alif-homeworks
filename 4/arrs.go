package main

import in "fmt"

var array1 = [5]int{1, 2, 3, 4, 5}

func main() {
	arr1()
	arr2()
	arr3()
	arr4()
	arr5()
	arr6()
	arr7()
	arr8()
	arr9()
}

func arr1() {

	in.Print("Task1): ", array1, len(array1))
}

func arr2() {
	array2 := array1

	in.Print("\nTask2): ", array2[1:4], len(array2[1:4]))
}

func arr3() {
	array3 := []int{12, 23, 33, 11, 1, 9}
	sum := 0
	for _, value := range array3 {
		sum += value
	}

	in.Print("\nTask3): ", sum)

}

func arr4() {
	var array4 []string

	slice1 := append(array4, "Hello")
	in.Print("\nTask4): ", slice1)

	slice2 := append(slice1, "Bye")
	in.Print(slice2)

	slice3 := append(slice2, "What")
	in.Print(slice3)

}

func arr5() {
	array5 := []int{1, 6, 4, 14}

	in.Print("\nTask5): ", array5)
	for i := range array5 {
		array5[i] *= 2
	}
	in.Print(" ", array5)
}

func arr6() {
	slc2 := []int{1, 9, 12, 44, 0, -1}
	max := slc2[0]

	for _, v := range slc2 {
		if v > max {
			max = v
		}
	}
	in.Print("\nTask6): ", max)
}

func arr7() {
	slc3 := []int{1, 8, 19, 0, 89, -1, -6}
	in.Print("\nTask7): ")

	for i := len(slc3) - 1; i >= 0; i-- {
		in.Print(" ", slc3[i])
	}

}

func arr8() {
	slc4 := make([]int, 1)

	in.Print("\nTask8): ", len(slc4), cap(slc4))

	for i := 0; i <= 5; i++ {
		slc4 = append(slc4, i)
		in.Print(" || ", len(slc4), cap(slc4))
	}
}

func arr9() {
	slc5 := []int{1, 2, 4, 6, 7, 10, 77, 5, 22, 9, 88, 90, 55, 67}
	oddNumbers := 0
	evenNumbers := 0

	for _, v := range slc5 {
		if v%2 == 0 {
			evenNumbers += 1
		} else {
			oddNumbers += 1
		}

	}

	in.Printf("\nTask9): Even numbers (%d), Odd numbers (%d)", evenNumbers, oddNumbers)

}
