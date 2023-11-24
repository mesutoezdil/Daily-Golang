/*
CLI Tools in Standard Library
OS  --> Stdin, Stdout, Stderr
fmt --> Scan*, Print*, (http://pkg.go.dev/bufio)


package main 

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Println("What would you like me to scream?")
	in := bufio.NewReader(os.Stdin)										//we need to store that NewReader, create var in
	s, _ := in.ReadString('\n')				//ReadString comes from bufio.NewReader..s stores the data
	s = strings.TrimSpace(s)  				//TrimSpace manipulates strings
	s = strings.ToUpper(s)
	fmt.Println(s + "!")
}

We've gone from our program to the monitor, using line xx, then line xx allows us to read input from the keyboard and send that to our program, and the last thing we need to do in order to really close this out is print the result back out. So our program needs to interface with the monitor once again.

Go to the terminal: go run . ---> What would you like me to scream?
Type: Go is really cool! ---> GO IS REALLY COOL!

----BUILDUNG A WEB SERVICE------------------
package main

import (
	"io"
	"net/http"
	"os"
)

func main(){
	http:HandleFunc("/", Handler)
	http.ListenAndServe(":3008", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt")
	io.Copy(w, f)
}


------DEBUGGING A PROGRAM---------------
package main

import (
	"io"
	"net/http"
	"os"
)

func main(){
	http:HandleFunc("/", Handler)
	http.ListenAndServe(":3008", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt")
	io.Copy(w, f)
}
