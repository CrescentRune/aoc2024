package main

import (
	"bytes"
	"fmt"
	"os"
)

type Antennae struct {
	row       int
	col       int
	frequency string
}

type FrequencyList map[string][]*Antennae

type Set map[string]bool

type Bounds struct {
	width  int
	height int
}

func main() {
	input, err := os.ReadFile("input.txt")

	if err != nil {
		panic("Poop file. No good!\n")
	}

	freqList, bounds := parseInput(input)

	p1 := findAntinodes(freqList, bounds)

	fmt.Printf("Part 1 solution: %d\n", p1)
	fmt.Printf("Part 2 solution: %d\n", findHarmonicAntinodes(freqList, bounds))

}

func findHarmonicAntinodes(freqList FrequencyList, bounds *Bounds) int {
	foundMap := make(Set)

	for _, list := range freqList {
		for a := range len(list) - 1 {
			antennaeA := list[a]
			for b := a + 1; b < len(list); b++ {
				antennaeB := list[b]
				dRow, dCol := antennaeB.row-antennaeA.row, antennaeB.col-antennaeA.col

				findAllHarmonics(bounds, foundMap, antennaeA, antennaeB, dRow, dCol)
			}
		}
	}

	return len(foundMap)
}

func findAllHarmonics(bounds *Bounds, foundMap Set, a *Antennae, b *Antennae, dRow int, dCol int) {
	aAnti := fmt.Sprintf("%d,%d", a.row, a.col)
	foundMap[aAnti] = true
	bAnti := fmt.Sprintf("%d,%d", b.row, b.col)
	foundMap[bAnti] = true

	currRow, currCol := a.row, a.col
	for {
		candRow, candCol := currRow-dRow, currCol-dCol
		if !inBounds(bounds, candRow, candCol) {
			break
		}
		candName := fmt.Sprintf("%d,%d", candRow, candCol)
		foundMap[candName] = true
		currRow, currCol = candRow, candCol
	}

	currRow, currCol = b.row, b.col
	for {
		candRow, candCol := currRow+dRow, currCol+dCol
		if !inBounds(bounds, candRow, candCol) {
			break
		}
		candName := fmt.Sprintf("%d,%d", candRow, candCol)
		foundMap[candName] = true
		currRow, currCol = candRow, candCol
	}

	//return foundMap
}

func findAntinodes(freqList FrequencyList, bounds *Bounds) int {
	foundMap := make(map[string]bool)

	for _, list := range freqList {
		for a := range len(list) - 1 {
			antennaeA := list[a]
			for b := a + 1; b < len(list); b++ {
				antennaeB := list[b]
				dRow, dCol := antennaeB.row-antennaeA.row, antennaeB.col-antennaeA.col
				candARow, candACol := antennaeA.row-dRow, antennaeA.col-dCol
				if inBounds(bounds, candARow, candACol) {
					candA := fmt.Sprintf("%d,%d", candARow, candACol)
					foundMap[candA] = true
				}

				candBRow, candBCol := antennaeB.row+dRow, antennaeB.col+dCol
				if inBounds(bounds, candBRow, candBCol) {
					candB := fmt.Sprintf("%d,%d", candBRow, candBCol)
					foundMap[candB] = true
				}
			}
		}
	}

	return len(foundMap)
}

func inBounds(bounds *Bounds, row int, col int) bool {
	return row >= 0 && col >= 0 && row < bounds.height && col < bounds.width
}

func parseInput(input []byte) (FrequencyList, *Bounds) {
	rawParse := bytes.Split(input, []byte("\n"))
	bounds := &Bounds{len(rawParse), len(rawParse[0])}

	freqList := make(FrequencyList)

	for i, row := range rawParse {
		for j, char := range row {
			if char != '.' {
				frequency := string(char)
				newAntennae := &Antennae{i, j, frequency}
				freqList[frequency] = append(freqList[frequency], newAntennae)
			}
		}
	}

	return freqList, bounds
}
