package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the feedback section"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our pizza")
	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating, ", input)
	//fmt.Println(input)
	//fmt.Printf("thanks for rating: %v", input)
}
