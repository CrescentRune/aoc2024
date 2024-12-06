package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type DepMap map[string][]string

func main() {
	input, err := os.ReadFile("input.txt")

	if err != nil {
		panic("input could not be read")
	}

	dependencies, workList := parseInput(input)

	safeUpdateSum, fixedUpdateSum := findSafeUpdateSum(dependencies, workList)

	fmt.Printf("Part 1 solution: %d\n", safeUpdateSum)
	fmt.Printf("Part 2 solution: %d\n", fixedUpdateSum)

}

func parseInput(input []byte) (DepMap, [][]string) {
	depRe := regexp.MustCompile(`(\d+)\|(\d+)`)

	depMap := make(DepMap)
	results := depRe.FindAllSubmatch(input, -1)

	for _, res := range results {
		lhs, rhs := string(res[1]), string(res[2])
		if depMap[rhs] == nil {
			depMap[rhs] = []string{lhs}
		} else {
			depMap[rhs] = append(depMap[rhs], lhs)
		}
	}

	updateRe := regexp.MustCompile(`(\d+,)+\d+`)
	updateResult := updateRe.FindAll(input, -1)
	updateList := [][]string{}
	for _, res := range updateResult {
		updateList = append(updateList, strings.Split(string(res), ","))
	}

	return depMap, updateList
}

func findSafeUpdateSum(dependencies DepMap, updates [][]string) (int, int) {
	safeSum := 0
	fixedSum := 0
	for _, update := range updates {
		if isUpdateSafe(dependencies, update) {
			safeSum += middleValue(update)
		} else {
			fixedUpdate := cureUpdate(dependencies, update)
			fixedSum += middleValue(fixedUpdate)
		}
	}
	return safeSum, fixedSum
}

func isUpdateSafe(dependencies DepMap, update []string) bool {
	presenceList := makePresenceList(update)

	workList := make(map[string]bool)
	for _, page := range update {
		if precons := dependencies[page]; precons != nil {
			for _, pageReq := range precons {
				if presenceList[pageReq] && !workList[pageReq] {
					return false
				}
			}
		}

		workList[page] = true
	}
	return true
}

func makePresenceList(update []string) map[string]bool {
	presencelist := make(map[string]bool)
	for _, page := range update {
		presencelist[page] = true
	}
	return presencelist
}

func cureUpdate(dependencies DepMap, update []string) []string {
	for i := 0; i < len(update); {
		page := update[i]
		pageReqs := dependencies[page]
		for j := i + 1; j < len(update); j++ {
			if slices.Contains(pageReqs, update[j]) {
				extract := update[j]
				update = slices.Delete(update, j, j+1)
				update = slices.Insert(update, i, extract)
				break
			}
		}

		if update[i] == page {
			i++
		}
	}

	return update
}

func middleValue(update []string) int {
	midPoint := len(update) / 2

	middleVal, err := strconv.Atoi(update[midPoint])
	if err != nil {
		return 0
	}
	return middleVal
}
