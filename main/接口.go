package main

import (
	"math"
	"fmt"
)

//

//定义一个几何基本接口
type geometry interface {
	area() float64
	perim() float64
}

//------定义两个结构
type rect_rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

//----- 实现定义的接口
func (r rect_rect) area() float64 {
	return r.width * r.height
}
func (r rect_rect) perim() float64 {
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {

	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry)  {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {

	r := rect_rect{width:3,height:4}
	c := circle{radius:5}

	measure(r)
	measure(c)
}
