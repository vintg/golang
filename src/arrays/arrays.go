package main

import (
	"fmt"
	"strings"
)

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [7]int{2, 3, 5, 7, 11, 13, 17}
	fmt.Println(primes)

	test := []int{1, 2, 3}
	fmt.Println(test)

	var slice []int = primes[1:3]
	fmt.Println(slice)

	arrOfStruct := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
		{17, true},
	}
	fmt.Println(arrOfStruct)

	testArr :=[10]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(testArr[:5], testArr[5:], testArr[:])

	s := testArr[:6]
	printSlice(s)

	var z []int // empty array = nil
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}

	// make slices directly, optional capacity 3rd arg
	sa := make([]int,5,10) // len(sa) = 5, cap 10 
	printSlice(sa)

	makeBoard()
}

func printSlice(s []int) { // functions can be after main
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func makeBoard(){
	board:=[][]string{
		[]string{"-","-","-"},
		[]string{"-","-","-"},
		[]string{"-","-","-"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
