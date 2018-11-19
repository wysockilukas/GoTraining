package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	wg.Add(2)
	go petla1()
	go petla2()

	fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("about to exit")
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())
}

func petla1() {
	var i int
	for i < 100 {
		runtime.Gosched()
		fmt.Println(1, "\t", i)
		i++
	}
	wg.Done()
}

func petla2() {
	var i int
	for i < 100 {
		runtime.Gosched()
		fmt.Println(2, "\t\t", i)
		i++
	}
	wg.Done()
}
