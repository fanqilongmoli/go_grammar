package main

import "fmt"

//go 协程在执行上来说是轻量级的线程

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from,":",i)
	}
}

func main() {
	f("direct")

	//使用go f(s) 在一个Go协程中调用这个函数。这个新的Go协程将会并行的执行这个函数的调用
	go f("goroutine")

	//匿名函数启动Go 协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//

	var input string

	fmt.Scanln(&input)
	fmt.Println("done")

}