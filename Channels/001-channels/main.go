package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	//c <- 42 tak nie dziala
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)

	d := make(chan int, 1)
	d <- 7
	fmt.Println(<-d)

	e := make(chan int, 2)
	e <- 3
	e <- 5
	fmt.Println(<-e)
	fmt.Println(<-e)

}
