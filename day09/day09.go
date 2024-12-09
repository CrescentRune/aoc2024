package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")

	if err != nil {
		panic("Ahh! File bad!")
	}

	diskmap := convertDiskmap(input)

	defrag := generateDefrag(diskmap)

	fmt.Printf("%v\n", defrag)
}

func generateDefrag(diskmap []int) []int {
	sum := 0
	for _, block := range diskmap {
		sum += block
	}

	defrag := make([]int, sum)

	fileCursor := len(diskmap) - 1
	filesRemaining := diskmap[fileCursor]

	diskCursor := 0

	for defragIdx := 0; defragIdx < sum; {
		writeCount := 0
		if diskmap[diskCursor] <= 0 {
			diskCursor++
			continue
		}

		for defragOffset := range diskmap[diskCursor] {
			if diskCursor%2 == 0 {
				defrag[defragIdx+defragOffset] = getIdFromIndex(diskCursor)
			} else {
				defrag[defragIdx+defragOffset] = getIdFromIndex(fileCursor)
				fileCursor, filesRemaining = getNextFileInfo(diskmap, fileCursor, filesRemaining)
			}
			writeCount += 1
		}

		diskCursor += 1
		defragIdx += writeCount
	}

	return defrag
}

func getNextFileInfo(diskmap []int, cursor int, leftInBlock int) (int, int) {
	if leftInBlock >= 1 {
		return cursor, leftInBlock - 1
	}

	for {
		cursor -= 2
		leftInBlock = diskmap[cursor]
		if leftInBlock > 0 {
			return cursor, leftInBlock - 1
		}
	}
}

func getIdFromIndex(cursor int) int {
	return (cursor + 1) / 2
}

func convertDiskmap(compMap []byte) []int {
	diskMap := make([]int, len(compMap))
	for i, char := range compMap {
		diskMap[i] = int(char - '0')
	}

	return diskMap
}
