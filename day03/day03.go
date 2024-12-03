package main

import (
	"fmt"
	"os"
	"regexp"
    "strconv"
)

func main() {
    day03()
}

func day03() {

    input, err := os.ReadFile("day03.txt")
    if err != nil {
        panic(err) 
    }
    
    p1 := part1(input)
    fmt.Printf("Part1 Result: %v\n", p1)

    p2 := part2(input)
    fmt.Printf("Part 2 Result: %v\n", p2)
}

func part1(input []byte) int {
    sum := 0
    re := regexp.MustCompile(`mul\(([\d]+),([\d]+)\)`)
    result := re.FindAllSubmatch(input, -1)
    for _, item := range result {
        a, err1 := strconv.Atoi(string(item[1]))
        b, err2 := strconv.Atoi(string(item[2]))

        if err1 != nil || err2 != nil {
            panic(err1)
        }

        sum += a * b 
        //fmt.Printf("%d * %d = %d (%d)\n", a, b, a * b, sum)
    }
    return sum
}

func part2(input []byte) int {
    sum := 0
    re := regexp.MustCompile(`(?:don't\(\)(?:.|\n)+?do\(\))|(?:mul\(([\d]+),([\d]+)\))`)
    result := re.FindAllSubmatch(input, -1)
    for _, item := range result {
        fmt.Printf("%s\n", item)
        if string(item[0][0:3]) == "mul" {
            a, err1 := strconv.Atoi(string(item[1]))
            b, err2 := strconv.Atoi(string(item[2]))

            if err1 != nil || err2 != nil {
                panic(err1)
            }

            sum += a * b 
            fmt.Printf("%d * %d = %d (%d)\n", a, b, a * b, sum)
        }
    }
    return sum
}

