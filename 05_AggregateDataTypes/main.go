/*
ARRAYS
An array is a fixed, sized collection of data elements that all have the same type. An array is viewed by Go as a single entity. 
Within the array we can define a certain number of positions, and all of these positions are going to contain the same type of data.
This is a very convenient data structure for us to work with from the standpoint of it we've got a related set of numbers.
Each pne of the numbers is unique, starts at 0, and indicates the position of a data element within the array, and this is called the index in the array.

var arr [3]int								//array of 3 ints
fmt.Println(arr)							//[0 0 0]
arr = [3]int{1 2 3}						//array literal
fmt.Println(arr[1])						//2
arr[1] = 99										//update value
fmt.Println(arr)							//[1 99 3]
fmt.Println(len(arr))					//3

arr := [3]string{"foo", "bar", "baz"}
arr2 := arr																//{"foo", "bar", "baz"}
fmt.Println(arr2)													//arrays are copied by value
arr[0] = "quux"														
fmt.Println(arr)													//{"quux", "bar", "baz"}
fmt.Println(arr2)													//{"foo", "bar", "baz"}
arr == arr2																//false - arrays are comparable


package main
import "fmt"
func main() {
	var arr [3]string
	fmt.Println(arr)
	arr = [3]string{"Coffee", "Espresso", "Cappucino"}
	fmt.Println(arr)
	fmt.Println(arr[1])
	arr[1] = "Chai Tea"
	fmt.Println(arr)
	arr2 := arr
	arr2[2] = "Chai Latte"
	fmt.Println(arr, arr2)
}

--------
SLICES

var s []int										//slices of ints
fmt.Println(s)								//[] (nil)
s = []int{1, 2, 3}						//slice literal

fmt.Println(s[1])							//2
s[1] = 99											//update value
fmt.Println										//[1 99 3]

s = append(s, 5, 10, 15)			//add elements to the slice
fmt.Println(s)								//[1 99 3 5 10 15]

s = slices.Delete(s, 1, 3)		//remove indices 1, 2 from slice (golang.org/x/exp/slices)
fmt.Println(s)								//[1 5 10 15]

-

package main
import "fmt"
func main() {
	var s []string
	fmt.Println(s)
	s = []string{"Coffee", "Espresso", "Cappucino"}
	fmt.Println(s)
	fmt.Println(s[1])
	s[1] = "Chai Tea"
	fmt.Println(s)
	s = append(s, "Hot Chocolate", "Hot Tea")
	fmt.Println(s)
	slices.Delete(s, 1, 2)
	fmt.Println(s)
	//s2 := s
	//s2[2] = "Chai Latte"
	//fmt.Println(s, s2)
}
if you want to add dependency type on terminal: go get golang.org/x/exp/slices

--------
MAP

var m map[string]int									//declare a map
fmt.Println(m)												//map[] (nil)
m = map[string]int{"foo": 1, "bar":2}	//map literal
fmt.Println(m)												//map[foo:1 bar:2]

fmt.Println(m["foo"])									//lookup value in map
m["bar"] = 99													//update value in map

delete(m, "foo")											//remove entry from map
m["baz"] = 418												//add value to map

fmt.Println(m)												//map[bar:99 baz: 418]

fmt.Println(m["foo"])									//0 - queries always return results
v, ok := m["foo"]											//comma okay syntax verifies presents
fmt.Println(v, ok)										//0, false
--

m := map[string]int{
	"foo":1,
	"bar":2,
	"baz":3}
m2 := m																//maps are copied by referance
																			//use maps.Clone to clone
m["foo"], m2["bar"] = 99, 42					//update values in maps
fmt.Println(m)												//map[foo:99 bar:42 baz:3]
fmt.Println(m2)												//map[foo:99 bar:42 baz:3]
																			//data is shared
m == m2																//compile time error - maps are not comparable
--

package main 

import "fmt"

func main() {
	var m map[string][]string
	fmt.Println(m)

	m = map[string][]string{
		"coffee": {"Coffee", "Espresso", "Cappuccino"},
		"tea": {"Hot Tea", "Chai Tea", "Chai Latte"},
	}
	fmt.Println(m)

	fmt.Println(m["coffee"])

	m["other"] = []string{"Hot Chocolate"}
	fmt.Println(m)

	delete(m, "tea")
	fmt.Println(m)

	fmt.Prinln(m["tea"])
	v, ok := m["tea"]
	fmt.Println(v, ok)

	m2 := m
	m2["coffee"] = []string{"Coffee"}
	m["tea"] = []string{"Hot Tea"}

	fmt.Println(m)
	fmt.Println(m2)
}

-------------
STRUCTS

var s struct{					//declare an anonymous struct
	name	string
	id		int
}
fmt.Println(s)				//{"" 0}

s.name = "Arthur"			//assign value to field
fmt.Println(s.name)		//query value of field

--

type myStruct struct {			//create custom type based on struct
	name 	string
	id		int
}

var s myStruct							//declare variable with custom type
fmt.Println(s)							//{"" 0}

s = myStruct{								//struct literal
	name: "Arthur",
	id: 42}
fmt.Println(s)							//{"Arthur" 42}

--

type myStruct struct {			
	name 	string
	id		int
}
var s myStruct													
s = myStruct{								
	name: "Arthur",
	id: 42}
s2 := s
s.name = "Tricia"						//structs are copied by value
fmt.Println(s, s2)					//{"Tricia", 42} {"Arthur" 42}

s == s2											//false - structs are comparable

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

	fmt.Println(menu)
}
--