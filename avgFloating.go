package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var fl string
	var b bool
	var value string

	for b == false {
		fmt.Println("Enter float numbers and q or Q to quit")
		fmt.Scan(&fl)
		if strings.Contains(fl, "q") || strings.Contains(fl, "Q") {
			b = true
		} else {
			fl += ","
			value += fl
		}

	}
	fmt.Println(value)
	var array = strings.Split(value, ",")
	slice := array[:len(array)]
	fmt.Println(slice)
	var a = make([]float64, len(slice))
	for i := 0; i < len(slice); i++ {
		a[i], _ = strconv.ParseFloat(slice[i], 32)
	}
	fmt.Println(a)

	av, s := avgfun(a)

	fmt.Println("Average is : ", av)
	std := stdDev(av, s, a)
	fmt.Println("Standard deviation is : ", std)
}

func avgfun(a []float64) (float64, float64) {
	var sum float64
	var avg float64
	for i := 0; i < len(a); i++ {
		sum += a[i]

	}
	size := float64(len(a))
	avg = sum / size

	return avg, size
}

func stdDev(mean float64, N float64, a []float64) float64 {
	var sum float64

	for i := 0; i < len(a); i++ {
		sum += math.Pow((a[i] - mean), 2)
	}
	std := math.Pow(sum/N, 0.5)

	return std
}
