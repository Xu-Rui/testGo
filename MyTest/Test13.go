package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	count := 1
	f0 := 1
	f1 := 1
	f2 := 2

	return func() int {

		if count < 3 {
			count++
			return f0
		} else {
			count++
			f0 = f1
			f1 = f2
			f2 = f1+f0
		}
		return f1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
