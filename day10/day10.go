package main

import (
	"fmt"
	"os"
)

type Point struct {
    row int
    col int
}

var directions = []Point{
    //{-1,-1},
    {-1,0},
    //{-1,1},
    {0,1},
    //{1,1},
    {1,0},
    //{1,-1},
    {0,-1},
}

func main() {
    input, err := os.ReadFile("input.txt")

    if err != nil {
        panic("Oh no file bad!")
    }

    grid := parseInput(input)

    p1 := computeTrailheadValues(grid)

    fmt.Printf("Part 1 solution: %d\n", p1)
}

func computeTrailheadValues(grid [][]int) int {
    sum := 0

    for i, row := range grid {
        for j, val := range row {
            if val == 0 {
                trailheadSet := make(map[string]bool)
                walk(grid, Point{i,j}, trailheadSet)
                fmt.Printf("(%d,%d): Set (len:%d) %v\n", i, j, len(trailheadSet), trailheadSet)
                sum += len(trailheadSet)
            }
        }
    }
    return sum
}

func walk(grid [][]int, pos Point, set map[string]bool) {
    val := grid[pos.row][pos.col]
    if val == 9 {
        strRep := fmt.Sprintf("%d,%d",pos.row, pos.col)
        set[strRep] = true
    }

    cands := getNextSteps(grid, pos)

    for _, cand := range cands {
        walk(grid, cand, set)
    }
}

func getNextSteps(grid [][]int, pos Point) []Point {
    nextSteps := []Point{}

    val := grid[pos.row][pos.col]

    for _, dir := range directions {
        cand := addPoints(pos, dir)
        if isInGrid(grid, cand) && grid[cand.row][cand.col] == val + 1 {
            nextSteps = append(nextSteps, cand)
        }
    }

    return nextSteps
}

func addPoints(a Point, b Point) Point {
    return Point{a.row + b.row, a.col + b.col}
}

func isInGrid(grid [][]int, point Point) bool {
    return point.row >= 0 && point.col >= 0 && point.row < len(grid) && point.col < len(grid[0]) 
}

func parseInput(input []byte) [][]int {
    grid := [][]int{{}}
    currRow := 0
    for i, char := range input {
        if i + 1 == len(input) && char == '\n' {
            break
        }
        if char == '\n' {
            currRow++
            grid = append(grid, []int{})
        } else {
            grid[currRow] = append(grid[currRow], int(char - '0'))
        }
    }

    return grid
}
