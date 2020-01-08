package main

import (
	"math/rand"
	"golang.org/x/tour/pic"
)

// Pic makes nested slices of unsigned 8 bit int of len dy
// each element contains slice of len dx
func Pic(dx, dy int) [][]uint8 {
	s:= make ([][]uint8, dy)
	
	for i:=0;i<dy;i++{
		s[i] = make([]uint8, dx)
		for j:=0;j<dx;j++{
			s[i][j] = uint8(rand.Intn(100))
		}
	}

	return s
}

func main(){
	pic.Show(Pic)
}