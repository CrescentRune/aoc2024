package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
    input, err := os.ReadFile("day04.txt")

    if err != nil {
        panic("Oh no! Input bad!")
    }

    inputGrid := parseInput(input)

    fmt.Printf("==Day 04==\n")

    p1 := part1(inputGrid)
    fmt.Printf("Part 1 Solution: %d\n", p1)

    p2 := part2(inputGrid)
    fmt.Printf("Part 2 Solution: %d\n", p2)
}

func parseInput(input []byte) [][]byte {
    inputGrid := bytes.Split(input, []byte("\n"))
    return inputGrid[:len(inputGrid)-1]
}

func drillDirection(inputGrid [][]byte, i int, iDirection int, j int, jDirection int) int {
    if (iDirection < 0 && i < 3) || 
       (jDirection < 0 && j < 3) || 
       (jDirection > 0 && j+3 >= len(inputGrid)) || 
       (iDirection > 0 && i+3 >= len(inputGrid[0])) {
        return 0
    }

    for offset, char := range []byte("MAS") {
        checkI := i + (offset + 1) * iDirection
        checkJ := j + (offset + 1) * jDirection

        if checkI >= 140 || checkJ >= 140 {
            fmt.Printf("Something went wrong at input (%d, %d), dir: (%d, %d)\n", i, j, iDirection, jDirection)
        }

        if char != inputGrid[checkI][checkJ] {
            return 0
        }
    }

    return 1
}

func findXMAS(inputGrid [][]byte, i int, j int) int {
    directions := [][]int{{-1,-1}, {0,-1}, {1,-1}, {1,0}, {1,1}, {0,1}, {-1,1}, {-1,0}};
    xmasCount := 0
    for _, dir := range directions {
        xmasCount += drillDirection(inputGrid, i, dir[0], j, dir[1]) 
    }
    
    return xmasCount
}

func part1(inputGrid [][]byte) int {
    
    sum := 0

    for i, row := range inputGrid {
        for j, char := range row {
            if string(char) == "X" {
                sum += findXMAS(inputGrid, i, j)      
            }
        }
    }

    return sum
}

func isXMas(inputGrid [][]byte, i int, j int) bool {
    if i < 1 || i+1 >= len(inputGrid[0]) || j < 1 || j+1 >= len(inputGrid) {
        return false
    }

    UL := inputGrid[i-1][j-1]
    UR := inputGrid[i+1][j-1]
    LL := inputGrid[i-1][j+1]
    LR := inputGrid[i+1][j+1]
    

    firstArm := (UL == 'M' && LR == 'S') || (UL == 'S' && LR == 'M')
    secondArm := (LL == 'M' && UR == 'S') || (LL == 'S' && UR == 'M')

    return firstArm && secondArm
}

func part2(inputGrid [][]byte) int {
    sum := 0

    for i, row := range inputGrid {
        for j, char := range row {
            if char == 'A' && isXMas(inputGrid, i, j) {
                sum += 1 
            }
        }
    }

    return sum
}
