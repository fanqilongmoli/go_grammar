package main

import (
	"errors"
	"fmt"
)

//go 使用一个独立明确的返回值来传递错误信息
//这与使用异常的Java和Ruby以及在C语言中经常见到的
//超重的单返回值/错误 go语言的处理方式能清楚的知道那个函数返回了错误
//并能像调用那些没有出错的函数一样调用

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("42 you wen ti de!")
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg: arg, prob: "42 you wen ti de!"}
	}

	return arg + 3, nil
}

func main() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(4)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
