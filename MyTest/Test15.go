package main

import (
	"encoding/json"
	"bytes"
	"fmt"
)

func jsonDecode(in []byte, out interface{}) (err error) {
	//获取一个读取字节数组的json解码器
	decoder := json.NewDecoder(bytes.NewReader(in))
	//将json数字转化为number类型
	decoder.UseNumber()
	//将解码结果保存在out里，out必须使用指针，不然只是值拷贝
	err = decoder.Decode(out)
	return
}

func main() {
	//创建参数集合
	paramset := make(map[string]interface{})
	//得到传来的参数
	rawValue := "{\"a\": \"b\",\"c\": \"d\",\"e\": \"f\"}"
	//Json反序列化
	jsonDecode([]byte(rawValue), &paramset)

	fmt.Println(paramset)
}