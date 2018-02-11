package main

import (
	"fmt"
	"time"
)

func main() {
	i:=2
	fmt.Print("Write",i," as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is the weekend")
	default:
		fmt.Println("it is weekday")
	}

	/**
		不带表达式的switch 是实现if/else逻辑的另一种方式
		case的表达式也可以不使用常量
	 */
	t:= time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}




}
