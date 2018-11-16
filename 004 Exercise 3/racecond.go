package main

import (
	"fmt"
	"runtime"
	"sync"
)

var i int = 0
var wg sync.WaitGroup

func main() {

	gs := 100

	wg.Add(gs)

	for x := 1; x <= gs; x++ {
		go fnInrement(x)

	}

	wg.Wait()

	fmt.Println(i)
}

func fnInrement(a int) {
	z := i
	runtime.Gosched()
	z++
	i = z
	//fmt.Println(a, i)
	wg.Done()
}
