package main

import "fmt"

func main() {
	test()
}

func treturn() int{
	fmt.Println("return")
	return 1
}

func test() int {
	defer func() {
		fmt.Println("defer")
	}()
	return  treturn()
}