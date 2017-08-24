package main

import (
	"time"
	"fmt"
)

func main() {
	theTime, err := time.Parse("2016-06-01 00:00:01","2016-06-01 00:00:01")
	fmt.Println(err)
	fmt.Println(theTime.String())
}


/*
	时间戳转换测试
	结论:时间戳转换为字符串以后就只能是个字符串了

*/