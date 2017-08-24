package main

import "fmt"

func main() {
	defer func(){
		err := recover()
		if err!=nil{
			fmt.Println(err)
		}
	}()

	panicTest()
}


func panicTest(){
	panic("panic Test")
}