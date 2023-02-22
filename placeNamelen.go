package main

import "fmt"

func main() {
	var array [7]string
	var s string
	for i := 0; i < 7; i++ {
		fmt.Println("Enter place name : ")
		fmt.Scan(&s)
		array[i] = s
	}
	fmt.Println(placeLen(array))

}

func placeLen(a [7]string) [7]int {
	var intarray [7]int
	for index, values := range a {
		intarray[index] = len(values)
	}
	return intarray
}
