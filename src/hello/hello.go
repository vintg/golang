package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum / 2
	y = sum - x
	return
}

func cToF(celsius int) int {
	return celsius*9/5 + 32
}

/*
  Big: 
	Create a huge number by shifting a 1 bit left 100 places.
	In other words, the binary number that is 1 followed by 100 zeroes.
*/
const ( // vars can be declared outside of main
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("now you have %g problems.\n", math.Sqrt(7))
	fmt.Println("value of pi", math.Pi)

	a, b := swap("a", "b")
	fmt.Println("a,b swapped is", a, b)

	fmt.Printf("split 21 is ")
	fmt.Println(split(21))

	fmt.Println("36 celsius to fahrenheit is", cToF(36))

	var typetest, t2, t3 bool
	fmt.Println("var type list", typetest, t2, t3)

	var c, py, java, js, golang = false, 1, "no!", true, true
	fmt.Println("initialized vars auto type", c, py, java, js, golang)

	/* go types
	bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte // alias for uint8

	rune // alias for int32
			// represents a Unicode code point

	float32 float64

	complex64 complex128
	*/

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z2 uint = uint(f)
	fmt.Println(x, y, z2)

	x = 3
	fmt.Println("x reassigned", x)	

	const Pi = 3.14159265358979323846264 // has built in assignment blocking in editor
	fmt.Println("Pi const", Pi)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
