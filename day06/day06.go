package main

import (
	"bytes"
	"fmt"
	"os"
)

var directions = []Point{
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

type Point struct {
	row int
	col int
}

var pointsTravelled = []string{}
var pointsFound = []string{}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic("File is bad. Womp womp")
	}

	inputGrid := parseInput(input)
	uniqueSteps, potentialLoops := calculateValues(inputGrid)

	fmt.Printf("Part 1 solution: %d (actual 4819)\n", uniqueSteps)
	fmt.Printf("Part 2 solution: %d\n", potentialLoops)
}

func calculateValues(grid [][]byte) (int, int) {
	start := findStart(grid)
	return traverseMaze(grid, 0, 0, start)
}

func sumPoints(p1 Point, p2 Point) Point {
	return Point{p1.row + p2.row, p1.col + p2.col}
}

func isInsideGrid(grid [][]byte, point Point) bool {
	return point.row >= 0 && point.col >= 0 && point.row < len(grid) && point.col < len(grid[0])
}

func pointByte(grid [][]byte, point Point) byte {
	return grid[point.row][point.col]
}

func markDirTraveled(grid [][]byte, point Point, dir int) {
	curr := pointByte(grid, point)
	if isTravelIndicator(curr) {
		grid[point.row][point.col] |= dirIndicator[dir]
	} else {
		grid[point.row][point.col] = dirIndicator[dir]
	}
}

func isTravelIndicator(indicator byte) bool {
	return indicator > 0 && indicator <= 0b1111
}

func traverseMaze(grid [][]byte, depth int, dir int, pos Point) (int, int) {
	traveled, loopsFound := 0, 0
	for {
		currIndicator := pointByte(grid, pos)
		currIsTravel := isTravelIndicator(currIndicator)

		// Loop detection escape
		if depth > 0 && currIsTravel && currIndicator&dirIndicator[dir] != 0 {
			return 0, 1
		}

		if !currIsTravel {
			traveled++
		}

		markDirTraveled(grid, pos, dir)

		//Start thinking about what's next
		candPoint := sumPoints(pos, directions[dir])
		if !isInsideGrid(grid, candPoint) {
			return traveled, loopsFound
		}
		nextIndicator := pointByte(grid, candPoint)

		if nextIndicator == '#' {
			//Turn, then continue to rerun evaluation for new next
			dir = (dir + 1) % len(directions)

			continue
		}

		if depth <= 0 && nextIndicator != '#' && !isTravelIndicator(nextIndicator) {
			//Traverse and see what happens if this becomes a blockage
			gridCopy := dupGrid(grid)
			gridCopy[candPoint.row][candPoint.col] = '#'
			_, loop := traverseMaze(gridCopy, 1, (dir+1)%len(directions), pos)
			if loop != 0 {
				loopsFound++
			}
		}

		pos = sumPoints(pos, directions[dir])
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

func findStart(grid [][]byte) Point {
	for i, row := range grid {
		for j, char := range row {
			if char == '^' {
				return Point{i, j}
			}
		}
	}
	return Point{-1, -1}
}

func parseInput(input []byte) [][]byte {
	return bytes.Split(input, []byte("\n"))
}
