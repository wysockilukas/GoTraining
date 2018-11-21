package main

import (
	"fmt"
	"sync"
)

var i int
var wg sync.WaitGroup
var mu sync.Mutex

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
	mu.Lock()
	z := i
	//runtime.Gosched() - przy mutex to nie jest potrzebne
	z++
	i = z
	//fmt.Println(a, i)
	mu.Unlock()
	wg.Done()
}
