package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	//wysylamy wartosc do kanalu
	go wyslij(c)

	//odbieramy z kanalu
	odbierz(c)

	fmt.Println("zaraz koniec")
}

func wyslij(do chan<- int) {
	do <- 42
}

func odbierz(z <-chan int) {
	fmt.Println(" ", <-z)
}
