package main

import (
	"time"
	"fmt"
)

// 超时对于一个连接外部资源，
// 或者其他一些需要花费执行时间的操作和程序而言是很重要的。
// 得益于通道和select，在Go中实现超时操作是简洁而优雅的

func main() {


	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "result 1"
	}()
	// 这里使用select实现一个超时操作
	// res := <- c1 等到结果， Time.After 等待超时时间 1 秒后发送的值
	// <- Time.After 等待超时时间 1 秒后发送的值
	// 由于select默认处理第一个已经准备好接收操作 如果这个操作超过
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

}
