package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("The one took the shortest time wins the race")
	dragRace := []int{}

	dragRace = append(dragRace, 100, 80, 79, 120, 96, 75, 150)
	fmt.Println("Time took to complete race by all the drivers are, ", dragRace)

	sort.Ints(dragRace)
	fmt.Println(dragRace)

}
