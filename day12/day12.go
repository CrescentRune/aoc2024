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

var surroundGrid = [][][]int{
	{{-1, -1}, {-1, 0}, {-1, 1}},
	{{0, -1}, {0, 0}, {0, 1}},
	{{1, -1}, {1, 0}, {1, 1}},
}

var diagDirections = [][]int{
	{-1, -1},
	{-1, 1},
	{1, 1},
	{1, -1},
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic("Ufda")
	}

	grid := parseInput(input)

	p1 := computeFenceCost(grid)

	fmt.Printf("Part 1 solution: %d\n", p1)

}

func isInGrid(grid [][]byte, i int, j int) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}

func getStrRep(i int, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}

func computeFenceCost(grid [][]byte) int {
	visited := make(map[string]bool)
	cost := 0

	for i, row := range grid {
		for j := range row {
			strRep := getStrRep(i, j)
			if !visited[strRep] {
				cost += computeFenceAround(visited, grid, i, j)
			}
		}
	}

	return cost
}

func computeFenceAround(visited map[string]bool, grid [][]byte, i int, j int) int {
	queue := [][]int{{i, j}}
	fences := 0
	uniq := 0
	for itemI := 0; itemI < len(queue); itemI++ {
		itemFences := 0
		item := queue[itemI]
		if visited[getStrRep(item[0], item[1])] {
			continue
		}
		char := grid[item[0]][item[1]]

		uniq++
		for _, dir := range directions {
			cI, cJ := item[0]+dir[0], item[1]+dir[1]
			if !isInGrid(grid, cI, cJ) {
				itemFences++
				continue
			}

			matches := char == grid[cI][cJ]
			strRep := getStrRep(cI, cJ)
			if !visited[strRep] && matches {
				queue = append(queue, []int{cI, cJ})
				//fmt.Printf("Enqueuing (%s) from (%s)\n", strRep, getStrRep(item[0], item[1]))
			} else if !matches {
				itemFences++
			}
			visited[getStrRep(item[0], item[1])] = true
		}
		//fences += itemFences
		corners := findCorners(grid, item[0], item[1])
		fences += corners
		//fmt.Printf("Item (%c) at (%s) requires %d fences, plot has %d corners\n", grid[item[0]][item[1]], getStrRep(item[0], item[1]), itemFences, corners)
	}

	//fmt.Printf("Total fencing cost for %c: %d*%d = %d\n", grid[i][j], fences, uniq, fences*uniq)

	return fences * uniq
}

func createMatchGrid(grid [][]byte, row int, col int) [][]bool {
	char := grid[row][col]
	matches := [][]bool{}
	for i, gridRow := range surroundGrid {
		//fmt.Printf("%d:", i)
		matches = append(matches, []bool{})
		for _, tup := range gridRow {
			matches[i] = append(matches[i], isInGrid(grid, row+tup[0], col+tup[1]) && char == grid[row+tup[0]][col+tup[1]])
		}
		//fmt.Printf("%d\n", len(matches[i]))
	}

	//fmt.Printf("Match grid: %v\n", matches)

	return matches
}

func findCorners(grid [][]byte, row int, col int) int {
	matchGrid := createMatchGrid(grid, row, col)

	//fmt.Printf("(%d,%d) match grid:\n%v\n", row, col, matchGrid)

	adjMatches := 0
	for _, dir := range directions {
		if matchGrid[1+dir[0]][1+dir[1]] {
			adjMatches++
		}
	}

	diagDiffs := 0
	for _, dir := range diagDirections {
		if !matchGrid[1+dir[0]][1+dir[1]] {
			diagDiffs++
		}
	}

	if adjMatches == 0 {
		return 4
	} else if adjMatches == 1 {
		return 2
	} else if adjMatches == 4 {
		return diagDiffs
	}

	horzCont := matchGrid[1][0] && matchGrid[1][2]
	vertCont := matchGrid[0][1] && matchGrid[2][1]

	if horzCont && vertCont {
		return diagDiffs
	} else if horzCont && !vertCont {
		//Check if either above or below match, if so, return number of non-matches left or right of that
		count := 0
		if matchGrid[0][1] {
			//Above
			if !matchGrid[0][0] {
				count++
			}
			if !matchGrid[0][2] {
				count++
			}
		} else if matchGrid[2][1] {
			//Below
			if !matchGrid[2][0] {
				count++
			}
			if !matchGrid[2][2] {
				count++
			}
		}
		return count

	} else if vertCont && !horzCont {
		//Check if either left or right match, if so, return number of non-matches above or below that
		count := 0
		if matchGrid[1][0] {
			//Left
			if !matchGrid[0][0] {
				count++
			}
			if !matchGrid[2][0] {
				count++
			}
		} else if matchGrid[1][2] {
			//Below
			if !matchGrid[2][2] {
				count++
			}
			if !matchGrid[0][2] {
				count++
			}
		}
		return count
	} else {
		//There IS a joint here, return 1
		if (matchGrid[0][1] && matchGrid[1][2] && !matchGrid[0][2]) ||
			(matchGrid[0][1] && matchGrid[1][0] && !matchGrid[0][0]) ||
			(matchGrid[2][1] && matchGrid[1][0] && !matchGrid[2][0]) ||
			(matchGrid[2][1] && matchGrid[1][2] && !matchGrid[2][2]) {
			return 2
		}
		return 1
	}

}

func parseInput(input []byte) [][]byte {
	inputGrid := bytes.Split(input, []byte("\n"))
	return inputGrid
}
