package main

import "fmt"

//默认通道是 无缓冲的 这意味着只有对用的就收(<- chan)通道准备好接收时，才允许进行发送(chan <-).
//可缓存通道允许在没有对应接收方的情况下，缓存限定数量的值

func main() {

	//make 一个通道 最多允许缓存两个值
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
