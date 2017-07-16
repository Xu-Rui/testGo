package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make (chan int, 1)
	ch2 := make (chan int, 1)


	go func() {
		for{

			time.Sleep(200)
			ch1 <- 2

			time.Sleep(200)
			ch2 <- 2
		}
	}()

	for {
		select {
		case <-ch1:
			fmt.Println("ch1 pop one element")
		case <-ch2:
			fmt.Println("ch2 pop one element")
		}
	}
	num := <- ch1

	fmt.Println(num)
}

/*
	结论若是

	select {
	case <-ch1:
		fmt.Println("ch1 pop one element")
	case <-ch2:
		fmt.Println("ch2 pop one element")
	}

	1，	这样调用select case
		则管道内的数据会被浪费掉

	2，	若程序判断管道中不会有数据通过，则会关闭管道
		此时再次读取管道数据会报异常

*/