package main

import (
	"fmt"
	"math"
)

func main() {
	var slice = make([]uint8, 5)
	slice[0] = 2
	slice[1] = 4
	slice[2] = 8
	slice[3] = 10
	slice[4] = 16
	var powerOfTwo = make([]float64, 5)

	for i := 0; i < 5; i++ {
		powerOfTwo[i] = powtwo(int(slice[i]))
	}
	fmt.Println(powerOfTwo)

}

func powtwo(a int) float64 {
	var k float64
	k = math.Pow(2, float64(a))
	return k
}
