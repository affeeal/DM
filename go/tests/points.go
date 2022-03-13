package main

import (
	"fmt"
	"math"
)

type Point2D struct {
	x, y int
}

type Point3D struct {
	x, y, z int
}

func (p Point2D) Len() float64 {
	return math.Sqrt(float64(p.x * p.x + p.y * p.y))
}

func (p Point3D) Len() float64 {
	return math.Sqrt(float64(p.x * p.x + p.y * p.y + p.z * p.z))
}

type Measurable interface {
	Len() float64
}

func main() {
	var a Point2D = Point2D { 3, 4 }
	var b Point3D = Point3D { 12, 0, 5 }

	var m2 Measurable = a
	var m3 Measurable = b

	fmt.Printf("%d\n", int(m2.Len()))
	fmt.Printf("%d\n", int(m3.Len()))
}