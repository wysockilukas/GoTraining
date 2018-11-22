package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	//wysylamy wartosc do kanalu
	go wyslij(c)

	//odbieramy z kanalu
	for v := range c {
		fmt.Println(" ", v)
	}

	fmt.Println("zaraz koniec")
}

func wyslij(do chan<- int) {
	for i := 1; i <= 100; i++ {
		do <- i
	}
	close(do)
}
