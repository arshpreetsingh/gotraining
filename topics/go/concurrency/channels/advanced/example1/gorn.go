package main

import (
	"fmt"
)

func SendValue(c chan int) {
	c <- 8 + 4
	fmt.Println("this is hello world from inside golang as well when work is done")
}

func main() {
	fmt.Println("hello this is hello world thing!!")

	values := make(chan int)
	defer close(values)
	for i := 0; i > 1; i++ {
		fmt.Println("hello world")
		go SendValue(values)

		value := <-values
		fmt.Println(value)
	}
}
