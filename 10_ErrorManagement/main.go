package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


var in = bufio.NewReader(os.Stdin)

func main() {

loop:
	for {
		fmt.Println("Please select on option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item,")
		fmt.Println("q) quit")
		choice, _ := in.ReadString('\n')
		
		switch strings.TrimSpace(choice) {
		case "1":
			menu.Print()
		case "2":
			err := menu.Add()
			if err != nil { // true, if an error occurs
				fmt.Println(fmt.Errorf("invalid input: %w", err))
			}			
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}
}

func addItem() {
	fmt.Println("Please enter the name of the new item")
		name, _ := in.ReadString('\n')
		menu = append(menu, menuItem{name: strings.TrimSpace(name), prices: make(map[string]float64)})
}

func printMenu() {
	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item,prices {
			fmt. Printf("\t%10s%10.2f\n, size, cost")
		}
	}
}