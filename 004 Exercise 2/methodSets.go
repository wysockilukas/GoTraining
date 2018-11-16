package main

import (
	"fmt"
)

type person struct {
	fName string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println((*p).fName, "mowi")

}

func saySomething(h human) {
	h.speak()
}

func main() {

	p1 := person{
		fName: "Kasia",
	}

	//fmt.Println()
	(&p1).speak()
	p1.speak()
	saySomething(&p1)

}
