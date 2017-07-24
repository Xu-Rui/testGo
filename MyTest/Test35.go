package main

import "fmt"

func main() {
	str := "\\n"
	fmt.Println([]byte(str))
	var strByte[4]byte = [4]byte{'a','b','c','\\'}
	//str := fmt.Sprintf("%s",strByte)
	fmt.Printf("%q",strByte)
}
