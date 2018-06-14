package main

import (
	"time"
	"fmt"
)

func main() {
	//定时器表示在未来某一个时候的独立事件。你告诉定时器需要等待的事件
	//然后他将提供一个用于通知的通道
	//这里的定时器将等待2秒

	time1:= time.NewTimer(time.Second *2)
	//<- time1.C 直到这个定时器的通道C 明确发送定时器失效的值之前，将一直阻塞
	<- time1.C
	fmt.Println("Timer 1 expired")

	time2 := time.NewTimer(time.Second)
	go func() {
		<- time2.C
		fmt.Println("1Timer 2 expired")
	}()
	stop2 := time2.Stop()
	if stop2 {
		fmt.Println("2Timer 2 stopped")
	}

}
