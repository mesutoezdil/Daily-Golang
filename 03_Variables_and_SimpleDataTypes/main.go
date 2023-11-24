// Package clause
package main

// Import statement
import "fmt"

// My Code
func main() {
	fmt.Println("Hello World")
}

//if you wan to create a new Go programme you need to create a module
//FIRST!!! create a module: go mod init demo

/*
Simple Data Types: Strings, Numbers, Booleans, Errors 

"this is a string" --> interpreted string

`this is also a string` --> raw string (if you have it, be careful spaces)

"this is an escape charecter: \n it creates a newline"
			output: this is an escape character:
								it creates a newline

`this is an escape charecter: \n it creates a newline`
			output: this is an escape character: \it creates a newline


NUMBERS
- Integers (int): 99, 0, -937 
- Unsigned integers (uint): 0, 15, 7233
- Floating point numbers (float32, float64): 6.02e23, 3.1415, 0.25
- Complex numbers (complex64, complex64): 1+2i, 0.833i, 6.02e23+3.1415i

error type: the error built-in interface type is the conventional interface for representing an error condition, with the nil value representing no error 

VARIABLES	
var myName string 						//declare variable
var myName string = "Mike"		//declare and initialize
var myName = "Mike"						//initialize with inferred type
myName := "Mike"							//short declaration syntax

One of Go's basic principles, one of its founding guidelines is to be clear over being clever.

TYPE CONVERSIONS:
var i int = 32
var f float32

f = i 								//error! - Go does not support implicit conversions
f = float32(i)				//type conversion allow explicit conversion

--
package main
import "fmt"

func main() {
	var a string
	a = "foo"
	fmt.Println(a)

	var b init = 99
	fmt.Println(b)

	var c = true
	fmt.Println(c)

	d := 3.1415 //float64
	fmt.Println(d)

	var e int = d
	fmt.Println(e)

}

ARITHMETIC
a, b := 10, 5		//Go allows multiple variables to be initialized at once!
c := a + b			//15 - addition
c = a - b				//5 - subtraction
c = a * b				//50 - multiplacation
c = a / b				//2 - division
c = a / 3				//3 - integer division used for integers
c = a % 3				//1 - modulus (remainder of integer division)
d := 7.0 / 2.0	//3.5 - decimal results given for floating point numbers

COMPARISON
a, b := 10, 5
c := a == b 		// false - equality
c = a != b 			// true - inequality
c = a < b				// false - less than
c = a <= b 			// false - less than or equal
c = a > b 			// true - greater
c = a >= b			// true - greater than or equal

--
package main
func main(){
	a, b := 5, 2
	fmt.Println(a + b)
	fmt.Println(a % b) //a%2 == 0
	fmt.Println(float(a) / float(b)) //2.5
}

CONSTANTS
const a = 42 //constant (implicitly typed)
const b string = "hello, world" //explicitly typed constant
const c = a //one constant can be asigned to another

const (
	d = true
	e = 3.14
)

const (
	a = "foo"
	b									//"foo" --> unassigned constants receive previous value
)

const c = 2 * 5									// 10 --> constant expression
const d = "hello, "+"world"			// must be calculable at compile time
const e = someFunction()				
// this won't work - cannot be evaluated at compile time, bcs of the memory allocation




IOTA
-----
const a = iota		// 0 --> iota is related to position in constant group

const (
	b = iota			// 0 --> iota starts at zero on first line
	c = 					// 1 --> constant expression copied, iota increments
	d = 3 * iota	// 6 --> iota increments again
)

const (
	e = iota			// iota resets to zero with each group
)


package main
imp "fmt"
func main(){
	const a = 42
	var i int = a 
	var f float64 = a

	fmt.Println(i, f)				// output: 42 42
}

package main
imp "fmt"
func main(){
	const a = 42
	const b float32 = 3
	var f32 float32 = b 
	var f64 float64 = float64(b)
	fmt.Println(f32, f64)							// output: 3 3
}

package main
imp "fmt"
func main(){
	const a = 42
	const b float32 = 3
	var f32 float32 = b 
	var f64 float64 = float64(b)
	fmt.Println(f32, f64)							
	const c = iota										// output: 3 3 0
	fmt.Println(c)
	const (
		d = 2*5
		e // = 2*5
		f = iota // not 0! but 2, it is always related to the position of the group
		g
		h = 10 * iota
	)
	fmt.Println(d, e, f, g, h)									// output: 3 3 0 10 10 2 3 40		

}

POINTERS AND VALUES
---------
--> Pointers are primarely used to share memory.
--> Use copies whenever possible.
a := "foo"		//create a string variable
b := &a				//address operator returns the address of a variable
*b = "bar"		//dereference a pointer with asterisk
c = new(int)	//built-in "new" function creates pointer to anonymous variable

package main
imp "fmt"
func main(){
	s := "Hello, world!"	//s is a variable
	p := &s								//p is a pointer of s	
	fmt.Println(p)				//output: the memory address, this actual address locates in memory
	fmt.Println(*p)				//output: Hello, world!.. p is pointed s, through p to the s graped the value

	*p = "Hello, gophers!"
	fmt.Print(s)					//Hello, gophers! -defererancing of operations

	p = new(string)
	fmt.Println(p)				//output:the memory address of p, different than first one!
}