package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to array")
	var fruitList = [4]string{}

	fruitList[0] = "Mango"
	fruitList[1] = "Apple"
	fruitList[3] = "Jamun"

	fmt.Println("the fruit list is: ", fruitList)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter first fruit")
	input, _ := reader.ReadString('\n')

	fmt.Println("thnkyou for entering first fruit which is, ", input)

	fmt.Println("Enter second fruit")
	input2, _ := reader.ReadString('\n')

	fmt.Println("your 2nd fruit is, ", input2)

	input3 := input + input2
	fmt.Println("addition of two fruit is, ", input3)

	// fmt.Println("Enter first number")
	// input, _ := reader.ReadString('\n')

	// fmt.Println("thnkyou for entering first number, ", input)

	// fmt.Println("Enter second number")
	// input2, _ := reader.ReadString('\n')

	// fmt.Println("your 2nd number is, ", input2)

	// input3 := input + input2
	// fmt.Println("addition of two number is, ", input2)
	// fmt.Printf("addition of two number is %v", input3)
}
