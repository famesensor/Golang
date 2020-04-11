package main

import (
	"fmt"
	"math"
)

func structure() {
	// var c Circle
	// ci := new(Circle)
	cycle := Circle{0, 0, 5}
	fmt.Println(cycle.x, cycle.y, cycle.r)
	fmt.Println(circleArea(cycle), "\n", cycle.x)
	PcircleArea(&cycle)
	fmt.Println(cycle.x)
}

type Circle struct {
	x, y, r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func PcircleArea(c *Circle) float64 {
	c.x = math.Pi * c.r * c.r
	return c.x
}
