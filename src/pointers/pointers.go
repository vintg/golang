package main

import (
	"fmt"
)

func pointers(){
	i, j := 42, 2701

	p := &i                   // point to operand i
	fmt.Println("i init", *p) // read i through pointer p

	*p = 21                                 //set i through pointer p
	fmt.Println("i set through p to 21", i) // i should be new val

	p = &j                                  // set pointer to j
	*p = *p / 37                            // divide j through pointer
	fmt.Println("j div from 2701 by 37", j) // new val of j
}

func main() {
	pointers()
}
