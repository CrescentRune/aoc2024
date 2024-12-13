package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Memory map[uint64]int

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic("File no good! File BAD!")
	}

	stones := createStonesFromInput(input)

	p1 := blinkAtStones(stones, 25)
	fmt.Printf("Part 1 solution: %d\n", p1)

	p2 := blinkAtStones(stones, 75)
	fmt.Printf("Part 2 solution: %d\n", p2)
}

func blinkAtStones(stones []uint64, blinks int) int {
	memory := make(Memory)

	for _, stone := range stones {
		memory[stone] = 1
	}

	for range blinks {
		nextMem := make(Memory)
		for key, val := range memory {
			count, v1, v2 := blinkAtStone(key)
			nextMem[v1] += val
			if count > 1 {
				nextMem[v2] += val
			}
		}

		memory = nextMem
	}

	sum := 0
	for _, val := range memory {
		sum += val
	}

	return sum
}

func blinkAtStone(stone uint64) (int, uint64, uint64) {
	if stone == 0 {
		return 1, 1, 0
	} else if width := intWidth(stone); width%2 == 0 {
		lhs, rhs := splitInt(stone, width)
		return 2, lhs, rhs
	} else {
		return 1, stone * 2024, 0
	}
}

func intWidth(val uint64) int {
	return int(len(strconv.FormatUint(val, 10)))
}

func splitInt(val uint64, width int) (uint64, uint64) {
	lhs := val / uint64(math.Pow10(width/2))
	rhs := val % (lhs * uint64(math.Pow10(width/2)))
	return lhs, rhs
}

func createStonesFromInput(input []byte) []uint64 {
	re := regexp.MustCompile(`\d+`)
	results := re.FindAll(input, -1)
	stones := []uint64{}
	for _, res := range results {
		if stoneVal, err := strconv.Atoi(string(res)); err == nil {
			stones = append(stones, uint64(stoneVal))
		}
	}

	return stones
}
