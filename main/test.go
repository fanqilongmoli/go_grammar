package main

import "fmt"

func main()  {
	/**/
	fmt.Println("你好我是樊启龙")

	const(
		a = iota // 0
		b   // 1
		c   // 2
		d = "ha" //独立值，iota += 1
 		e       //"ha"   iota += 1
		f = 100 //iota +=1
		g     //100  iota +=1
		h = iota //7,恢复计数
		i      //8
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)

	fmt.Println("===========Go 语言运算符===========")
	operator()
	fmt.Println("===========Go 语言 select 语句===========")
	selectyuju()

	fmt.Println("===========Go 语言 chan===========")
	chanTest()

	fmt.Println("===========Go 语言 语言循环语句===========")
	forTest()

}
/*for 循环*/
func forTest() {
	var  C,c int
	C = 1
	A:for C < 100 {
		C++
		for c = 2; c < C; c++ {
			if C%c == 0 {
				goto A
			}
		}
		fmt.Println(C,"是素数")
		}
}

/*chan*/
func chanTest() {
	/*
		chan是golang 非常重要的东西 是用来做goroutine的通信 因为golang程序必然会有多个goroutine
		chan 使用符号<- 进行读写   v,ok := <-c 是读取  c <- v 是写入
		读取时判断时候close使用 v,ok := <-c而不是用v := <-c
	*/

	// 读取chan

}

/*Go 语言运算符*/
func operator()  {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c )
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c )
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c )
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c )
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c )
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a )
	a=21   // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a )

}
/*Go 语言 select 语句*/
func selectyuju()  {
	var c1, c2, c3 chan int
	var  i1,i2  int
	select {
	case i1  = <- c1:
		fmt.Printf("received %d   from c1\n", i1)
	case c2 <- i2:
		fmt.Printf("sent %d  to c2\n", i2)
	case i3,ok := (<- c3):  // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received %d  from c3\n", i3)
		}else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
		
	}

}
