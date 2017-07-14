package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"regexp"
)

func main(){
	house("http://bj.58.com/ershoufang/?PGTID=0d30000c-0000-16fc-1f3f-e51b929055a9&ClickID=1")

}

//    Hello 世界！123 G"o


func house(url string) {
	response, _ := http.Get(url)
	data, _ := ioutil.ReadAll(response.Body)
	str := string(data)

	reg := regexp.MustCompile(`<!--房源列表信息-->(.|\s)*<div class="page">`)
	//reg := regexp.MustCompile(`<!--房源列表信息-->...<div class="page">`)
	fmt.Println(reg.FindAllString(str, -1))

	reg = regexp.MustCompile(`class="next".*<span>下一页</span>`)
	urlStrArray := reg.FindAllString(str, -1)
	if len(urlStrArray)>0 {
		temp := []byte(urlStrArray[0])
		urlStr := string(temp[19:len(temp)-24])
		fmt.Println(urlStr)
		house(urlStr)
	}

	defer response.Body.Close()
}