package main

import "fmt"

func main() {
	var slice [3]int
	var a int
	for i := 0; i < len(slice); i++ {
		fmt.Println("Enter a integer : ")
		fmt.Scan(&a)
		slice[i] = a
	}
	s := 0
	for s < len(slice) {
		for d := 0; d < len(slice)-1; d++ {
			if slice[d] > slice[d+1] {
				j := slice[d+1]
				slice[d+1] = slice[d]
				slice[d] = j
			}
		}
		s++
	}
	fmt.Println(slice)
}
