package main

import "fmt"

func main() {
	// var 可以声明一个或者是多个变量
	var a string = "initial"
	fmt.Println(a)
	// var 可以一次性声明多个变量
	var b, c int = 1, 2
	fmt.Println(b,c)

	// Go将自己推断已经初始化的变量类型
	var d = true
	fmt.Println(d)

	// 声明之后没有给出相应的初始值 变量将初始话零值  int的零值就是0
	var e int
	fmt.Println(e)

	// := 语法是声明并初始变量的简写
	f :="haha"  // 等价于 var f string = "haha"
	fmt.Println(f)

}
