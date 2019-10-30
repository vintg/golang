package main

import (
	"fmt"
	"math"
	"runtime"
)

func sum1toN(n int) int {
	sum := 0 // : shortcut for var declaration
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}

func doubleUnderN(n int) int {
	// while loops written as for conditions
	// empty for -> inf loop

	sum := 1
	for sum <= n {
		sum += sum
	}
	return sum
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// vars can have executions before condition
	// and are block scoped
	if v := math.Pow(x, n); v < lim {
		lim = v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func sqrtFinder(x float64) float64 {
	z := float64(1)

	for i := 1; i <= 10; i++ {
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func envCheck() {
	switch os := runtime.GOOS; os { // if condition left out, matches any true
	case "darwin":
		fmt.Println("os x")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s. \n", os)
	}
}

func main() {
	var n = 21
	var res = sum1toN(n)
	fmt.Println("sum loop 1 to", n, "=", res)
	res = doubleUnderN(n)
	fmt.Println(
		"double while <=", n, "=", res, "\n",
		"2^.5 =", sqrt(2), "\n",
		"-4^.5=", sqrt(-4), "\n",
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println("newton sqrt", sqrtFinder(2), sqrtFinder(9))
	envCheck()

	for x := 1; x <= 10; x++ {
		defer fmt.Print(x, " ") // defer is push pop stack
	}

	defer fmt.Print("peaches\n")
	fmt.Print("pat is ")

}
