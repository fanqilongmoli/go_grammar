package main

func main() {
	//for 是 Go 中唯一的循环结构

	i :=1
	for i<=3 {
		println(i)
		i++
	}

	//==========================

	for j := 7; j <= 9; j++ {
		println(j)
	}

	//==========================
	for {
		println("loop")
		break
	}

	//==========================

	for n := 0; n < 10; n++ {
		if n % 2 == 0 {
			continue
		}
		println(n)
	}
}