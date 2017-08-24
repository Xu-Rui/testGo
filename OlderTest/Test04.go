package main

import (
	"fmt"
	"strings"
)

func main(){

	var str string

	fmt.Scanf("%s",&str)

	var temp_1 = strings.Split(str,",")

	temp_2 := strings.Replace(str," ","%2",-1)

	fmt.Println(temp_1)
	fmt.Println(temp_2)
}
