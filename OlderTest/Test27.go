package main

import (
	"time"
	"fmt"
)

func main() {

	ticker := time.NewTicker(1)
	for i := 0; i < 10; i++ {
		now := time.Now()
		<-ticker.C
		ticker.Stop()
		fmt.Println(time.Now().Sub(now))
	}
}


/*
	结论
		time := <-ticker.C
		t := time.Now()
		若是定义和报名相同的变量名


		ticker 定时器的作用
			创建定时器的时候参数中传递一个整形参数代表时间周期毫秒值
			每经过一个时间周期ticker.C这个通道内就会有时间字符串传递过来
			又由于通道的阻塞性,所以ticker起到了精确延时的作用
			可以通过ticker.stop停止定时

			缺点：毫秒级别不是很精确，可能是由于GC的影响
*/