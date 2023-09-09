package main

import "fmt"

func main() {

	arr := [3]int{101, 102, 103}
	for i, v := range arr {
		fmt.Println(i, v)

	}
	fmt.Println("done!")

}
