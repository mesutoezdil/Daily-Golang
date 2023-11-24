A goroutine is a function executing concurrently with other goroutines in the same address space. It is lightweight, costing little more than the allocation of stack space.

WaitGroups are simply counters that have special behaviour when their value is zero.

type WaitGroup:
func (wg WaitGroup) Add(delta int)	//increment counter by delta
func (wg WaitGroup) Done()					//decrement counter by 1
func (wg WaitGroup) Wait()					//wait till counter is zero

---

func main(){
	go func(){
		fmt.Println("do some asyn thing")
	}()
}

---

func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		fmt.Println("do some async thing")
		wg.Done()
	}()
	wg.Wait()
}

------

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("This happens asynchronously")
		wg.Done()
	}()

	fmt.Println("This is snychronous")
	wg.Wait()
}

↓↓↓↓
----- CHANNELS ----

//create a channel
ch := make (chan string)

//send a message
ch <- "hello!"

//receive a message
msg := <- ch

Channel operations block until the complementary operation is ready

↓↓↓↓↓↓↓↓
func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)

	go func() {
			ch <- 42
	}()

	go func(){
		fmt.Println(<-ch)
		wg.Done()
	}()

	wg.Wait
}


package main
func main()  {
	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(1)

	go func() {
		ch <- "the message"
	}()

	go func(){
		fmt.Println(<-ch)
		wg.Done()
	}()

	wg.Done()
}


SELECT STATEMENT ↓↓↓↓↓↓↓↓
In a select statement, if more than one case can be acted upon then one case is chosen randomly.

select {
	case channel operation:
		statements
	case channel operation:
		statements
	default:													//optional
		statements
}

↓↓↓↓↓

--DEMO--


package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		ch1 <- "message to channel 1"
	}()

	go func() {
		ch2 <- "message to channel 2"
	}()

	time.Sleep(10 * time.Millisecond)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	default:
		fmt.Println("no messages available")
	}
}

----LOOPING---↓↓↓↓
ch := make(chan list)

go func() {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)													//no more messages can be sent
}()

for  msg := range ch {
	statements
}

↓↓↓↓↓

--DEMO--

package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
