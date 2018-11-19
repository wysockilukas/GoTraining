package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var i int64
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
	atomic.AddInt64(&i, 1)
	//fmt.Println(atomic.LoadInt64(&i))
	wg.Done()
}
