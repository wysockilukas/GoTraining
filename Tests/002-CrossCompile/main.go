package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println(runtime.GOOS, runtime.GOARCH)
	/*
	   komenda do kompilacji
	   GOOS=linux GOARCH=amd64 go build main.go
	   GOOS=windows GOARCH=386 go build -o hello.exe hello.go
	*/
}
