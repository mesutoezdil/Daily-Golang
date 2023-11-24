LOOPING

/* looping execution

func main(){
	statement 1
	loop {
		statement 2
		statement 3
	}
	...
	statement n
}

------
for { . . . }																		//infinite loop
for condition { . . . }													//loop till condition 
for initializer; test; post clause { . . . }		//counter-based loop
------
i := 1
for {															//infinite loop
	fmt.Println(i)
	i += 1
}
------
i := 1
for i < 3 {												//loop till condition
	fmt.Println(i)
	i += 1
}
fmt.Println("Done!")
------
for i:=1; i < 3; i++ {						//counter-based loops
	fmt.Println(i)
}
fmt.Println("Done!")
-------
package main
func main(){
	i := 99

	for {
		fmt.Println(i)
		i + 1
		break
	}

	i = 5
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	} 
}
------
Looping with Collections

for key, value := range collection { . . . }
for key := range collection { . . . }
for _, value := range collection { . . . }
--
arr := [3]int{101, 102, 103}
for i, v := range arr {
	fmt.Println(i, v)
}
fmt.Println("Done!")
--
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Please select on option")
	fmt.Println("1) Print menu")
	in := bufio.NewReader(os.Stdin)
	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(choice) //we do not know what to do with this yet!

	type menuItem struct {
		name string
		prices map[string]float64
	}

	menu := []menuItem{
		{name :"Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80}},
		{name :"Espresso", prices: map[string]float64{"single": 1.90, "double": 2.25}},
	}

	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item,prices {
			fmt. Printf("\t%10s%10.2f\n, size, price")
		}
	}
}
--