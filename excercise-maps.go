package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int)
	for i := 0; i < len(words); i++ {
		x := words[i]
		if counts[x] != 0 {
		 	counts[x] += 1
		} else {
			counts[x] = 1
		}
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
