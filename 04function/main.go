package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
	name   string
	prices map[string]float64
}

var menu = []menuItem{
	{name: "Coffee", prices: map[string]float64{"small": 1.65, "Medium": 1.80, "Large": 1.95}},
	{name: "Espresso", prices: map[string]float64{"single": 1.90, "double": 2.25, "triple": 2.55}},
}
var in = bufio.NewReader(os.Stdin)

func main() {
loop:
	for {
		fmt.Println("please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("3) Remove item")
		fmt.Println("4) quit")
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {

		case "1":
			printMenu()
		case "2":
			addItem()
		case "3":
			delItem()
		case "4":
			break loop
		default:
			fmt.Println("Invalid input!")
		}
	}
}

func printMenu() {
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, cost := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, cost)
		}
	}
}

func addItem() {
	fmt.Println("Please enter the name of the item")
	name, _ := in.ReadString('\n')
	menu = append(menu, menuItem{name: strings.TrimSpace(name), prices: make(map[string]float64)})
}

func delete_at_index(abc []menuItem, index int) []menuItem {
	return append(abc[:index], abc[index+1:]...)
}
func delItem() {
	for _, item := range menu {
		fmt.Println(item.name)
	}
	fmt.Println("Please enter the item to be deleted")
	name, _ := in.ReadString('\n')
	n := strings.TrimSpace(name)
	idx := findIndex(n)
	menu = delete_at_index(menu, idx)
	fmt.Println(menu)
}

func findIndex(name string) int {
	for i, v := range menu {
		loopvar := v.name
		if name == loopvar {
			return i
		}
	}
	return -1
}
