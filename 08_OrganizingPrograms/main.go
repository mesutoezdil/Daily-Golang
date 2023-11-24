/*
FUNCTION SIGNATURES

func functionName (parameters)(return values){
	function body
}

-----
func greet(name1 sring, name2 string) {						//func greet(name1, name2 string)
	fmt.Println(name1)
	fmt.Println(name2)
}

-----VARIADIC PARAMETERS--											//received as slice and must be final parameter
func greet(names ...string){
	for _, n := range names {
		fmt.Println(n)
	}
}


-----PASSING VALUES AND POINTERS--	
func main(){
	name, otherName := "Name", "Other name"
	fmt.Println(name)
	fmt.Println(otherName)
	myFunc(name, &otherName)
	fmt.Println(name)
	fmt.Println(otherName)
}

func myFunc(name string, otherName *string){
	name = "New name"
	*otherName = "Other new name"
}
--->name ->Other name -> Name -> Other new name
Hint: Use pointers to share memory, otherwise use values.

-----RETURN VALUES--
func functionName (parameters)(return values){
	function body
}	

--> NR1- Returning Single Value:
func main() {
	result := add(1, 2)
	fmt.Println(result)
}

func add(l, r int) int {
	return l + r
}

--> NR2- Returning Multiple Values
func main() {
	result, ok := divide(1, 2)
	if ok {
		fmt.Println(result)
	}
}

func divide(l, r int) (int, bool) {
	if r == 0 {
		return 0, false
	}
	return l/r, true
}

--> NR3- Named Return Values
func main() {
	result, ok := divide(1, 2)
	if ok {
		fmt.Println(result)
	}
}

func divide(l, r int) (result int, ok bool) {         //rarely used
	if r == 0 {
		return																						//0, false
	}
	return l/r
	ok = true
	return
	//optional: return l/r, true
}

-----DEMO-
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type menuItem struct {
		name string
		prices map[string]float64
}

	menu := []menuItem{
		{name :"Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80}},
		{name :"Espresso", prices: map[string]float64{"single": 1.90, "double": 2.25}},
}

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
			printMenu()
		case "2":
			addItem()			
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
--......------------------------
PACKAGES
-> Directory within a module
-> Contains at least one source file
-> All members are visible to other package members

package user													//package declaration
import "strings"											//import statement
var currentUsers []*User							//variable as a member of the package
const MaxUsers = 100									//constant as a member of the package
func GetByID(id int) (User, bool) {}	//function as a member of the package
--
package user													//package identifier
type User struct {										//public struct
	ID					int											//public field
	Username		string									//public field
	password		string									//package-level field
}
--DEMO--

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"demo/menu"
)

type menuItem struct {
		name string
		prices map[string]float64
}

	menu := []menuItem{
		{name :"Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80}},
		{name :"Espresso", prices: map[string]float64{"single": 1.90, "double": 2.25}},
}

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
			printMenu()
		case "2":
			menu.Add()			
		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}
	}
}
