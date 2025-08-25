package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	point := new(Point)
	point.x = x
	point.y = y
	return point
}

func (p Point) Distance(point *Point) float64 {
	dist := math.Sqrt((math.Pow(point.x, 2) - 2*p.x*point.x + math.Pow(p.x, 2)) + (math.Pow(point.y, 2) - 2*p.y*point.y + math.Pow(p.y, 2)))
	return dist
}

func main() {
	TochkaA := NewPoint(1, 3)
	TochkaB := NewPoint(2, 4)
	fmt.Println(TochkaA.Distance(TochkaB))
}
