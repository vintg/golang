package main

import (
	"fmt"
	"math"
)

// Vertex struct is a coordinate pair
type Vertex struct {
	x, y float64
}

// Abs returns sqrt of coordinates ^2 to ensure > 0
// input type struct Vertex
// note the fn declaration indicates it is a method for struct Vertex
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// Abs contains the same functionality as the struct method but as a regular fn
func Abs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// Flip returns mirror over y=x diagonal
func (v Vertex) Flip() Vertex {
	flipped := Vertex{v.y, v.x}
	return flipped
}

// MyFloat non-struct type ext float64
type MyFloat float64

// Abs is a method ext myFloat 
func (myf MyFloat) Abs() float64 {
	if myf < 0 {
		return float64(-myf)
	}
	return float64(myf)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs(), "=", Abs(v))

	vi := v.Flip()
	fmt.Println(v, "flipped over y=x", vi)

	f:= MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
