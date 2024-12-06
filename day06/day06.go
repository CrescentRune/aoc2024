package main

import (
	"bytes"
	"fmt"
	"os"
)

var directions = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

const (
	UP = 0
	RIGHT
	DOWN
	LEFT
)

var dirIndicator = []byte{
	0b0001,
	0b0010,
	0b0100,
	0b1000,
}

func main() {
	input, err := os.ReadFile("sample.txt")
	if err != nil {
		panic("File is bad. Womp womp")
	}

	inputGrid := parseInput(input)
	uniqueSteps, potentialLoops := calculateValues(inputGrid)

	fmt.Printf("Part 1 solution: %d\n", uniqueSteps)
	fmt.Printf("Part 2 solution: %d\n", potentialLoops)
}

func calculateValues(grid [][]byte) (int, int) {
	i, j := findStart(grid)
	return traverseMaze(grid, -1, 0, i, j)
}

func traverseMaze(grid [][]byte, depth int, dir int, i int, j int) (int, int) {
	bytesChanged := 0
	loopsFound := 0
	for {
		point := grid[i][j]

		if point > 0x0F {
			grid[i][j] = dirIndicator[dir]
			bytesChanged++
		} else {
			if (point & dirIndicator[dir]) == dirIndicator[dir] {
				return -1, 0
			}

			grid[i][j] |= dirIndicator[dir]
		}

		dI, dJ := i+directions[dir][0], j+directions[dir][1]

		if !isOutsideGrid(grid, dI, dJ) && grid[dI][dJ] == '#' {
			dir = (dir + 1) % len(directions)
			grid[i][j] |= dirIndicator[dir]
		}

		prevI, prevJ := i, j

		i += directions[dir][0]
		j += directions[dir][1]
		if isOutsideGrid(grid, i, j) {
			return bytesChanged, loopsFound
		} else if depth < 1 {
			testGrid := dupGrid(grid)
			testGrid[prevI][prevJ] = '.'
			testGrid[i][j] = '#'
			path, _ := traverseMaze(testGrid, 1, dir, prevI, prevJ)
			if path < 0 {
				testGrid[i][j] = 'O'
				printGrid(testGrid, i, j)
				fmt.Printf("(%d, %d) would form a loop\n", i, j)
				loopsFound++
			}
		}
	}

}

func printGrid(grid [][]byte, i int, j int) {
	fmt.Printf("Hash added: %d, %d\n", i, j)
	header := "   "
	lsd := 0
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < len(grid[0]); i++ {
		header += digits[lsd]
		lsd = (lsd + 1) % 10
	}

	fmt.Printf("%s\n", header)

	for i, row := range grid {
		rowString := ""
		for _, char := range row {
			if char <= 0xf {
				hasVert := (char&dirIndicator[0]) != 0 || (char&dirIndicator[2]) != 0
				hasHorz := (char&dirIndicator[1]) != 0 || (char&dirIndicator[3]) != 0
				if hasVert && hasHorz {
					rowString += "+"
				} else if hasVert {
					rowString += "|"
				} else if hasHorz {
					rowString += "-"
				}
			} else {
				rowString += string(char)
			}
		}
		fmt.Printf("%d: %s\n", i, rowString)
	}
}

func dupGrid(grid [][]byte) [][]byte {
	newGrid := make([][]byte, len(grid))
	for i := range newGrid {
		newGrid[i] = make([]byte, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	return newGrid
}

func isOutsideGrid(grid [][]byte, i int, j int) bool {
	return i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0])
}

func wouldCauseLoop(grid [][]byte, i int, j int, dir int) bool {
	nextIndicator := dirIndicator[(dir+1)%len(dirIndicator)]
	return grid[i][j]&nextIndicator != 0
}

func findStart(grid [][]byte) (int, int) {
	for i, row := range grid {
		for j, char := range row {
			if char == '^' {
				return i, j
			}
		}
	}
	return -1, -1
}

func parseInput(input []byte) [][]byte {
	return bytes.Split(input, []byte("\n"))
}
