/*
IF STATEMENT

if test { . . . } 

if test { . . . } 
else { . . . } 

if test { . . . } 
else if test { . . . } 

if test { . . . } 
else if test { . . . } 
else { . . . } 

if initializer; test { . . . } 

--

package main
import "fmt"
func main(){
	i:= 15
	if i < 5 {																			//alternative--> if i := 15; i < 5 {...}
		fmt.Println("i is less than 5")
	} else if i < 10 {
		fmt.Println("i is less than 10")
	} else {
		fmt.Println("i is at least 10")
	}
	fmt.Println("after the if statement")
}

-------------
SWITCH STATEMENTS

switch test expression {
	case expression1:
		statements
	case expression2, expression3:
		statements
	defaults:
		statements
}

--
switch i:=999; i {										//initializer
	case 1:
		fmt.Println("first case")
	case 2 + 3, 2*i*3;
		fmt.Println("second case")
	default:
		fmt.Println("default case")
}

--
LOGICAL SWITCH

switch i := 8; {                       //true is implied
case i < 5:
	fmt.Println("i is less than 5")
case i < 10:
	fmt.Println("i is less than 10")
default:
	fmt.Println("i is greater than 10")
}

--
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type menuItem struct {
		name string
		prices map[string]float64
	}

	menu := []menuItem{
		{name :"Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80}},
		{name :"Espresso", prices: map[string]float64{"single": 1.90, "double": 2.25}},
	}

	in := bufio.NewReader(os.Stdin)

loop:
	for{
		fmt.Prinln("Please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) quit")
		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			for _, item := range menu {
				fmt.Println(item.name)
				fmt.Println(strings.Repeat("-", 10))
				for size, price := range item,prices {
					fmt. Printf("\t%10s%10.2f\n, size, price")
				}
			}
		case "2":
			fmt.Println("Please enter the name of the new item")
			name, _ := in.ReasdString('\n')
			menu = append(menu, menuItem{name: name, prices:make(map[string]float64)})
		case "q":
			break loop
		default: 
			fmt.Println("Unknown option")
		}
	}
}
--
DEFERRED FUNCTIONS
invocation --> Execute statements --> Exit --> Return focus to caller

func main(){
	fmt.Println("main 1")
	defer fmt.Println("defer 1")
	fmt.Println("main 2")
	defer fmt.Println("defer 2")
}

output: main 1 --> main 2 --> defer 2 --> defer 1

--
package main
import "database/sql"
func main(){
	db, _ := sql.Open("drivenName", "connection string")
	defer db.Close()

	rows, _ := db.Query("some sql query here")
	defer rows.Close()
}

----------------------------
PANIC and RECOVER

func main() {
	fmt.Println("main 1")
	func1()
	fmt.Println("main 2")
}

func func1(){
	defer func() {
		fmt.Println(recover())
	}()
	fmt.Println("func1 1")
	panic("uh-oh")
	fmt.Println("func1 2")
}

---> main -> func1 1 -> uh-oh -> main 2


--
main package
import "fmt"

func main() {
	dividend, divisor := 10, 5
	fmt.Printf("%v divided by %v is %v\n", dividend, divisor, divide(dividend, divisor))

	dividend, divisor := 10, 0
	fmt.Printf("%v divided by %v is %v\n", dividend, divisor, divide(dividend, divisor))
}

func divide(dividend, divisor int) in {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}
	}()
	return dividen / divisor
}
----
GOTO STATEMENT¬

func myFunc(){
	i := 10
	if i < 15 {
		goto myLabel					//can leave a block
	}
myLabel:									//can jump to containing block
	j := 42									//cannot jump after variable declaration
	for ; i < 15; i++ {     
		...										//cannot jump into another block
	}
}
