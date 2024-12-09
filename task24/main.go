// Разработать программу нахождения расстояния между двумя точками,
//  которые представлены в виде структуры Point с инкапсулированными
//  параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

// инкапсуляция достигается за счет того, что поля x, y являются приватными, то есть доступными только для данного пакета
// Благодаря тому, что они инкапсулированы нельзя произвольно задавать эти поля в рамках другого пакета.
type Point struct {
	x, y float64
}

func New(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) Distance(other *Point) float64 {
	dx := other.x - p.x
	dy := other.y - p.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point1 := New(5, 6)
	point2 := New(4, 2)

	distance := point1.Distance(point2)

	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
