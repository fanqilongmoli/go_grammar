package main

import "fmt"

// go的结构体是各个字段的类型集合。这在组织数据是非常有用

type person struct {
	name string
	age  int
}

func main() {
	//使用这个语法创建了一个新的结构体元素。
	fmt.Println(person{"Bob", 20})

	//初始化一个结构体元素时指定字段名字。
	fmt.Println(person{name: "wang", age: 66})

	//省略的字段将被初始化为零值。
	fmt.Println(person{name: "fred"})

	//& 前缀生成一个结构体指针
	fmt.Println(&person{name: "ann", age: 29})

	s := person{name: "sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	//结构体指针使用. - 指针会被自动解引用
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

}
