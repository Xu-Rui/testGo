package main

import "fmt"

func main() {
	kvmap := make(map[string]string)
	kvmap["asd"] = "123"
	fmt.Println(kvmap)
	if v,ok := kvmap["asd"];ok{
		fmt.Println(v)
	}
}
