package main

import (
	"log"
	"net/http"
	"fmt"
)



type String string

type Struct struct{
	Greeting 	string
	Punct		string
	Who			string
}

func (s Struct)ServerHTTP(
	w http.ResponseWriter,
	r *http.Request){
	fmt.Println("asda")
}

func main() {
	// your http.Handle calls here
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
}
