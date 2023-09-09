package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza company")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("please rate our pizza")

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("type of the rating: %T\n", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("err")
	} else {
		//fmt.Printf("thnkas for new rating %v", numRating+1)
		fmt.Println("added 1 to your rating, ", numRating+1)
		fmt.Printf("type of the rating: %T\n", numRating)
	}
}
