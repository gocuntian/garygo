package main

import "fmt"
import "math"

type geometry interface {
	area() float64
	perim() float64
}

type square struct {
	width float64
	height float64
}

type circle struct {
	radius float64
}

//正方形实现的接口
func (s square) area() float64 {
	return s.width * s.height
}

func (s square) perim() float64 {
	return 2*s.width + 2*s.height
}

//圆形实现的接口
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main(){
	s := square{width: 3, height: 4}
	measure(s)
	c := circle{radius: 5}
	measure(c)
}

// {3 4}
// 12
// 14
//=========================
// {5}
// 78.53981633974483
// 31.41592653589793