package main

import "fmt"

func main() {
	var values = [5]int{10, 12, 15, 19, 24}
	var array [5][2]int
	var a int = 0
	for a < 5 {
		for i := 0; i < 2; i++ {
			if i == 0 {
				array[a][i] = values[a]
			} else if i == 1 {
				array[a][i] = values[a] * 2
			}
		}
		a++
	}
	fmt.Println(array)
}
