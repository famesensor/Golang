package main

import (
	"fmt"
	"math"
)

func method() {
	c := Circle{0, 0, 10}
	rac := Ractangle{20, 10}
	fmt.Println(c.area())
	fmt.Println(rac.area())
	p := Person{"Fame"}
	fmt.Println(p.Name)
	r := new(Robot)
	r.Name = "God speed"
	r.Talk()

	// interface
	fmt.Println(totalAarea(&c, &rac))
}

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Ractangle struct {
	l, w float64
}

func (r *Ractangle) area() float64 {
	return r.w * r.l
}

func totalAarea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

type Person struct {
	Name string
}

type Robot struct {
	Person
	Model string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}
