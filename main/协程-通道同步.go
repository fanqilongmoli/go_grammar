package main

import (
	"fmt"
	"time"
)

//我们可以使用通道来同步Go协程之间的执行状态，这里是一个使用阻塞的接收方式来等待一个Go协程的运行结束

func worker(done chan bool) {

	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool, 1)

	go worker(done)

	<-done
}
