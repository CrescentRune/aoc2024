package main

import (
	"bufio"
	"os"
    "log"
    "strings"
    "strconv"
)

func main() {
    day2();
}

func day2() {
    file, err := os.Open("./day2.txt")
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)


   // p1 := part1(scanner)
   // log.Printf("Part 1 solution: %d\n", p1)

    scanner = bufio.NewScanner(file)
    p2 := part2(scanner)

    log.Printf("Part 2 solution: %d\n", p2)
    //part2();
}

func part1(scanner *bufio.Scanner) int {
    
    badLines := 0
    lineNum := 0

    for scanner.Scan() {
        report := strings.Split(scanner.Text(), " ")
        prevVal := 0
        direction := 0
        differenceList := []int {}
        isBad := false
        for _, level := range report {
            if (isBad) {
                continue
            }
            currVal, err := strconv.Atoi(level)
            if err != nil {
                log.Fatal(err)
            }

            if (prevVal != 0) {
                difference := currVal - prevVal
                if (direction == 0) {
                    direction = difference
                }

                differenceList = append(differenceList, difference)
                if (direction < 0) {
                    difference *= -1
                }

                if (difference < 1 || difference > 3) {
                    badLines++
                    isBad = true;
                }
            }


            prevVal = currVal;
        }

        if (isBad) {
            log.Printf("== Line is bad: Report: \"%s\", %d, %v", report, direction, differenceList)
        } else {
            log.Printf("== Line is good: Report: \"%s\", %d, %v", report, direction, differenceList)
        }

        lineNum++
    }
    return lineNum - badLines
}

func part2(scanner *bufio.Scanner) int {
    safeReports := 0
    for scanner.Scan() {
        report, err := parseReport(scanner.Text())
        if err != nil {
            log.Fatal("Report contained invalid values")
        }

        
        reportSafe := isReportSafe(report, true, -1)

        if reportSafe {
            safeReports++    
        }
    }

    return safeReports
}

func parseReport(report string) ([]int, error) {
    levels := strings.Split(report, " ")
    nums := []int {}
    for _, level := range levels {
        num, err := strconv.Atoi(level)
        if err != nil {
            return nil, err
        }
        nums = append(nums, num)
    }
    return nums, nil
}

func isReportSafe(report []int, allowSkip bool, skip int) bool {
    prevVal := 0
    currVal := 0
    direction := 0

    for i, level := range report {
        if allowSkip && skip == i {
            continue
        }
        currVal = level

        if (prevVal != 0) {
            difference := currVal - prevVal

            if (direction == 0) {
                direction = difference
            }

            if (direction < 0) {
                difference *= -1
            }
            
            if (difference < 1 || difference > 3) {
                if (allowSkip && skip == -1) {
                    result := isReportSafe(report, true, i-1) || isReportSafe(report, true, i)

                    if !result && i == 2 {
                        result = result || isReportSafe(report, true, i-2)
                        if result {
                            log.Printf("Report: hanks rule violated (%d, %d) report: %v", i-1, i, report)
                        }
                    }
                    return result
                }
                return false
            }
            
        }

        prevVal = currVal
    }
    
    return true
}
