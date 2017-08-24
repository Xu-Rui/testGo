package main

import (
	"strings"
	"fmt"
)

func main() {
	kvMap := make(map[string]string)
	resolveData("set asd asd",kvMap)
}

func temp(data string,kvMap map[string]string) string{

	dataSplit := strings.Split(data," ")

	fmt.Println(dataSplit)

	switch dataSplit[0] {
	case "set":
		kvMap[dataSplit[0]] = dataSplit[1]
	case "get":
		return kvMap[dataSplit[0]]
	}
	return "error"
}
