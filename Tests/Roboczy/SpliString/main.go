package main

import (
	"fmt"
)

func main() {

	fmt.Println(solution("abcdef"))

}

func solution(str string) []string {
	var res []string

	for i := 0; i < len(str)-1; i = i + 2 {
		res = append(res, str[i:i+2])
	}
	if len(str)%2 != 0 {
		res = append(res, string(str[len(str)-1])+"_")

	}

	return res
}
