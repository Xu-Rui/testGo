package main

import "fmt"

func main() {
	numArray := make([]int,10)
	fmt.Println(len(numArray))
	fmt.Println(cap(numArray))
	fmt.Println(numArray)
}
