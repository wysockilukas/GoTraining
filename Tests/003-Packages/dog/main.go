package main

import (
	"fmt"
	"runtime"

	"github.com/wysockilukas/GoTraining/003-Packages/cat"
	"github.com/wysockilukas/GoTraining/003-Packages/trzeci"
)

func main() {

	fmt.Println(runtime.GOOS, runtime.GOARCH)
	cat.FromCat()
	fmt.Println(trzeci.X)
	trzeci.Abc()

	/*
	   komenda do kompilacji
	   GOOS=linux GOARCH=amd64 go build main.go
	   GOOS=windows GOARCH=386 go build -o hello.exe hello.go
	*/
}
