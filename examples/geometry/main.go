package main

import "fmt"


type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base float64
}
type square struct {
	sideLength float64
}

func main()  {
	tr := triangle{
		height: 10,
		base: 10,
	}
	sq := square{
		sideLength: 10,
	}
	printArea(tr)
	printArea(sq)

}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func (t triangle) getArea()  float64{
	return t.height * t.base * 0.5
}

func (s square) getArea()  float64{
	return s.sideLength * s.sideLength
}