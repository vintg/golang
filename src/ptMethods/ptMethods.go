package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Vertex coordinate pair
type Vertex struct {
	x, y float64
}

// methods with value receivers may take either value or pointer

// Abs Vertex.Abs()
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// Scale multiplies coords by a constant factor
func (v *Vertex) Scale(l float64) {
	v.x *= l
	v.y *= l
}

// non method funcs must take inputs of declared type

// Abs regular func with vertex input
func Abs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// Scale regular func with vertex input
func Scale(v *Vertex, f float64) {
	v.x *= f
	v.y *= f
}

// all methods should have either value or pointer receiveers (args) but not a mixture

func main() {
	v := Vertex{1, 2}
	c := float64(rand.Intn(1337))

	fmt.Println("original", v.Abs())

	v.Scale(c)
	fmt.Println(v.Abs())

	Scale(&v, 1/c)
	fmt.Println("rescaled", Abs(v))

	p := &Vertex{3,4}
	p.Scale(3)
	Scale(p,8)

	fmt.Println(v,p)
}
