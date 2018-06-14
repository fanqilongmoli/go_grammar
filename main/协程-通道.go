package main

import "fmt"

//通道是连接多个Go协程的管道，你可以从一个Go协程将值发送到通道
//然后在别的Go协程中接受

func main() {
	// 使用 make(chan val-type) 创建一个新的通道
	// 通道类型就是我们需要传递的值得类型

	messages := make(chan string)

	//使用 channel <- 语法 发送 一个新的值到通道中。
	go func() {messages <- "ping"}()

	//使用 <-channel 语法从通道中 接收 一个值。
	msg := <-messages
	fmt.Println(msg)
}
