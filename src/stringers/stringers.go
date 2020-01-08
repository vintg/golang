package main

import (
	"fmt"
	"strings"
)

// IPAddr 4 numbers
type IPAddr [4]byte

func (a IPAddr) String() string{
	var b strings.Builder
	for i := range a{
		fmt.Fprint(&b, a[i], ".")
	}
	s := b.String()
	s = s[:b.Len()-1]
	return fmt.Sprintf(s)
}

func main() {
	hosts := map[string]IPAddr{
		"loopback" : {127,0,0,1},
		"googleDNS": {8,8,8,8},
	}
	for name, ip := range hosts{
		fmt.Printf("%v: %v\n", name, ip)
	}
}