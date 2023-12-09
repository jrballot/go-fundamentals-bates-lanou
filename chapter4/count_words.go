package main

import (
	"fmt"
	"strings"
)


func main() {
	counts := map[string]int{}

	sentence := "The quick brown fox jumps over the lazy dog"

	words := strings.Fields(strings.ToLower(sentence))

	for _, w := range words {
		counts[w] = counts[w] + 1
	}

	fmt.Println("Counts: ",counts)
}