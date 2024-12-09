package main

import (
	//"bytes"
	"bytes"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type RowSet struct {
    goal int64
    values []int
}

func main() {
    input, err := os.ReadFile("input.txt")

    if err != nil {
        panic("File all bad!")
    }

    rows := parseInput(input) 
    
    p1, p2 := findSumOfCorrectRows(rows)

    fmt.Printf("Part 1 solution: %d\n", p1)
    fmt.Printf("Part 2 solution: %d\n", p2)
}

func findSumOfCorrectRows(rows []*RowSet) (int64, int64) {
    var sum, concatSum int64 = 0, 0
    for _, row := range(rows) {
        if canMakeGoal(row.goal, row.values, false) {
            sum += row.goal   
        } else if canMakeGoal(row.goal, row.values, true) {
            concatSum += row.goal
            //fmt.Printf("Row %d: %v could be made using concatenation!\n", row.goal, row.values)
        }
    }

    return sum, sum + concatSum
}

func catNumbers(left int64, right int64) int64 {
    rightLen := len(strconv.FormatInt(right, 10))
    
    return left * int64(math.Pow10(rightLen)) + right
}

func canMakeGoal(goal int64, ops []int, allowCat bool) bool {
    if allowCat {
        return canMakeValP2(goal, int64(ops[0]), ops[1:])
    }
    return canMakeVal(goal, int64(ops[0]), ops[1:])
}

func canMakeVal(goal int64, currentValue int64, ops []int) bool {
    if (len(ops) == 1) {
        return currentValue * int64(ops[0]) == goal || currentValue + int64(ops[0]) == goal 
    }
    return canMakeVal(goal, currentValue + int64(ops[0]), ops[1:]) ||
            canMakeVal(goal, currentValue * int64(ops[0]), ops[1:])
}

func canMakeValP2(goal int64, currentValue int64, ops []int) bool {
    //fmt.Printf("Calling can make val for goal: %d, currentValue: %d, vals remaining: %v\n", goal, currentValue, ops)
    if (len(ops) == 1) {
     //   fmt.Printf("Reached the end!\n")
        return currentValue * int64(ops[0]) == goal ||
                currentValue + int64(ops[0]) == goal ||
                catNumbers(currentValue, int64(ops[0])) == goal
    }
    return canMakeValP2(goal, currentValue + int64(ops[0]), ops[1:]) ||
            canMakeValP2(goal, currentValue * int64(ops[0]), ops[1:]) ||
            canMakeValP2(goal, catNumbers(currentValue, int64(ops[0])), ops[1:])
}


func getValuesFromRow(row [][]byte) *RowSet {
    row = row[1:]
    var rowNums []int
    var goal int64

    goal, err := strconv.ParseInt(string(row[0]), 10, 64)
    if err != nil {
        panic("Oh man!\n")
    }

    for _, val := range bytes.Split(row[1], []byte(" ")) {
        val, err := strconv.Atoi(string(val))
        if err != nil {
            panic("No way!\n")
        }
        rowNums = append(rowNums, val)
    }

    return  &RowSet{
        goal,
        rowNums,
    }
}

func parseInput(input []byte) []*RowSet {
    //rowByRow := bytes.Split(input, []byte("\n"))
    
    var output []*RowSet

    re := regexp.MustCompile(`(\d+): ((?: ?\d+)+)`)
    results := re.FindAllSubmatch(input, -1)
   
    for _, line := range results {
        output = append(output, getValuesFromRow(line))
    }

    return output
}
