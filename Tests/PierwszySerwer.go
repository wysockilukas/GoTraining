package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {

	Licznik := 1

	fmt.Println("print")

	FunkcjaWykonywanaNaWWW := func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("<html><body><h1>" + strconv.Itoa(Licznik) + "</h1></body></html>"))
		Licznik++
	}

	http.HandleFunc("/GOWWW/", FunkcjaWykonywanaNaWWW)
	if err := http.ListenAndServe(":4444", nil); err != nil {
		log.Fatal("Jakis error", err)
	}

}

