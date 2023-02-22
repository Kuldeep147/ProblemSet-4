package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter the comma seperated values : ")
	var s string
	fmt.Scan(&s)
	var array []string
	array = strings.Split(s, ",")
	var a = make([]int, len(array))
	for i := 0; i < len(array); i++ {
		b, _ := strconv.ParseInt(array[i], 10, 64)
		a[i] = int(b)
	}

	fmt.Println("Max And max next is : ")
	c, b := maxNext(a)
	fmt.Println(c, b)
	fmt.Println("Min and min next is : ")
	g, h := minNext(a)
	fmt.Println(g, h)
}

func maxNext(a []int) (int, int) {
	var (
		max     int
		maxNext int
	)

	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}

	for i := 0; i < len(a); i++ {
		if a[i] > maxNext && a[i] != max {
			maxNext = a[i]
		}
	}
	return max, maxNext
}

func minNext(a []int) (int, int) {
	var (
		min     int = 32767
		minNext int = 32767
	)

	for i := 0; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}
	}

	for i := 0; i < len(a); i++ {
		if a[i] < minNext && a[i] != min {
			minNext = a[i]
		}
	}
	return min, minNext
}
