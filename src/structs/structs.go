package main

import (
	"fmt"
)

// Vertex struct contains a pair of int coordinates
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}  // Y:0 def/implicit
	v3 = Vertex{}      // default 0,0
	p  = &Vertex{1, 2} // type *Vertex
)

func main() {
	v := Vertex{1, 6}
	v.X = 5
	p := &v
	p.Y = 10
	fmt.Println(v)

	fmt.Println(v1, p, v2, v3)
}
