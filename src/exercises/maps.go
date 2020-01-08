package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

// WordCount returns a map of distinct word counts
func WordCount(s string) map[string]int {
	var out map[string]int
	out = make(map[string]int)

	var words []string = strings.Fields(s)

	for i := 0; i < len(words); i++ {
		ok := out[words[i]]

		switch ok {
		case 1:
			out[words[i]]++
		case 0:
			out[words[i]] = 1
		}
	}

	return out
}

func main() {
	wc.Test(WordCount)
}