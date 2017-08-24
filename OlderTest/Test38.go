package main

import (
	"net/http"
	"fmt"
)

func main() {
	fun := getHandlerMapping("123")
	fmt.Println(fun == nil)
}


func getHandlerMapping (urlStr string) (http.Handler) {
	handlerMapping := make(map[string]http.HandlerFunc)
	fmt.Println(handlerMapping["123"] == nil)
	return handlerMapping[urlStr]
}

// 注意空接口指针