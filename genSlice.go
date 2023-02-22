package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter the comma seperated values : ")
	var s string
	fmt.Scan(&s)
	var array []string
	array = strings.Split(s, ",")
	fmt.Println(array)
}
