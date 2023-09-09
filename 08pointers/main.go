package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointer class")

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("the actual value of pointer is: ", ptr)
	fmt.Println("the actual value of pointer is: ", *ptr)

	*ptr = *ptr + 2

	fmt.Println("the new value is: ", myNumber)

}
