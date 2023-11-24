/*
A method in object-oriented programming is a function associated with an invocation and a variable.

Function:

var i int
func isEven(i int) bool {
	return i%2==0
}

--> Methods indicate a tighter coupling between a function and a type

type myInt int												//need a type to bind method to. 
																			//does not have to be a struct
func (i myInt) isEven() bool {				//method receiver : i myInt
	return int(i)%2==0
}

--
//function
var i int
func isEven(i int) bool {
	return i%2==0
}
ans := isEven(i)

//method
type myInt int
var mi myInt
func (i myInt) isEven() bool {
	returm int(i)%2==0
}
ans = mi.isEven()

---------------------
-> Method Receivers: Use pointer receivers to share variable between caller and method. 

type user struct {
	id				int
	username	string
}

func (u user) String() string{														//value receiver
	return fmt.Sprintf("%v (%v)\n", u.username, u.id)
}

func (u *user) UpdateName(n name) {												//pointer receiver
	u.username = name
}

--DEMO--

package menu

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

type menuItem struct {
	name string
	prices map[string]float64
}

type menu []menuItem

func (m menu) print() {
	for _, item := range m {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item,prices {
			fmt. Printf("\t%10s%10.2f\n, size, cost")
		}
	}
}

func (m *menu) add(){
	fmt.Println("Please enter the name of the new item")
		name, _ := in.ReadString('\n')
		m = append(*m, menuItem{name: strings.TrimSpace(name), prices: make(map[string]float64)})
}

var in = bufio.NewReader(os.Stdin)

func Add() {
	data.add()
}

func Print() {
	data.print()
}

----------------------------
METHODS AND CONCRETE TYPES
* os.File			 ----> Read(...) ----> [byte]
* net.TCPConn	 ----> Read(...) ----> [byte]
or
* os.File/net.TCPConn ----> io.Reader ----> Read(...) ----> [byte]

--> INTERFACE:

type Reader interface {
	Read([]byte) (int, error)
}

type File struct { . . .}
func (f File)Read(b []byte) (n int, err error)

type TCPConn struct { . . . }
func (t TCPConn) Read(b []byte) (n int, err error)

var f File
var f TCPConn

var r Reader
r = f
r.Read(...)																				// read from File
r = t 
r.Read(...)																				// read from TCPConn

--> TYPE ASSERTIONS:
type Reader interface {
	Read([]byte) (int, error)
}

type File struct { . . . }
func (f File)Read(b []byte) (n int, err error)

var f File
var r Reader = f

var f2 File = r								// error, Go cannot be sure this will work
f2 = r.(File)									// type assertion, panics upon failure
f2, ok := r.(File)						// type assertion with comma okay, does not panic

--> TYPE SWITCHES:
 var f File
 var r Reader = f

 switch v := r.(type) {
 	case File:
		//v is now a File object
	case TCPConn:
		//v is now a TCPConn object
	default:
		//this is selected if no types were matched 
 }


-------DEMO-----------------------------------------
package main

type printer interface {
	Print() string
}

type user strict {
	username string
	id int
}

func (u user) Print() string{
	return fmt.Sprintf("%v [%v]\n", u.username, u.id)
}

type menuItem struct {
	name string
	prices map[string]float64
}

func (mi menuItem) Print() string {
	var b bytes.Buffer
	b.WriteString(mi.name + "\n")
	b.WriteString(string.Repeat("-", 10) + "\n")
	for size, cost := range mi.prices {
		fmt.Fprint(&b, "\t%10s%10.2f\n", size, cost)
	}

	return b.String()
}

func main() {
	var p printer
	p = user{username: "adent", id:42}
	fmt.Println(p.Print())

	p = menuItem{name: "Coffee",
		prices: map[string]float64{"small": 1.65,
			"medium": 1.80,
			"large": 1.95,
		},
	}
	fmt.Println(p.Print())

	u, ok := p.(user)
	fmt.Println(u, ok)
	mi, ok := p.(menuItem)
	fmt.Println(mi, ok)

	switch v := p.(type) {
		case user:
			fmt.Println("Found a user!", v)
		case menuItem:
			fmt.Println("Found a menuItem!", v)
		defualt:
			fmt.Println("I am not sure what this is . . .")
	}
}

---------------------------
----GENERIC PROGRAMMING----

type Reader interface {
	Read([]byte) (int, error)
}

type File struct { . . .}
func (f File)Read(b []byte) (n int, err error)

type TCPConn struct { . . . }
func (t TCPConn) Read(b []byte) (n int, err error)

var f File
var f TCPConn

var r Reader
r = f																			//types lose identity
r = t																			//types lose identity
--->But on generic programming types do not lose their identites!

--------
DEMO for Creating Generic Function¬

package main

import "fmt"

func main(){
	testScores := []float64{
		87.3
		105
		63.5
		27
	}

	c := clone(testScores)

	fmt.Println(%testScores[0], &c[0], c)
}

func clone[V any](s []V) []V {												//V: Generic variable
	result := make([]V, len(s))
	for o, v := range s {
		result[i] = v
	}
	return result
}

--------
DEMO for Generics with the Comparable Constraint¬

package main

import "fmt"

func main(){
	testScores := map[string]float32{
		"Hary": 87.3,
		"Hermione": 105,
		"Ronald": 63.5,
		"Neville": 27,
	}

	c := clone(testScores)

	fmt.Println(c)
}

func clone[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}


--------
DEMO for Creating Custom Type Constraints

package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3}
	a2 := []float64{3,14, 6.02}
	a3 := []string{"foo", "bar", "baz"}

	s1 := add(a1)
	s2 := add(a2)
	s3 := add(a3)

	fmt.Printf("Sum of %v: %v\n", a1, s1)
	fmt.Printf("Sum of %v: %v\n", a2, s2)
	fmt.Printf("Sum of %v: %v\n", a3, s3)
}

type addable interface {
	int | float64 | string
}

func add[V addable](s []V) V{
	var result V
	for _, v := range s {
		result += v
	}
	return result
}


-----------
-------------
------------------
func foo[T any]() {...}								//create a function with a generic parameter 'T'
func bar[T any, S any]() {...}				//can use multiple generic types per function
func baz[T any](in T) T {							//generics maintain type from consumer's perspective
	return in
}
fmt.Printf("%T", baz(3))							// int
fmt.Printf("%T", baz(true))						// bool

any																		//matches any type, like interface[]
comparible														//matches types that can ber compared

type Addable interface {							//create a type interface
	int | float64	
}
func add[T Addable](){ . . . }				//used like other types as generic parameter

USEFUL PACKAGES:
--> xexp/constraints
--> x/exp/slices
--> x/exp/maps