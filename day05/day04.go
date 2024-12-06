package main

import (
	"fmt"
	"os"
	"regexp"
)

type DepMap map[int]int 

func main() {
    input, err := os.ReadFile("input.txt")

    if err != nil {
        panic("input could not be read")
    }

    dependencies, workList := parseInput(input)

    safeReports := findSafeUpdate(dependencies, workList)

    fmt.Printf("Part 1 solution: %d\n", safeReports)
}

func parseInput(input []byte) (DepMap, [][]byte) {
    //depRe := regexp.MustCompile(`(\d+)\|(\d+)`)

    return nil, nil

}

func findSafeUpdate(dependencies DepMap, updates [][]byte) int {
    return 0
}
