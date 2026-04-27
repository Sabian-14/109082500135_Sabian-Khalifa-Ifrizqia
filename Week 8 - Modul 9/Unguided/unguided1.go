package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

type Circle struct {
	center Point
	r      int
}

func jarak(p, q Point) float64 {
	dx := float64(p.x - q.x)
	dy := float64(p.y - q.y)
	return math.Sqrt(dx*dx + dy*dy)
}

func diDalam(c Circle, p Point) bool {
	return jarak(c.center, p) <= float64(c.r)
}

func main() {
	var c1, c2 Circle
	var p Point

	fmt.Scan(&c1.center.x, &c1.center.y, &c1.r)
	fmt.Scan(&c2.center.x, &c2.center.y, &c2.r)
	fmt.Scan(&p.x, &p.y)

	in1 := diDalam(c1, p)
	in2 := diDalam(c2, p)

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}