package main

import "fmt"

func main() {
	a,b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//仅仅想返回值的一部分的话，你可以使用空白定义符 _
	_, c:= vals()
	fmt.Println(c)

}

func vals() (int,int)  {
	return 3,7
}


