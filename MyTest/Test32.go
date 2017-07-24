package main

import (
	"fmt"
)

//接口
type realize func(int, *int)

//实现类
type virtual interface {
	Handler(int, *int)
}
func (f realize) Handler(w int, r *int) {
	f(w, r)
}

func main() {
	num := 1
	f := funchain(funchain(realize(hello)))

	//f类指针指向实现类，调用实现类方法
	f.Handler(num,&num)
}

//参数为实现类，返回值为实现类
func funchain(next virtual) virtual{

	fn := func(w int, r *int) {
		fmt.Println("1")
		next.Handler(w, r)
	}

	//将fn转化为实现类引用，并返回
	return realize(fn)
}

//未转化的实现类
func hello(w int, r *int) {
	fmt.Println("hello")
}