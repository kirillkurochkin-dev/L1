package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Вычисляем расстояние
func (p1 Point) DistanceTo(p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	point1 := NewPoint(0, 0)
	point2 := NewPoint(3, 4)

	distance := point1.DistanceTo(point2)

	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
