package main

import (
	"fmt"
)

//常规的通过通道发送和接收数据是阻塞的。然而，我们可以
//使用带一个default字句的select来实现非阻塞的发送、接收
//甚至是非阻塞的多路select

func main() {

	messages := make(chan string)
	signals := make(chan bool)

	go func() {
		messages <- "da"

	}()

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "Hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
