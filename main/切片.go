package main

import "fmt"

//slice 是go中关键的数据类型 是一个比数组更强大的序列接口

func main() {
	// 要创建一个长度非零的空slice，需要使用内建的方法 make
	s := make([]string, 3)
	fmt.Println("emp", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set", s)
	fmt.Println("get", s[2])
	fmt.Println("len", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy", c)

	//slice[low:high] 语法进行“切片”操作
	l := s[2:5]
	fmt.Println("sl1", l)

	l = s[:5] //从第一个开始
	fmt.Println("sl2:", l)

	l = s[2:] // 切到最后一个
	fmt.Println("sl3:", l)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d",twoD)
}
