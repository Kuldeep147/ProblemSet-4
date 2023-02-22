package main

import "fmt"

func main() {
	var array [20]int
	for i := 0; i < 20; i++ {
		array[i] = i + 7
	}
	var slice []int = array[5:15]
	fmt.Println(slice)

	var odd int
	var even int

	for d := 0; d < len(slice); d++ {
		if isOdd(slice[d]) {
			odd++
		} else if isEven(slice[d]) {
			even++
		}
	}
	var evenSlice = make([]int, even)
	var oddSlice = make([]int, odd)
	var o int
	var e int
	for d := 0; d < len(slice); d++ {
		if isOdd(slice[d]) {
			oddSlice[o] = slice[d]
			o++
		} else if isEven(slice[d]) {
			evenSlice[e] = slice[d]
			e++
		}
	}
	fmt.Println(evenSlice)
	fmt.Println(oddSlice)

}

func isOdd(a int) bool {
	if a%2 == 1 {
		return true
	}
	return false
}

func isEven(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}
