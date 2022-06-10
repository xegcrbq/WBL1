package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры
//Point с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p1 *Point) GetDistanceToPoint(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}
func main() {
	p1, p2 := NewPoint(0, 0), NewPoint(2.5, 2.5)
	fmt.Printf("Distance between %#v and %#v is %.4f", p1, p2, p1.GetDistanceToPoint(p2))
}
