package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(sequence(34))
	fmt.Println(fct(17))
}

func sequence(n uint) uint {

	res := []uint{1, 2}
	var nom uint
	var denom uint
	var newItem uint
	var i uint
	//var a int
	//var b int
	for i = 2; i <= n; i++ {
		//a = int(res[i-2])
		//b = int(res[i-1])
		nom = 6 * res[i-2] * res[i-1]
		denom = (5*res[i-2] - res[i-1])
		if denom == 0 {
			return 0
		} else {
			newItem = nom / denom
		}
		res = append(res, uint(newItem))
	}
	return res[n]
}

func fct(n uint) uint {
	return uint(math.Pow(2, float64(n)))
}
