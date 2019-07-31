package main

import (
	"fmt"
	"sort"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func largest(a []int, c chan int) {
	sort.Ints(a)
	fmt.Println(a[len(a)-1])
	c <- a[len(a)-1]
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	go largest(a, c)
	x, y, z := <-c, <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
	fmt.Println("this is largest number", z)

}
