package main

import (
	"math"
	"fmt"
)

// Abser is a set of method signatures
type Abser interface {
	Abs() float64
}

// MyFloat for custom float methods
type MyFloat float64


// Vertex coordinate pair
type Vertex struct {
	x, y float64
}

// T struct
type T struct {
	S string
}

// I interface
type I interface{
	Method()
}


// Method of type T
func (t T) Method(){
	fmt.Println(t.S)
}

// Abs method for MyFloat
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Abs vertex abs method
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat
	a = &v // a *Vertex

	fmt.Println(a.Abs())

	var i I = T{"hello"}
	i.Method()
}
