package main

import (
	"fmt"
	"math"
)

type iShape interface {
	Name() string
	GetArea() float64
	String() string
}
type shape struct {
	name   string
	height float64
	width  float64
}
type triangle struct {
	shape
}
type square struct {
	shape
}
type circle struct {
	shape
}

func (s *shape) Name() string {
	return s.name
}
func (t triangle) GetArea() float64 {
	return (0.5 * t.height * t.width)
}
func (t triangle) String() string {
	return fmt.Sprintf("\tName: %s,\n\tHeight: %f ,\n\tWidth: %f,\n\tArea:%f", t.Name(), t.height, t.width, t.GetArea())
}
func (t square) GetArea() float64 {
	return (t.height * t.width)
}
func (t square) String() string {
	return fmt.Sprintf("\tName: %s,\n\tHeight: %f ,\n\tWidth: %f,\n\tArea:%f", t.Name(), t.height, t.width, t.GetArea())
}
func (t circle) GetArea() float64 {
	return (math.Pi * (t.height / 2))
}
func (t circle) String() string {
	return fmt.Sprintf("\tName: %s,\n\tHeight: %f ,\n\tWidth: %f,\n\tArea:%f", t.Name(), t.height, t.width, t.GetArea())
}

func main() {
	t := triangle{shape{"Triangle", 10, 5}}
	c := circle{shape{"Circle", 5, 5}}
	s := square{shape{"Square", 10, 10}}
	fmt.Println(t.String())
	fmt.Println(c.String())
	fmt.Println(s.String())
}
