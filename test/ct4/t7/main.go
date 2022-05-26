package main

import (
	"fmt"
	"strconv"
)

// Panic panics with a divide by zero
func panicTest() {
	zero, err := strconv.ParseInt("0", 10, 64)
	if err != nil {
		panic(err)
	}

	a := 1 / zero
	fmt.Println("we'll never get here", a)
}

// Catcher calls Panic
func catcher() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic occurred:", r)
		}
	}()
	panicTest()
}

func main() {
	fmt.Println("before panic")
	catcher()
	fmt.Println("after panic")
}
