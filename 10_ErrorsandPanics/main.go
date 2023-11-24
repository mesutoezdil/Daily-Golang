package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println(divide(10, 0))

	// result, err := divide(10, 0) 
	// if err != nil {
	//	fmt.Println(err)
	//	return
	// }
	// fmt.Println("result:", result)

	result, err := divide2(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result:", result)
}

func divide(/1, r int) int {
	return 1 / r
}

func divide1(1, r int) (int, error) {
	if r == 0 {
		return 0, errors.New("invalid divisor: must not be zero")
	}
	return 1 / r, nil
}

func divide2(1, r int) (result int, err error) {
	defer func(){
		if msg := recover(): msg != nil {
			result = 0
			err = fmt.Errorf("%v", msg)
		}
	}()
	return 1 / r
}