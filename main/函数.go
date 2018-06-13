package main

import "fmt"

func main() {

	res := plus(1,2)
	fmt.Println("plus",res)
}

func plus(a int, b int) int {
	return a + b
}