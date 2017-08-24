package main

import (
	"net/http"
	//"fmt"
	//"io/ioutil"
	"fmt"
	"io/ioutil"
)

var data[] byte

func main(){

	response,err:=http.Get("http://www.baidu.com")
	if err != nil{
		fmt.Println(err)
	}
	defer response.Body.Close()
	data,_= ioutil.ReadAll(response.Body)

	http.HandleFunc("/",Baidu)
	http.ListenAndServe(":80",nil)

}

func Baidu(w http.ResponseWriter,req *http.Request){
	w.Write(data)
}
