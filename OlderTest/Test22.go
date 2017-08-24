package main

import "fmt"

func fibonac(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func print(c,quit chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	quit <- 0
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go print(c,quit)
	fibonac(c, quit)
}
