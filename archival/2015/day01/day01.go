package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic("Bad file")
	}

	floor := 0
	basementIdx := -1
	for i, char := range input {
		if char == '(' {
			floor++
		} else if char == ')' {
			floor--
		}
		if floor < 0 && basementIdx < 0 {
			basementIdx = i + 1
		}
	}

	fmt.Printf("Part 1: %d\n", floor)
	fmt.Printf("Part 2: %d\n", basementIdx)
}
