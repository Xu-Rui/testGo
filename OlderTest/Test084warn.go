package main

import (
	"net/http"
	"fmt"
	"time"
	"strconv"
)


func main() {
	check()
	single()
}

func check(){
	var response *http.Response
	response, err := http.Get("http://wiki.intra.xiaojukeji.com/pages/viewpage.action?pageId=98022322")

	for err != nil {
		response, err = http.Get("http://wiki.intra.xiaojukeji.com/pages/viewpage.action?pageId=98022322")
		if err != nil {
			fmt.Print(time.Now())
			fmt.Println(" faile")
		}
		time.Sleep(1000000000)
	}
	fmt.Print(time.Now())
	fmt.Println(" succeed")
	defer response.Body.Close()
}

func single(){
	hour,min,sec := time.Now().Clock()
	nowTime := strconv.Itoa(hour)+":"+strconv.Itoa(min)+":"+strconv.Itoa(sec)
	str := "http://118.89.184.104:8092/SendMessage" + "?" + "time="+nowTime
	//str := "http://localhost:8092/SendMessage" + "?" + "time="+nowTime
	http.Get(str)
}