package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

type FreeListEntry struct {
    start int
    space int
}

type FileListEntry struct {
    fileId int
    index int
    width int
}

func main() {
	input, err := os.ReadFile("input.txt")

	if err != nil {
		panic("Ahh! File bad!")
	}

	diskmap := convertDiskmap(input)

	defrag := generateDefrag(diskmap)

    p2Defrag := make([]int, len(defrag))
    copy(p2Defrag, defrag)

    fmt.Printf("Part 1 solution: %d\n", findSumOfDefrag(p1(defrag)))
    fmt.Printf("Part 2 solution: %d\n", findSumOfDefrag(p2(p2Defrag)))
}

func p1(defrag []int) []int {

    emptyCursor := findNextFileSlot(defrag)
	fileCursor := findLastFileSlot(defrag)

    for ; emptyCursor > 0 && fileCursor > 0 && emptyCursor < fileCursor; {
        defrag[emptyCursor] = defrag[fileCursor]
        defrag[fileCursor] = -1

        emptyCursor = findNextFileSlot(defrag)
	    fileCursor = findLastFileSlot(defrag)
    }
    return defrag
}

func p2(defrag []int) []int {
    freeList, fileList := generateLists(defrag)
    
    //printFileList(fileList)
    //printFreeList(freeList)

    defrag = defragLists(defrag, fileList, freeList)
    

    return defrag
}

func moveFile(defrag []int, file FileListEntry, freeBlock FreeListEntry) []int {
    for i := freeBlock.start; i < freeBlock.start + file.width; i++ {
        defrag[i] = file.fileId
    }

    for i := file.index; i < file.index + file.width; i++ {
        defrag[i] = -1
    }

    //fmt.Printf("File moved: %v\n", defrag)

    return defrag
}

func defragLists(defrag []int, fileList []FileListEntry, freeList []FreeListEntry) []int {
    for _, file := range fileList {
        for i, freeBlock := range freeList {
            extraSpace := freeBlock.space - file.width
            if extraSpace >= 0 {
                //fmt.Printf("Moving file! FileId %d, freeblock index %d\n", file.fileId, freeBlock.start)
                defrag = moveFile(defrag, file, freeBlock)
                if extraSpace > 0 {
                    freeList[i] = FreeListEntry{freeBlock.start+file.width, extraSpace}
                } else {
                    freeList = slices.Delete(freeList, i, i+1)
                }
                break
            }
        }
    }

    return defrag
}

func printFileList(fileList []FileListEntry) {
    fmt.Printf("File list:\n")
    for _, entry := range fileList {
        fmt.Printf("{id: %d, index: %d, size: %d}\n", entry.fileId, entry.index, entry.width)
    }
}


func printFreeList(freeList []FreeListEntry) {
    fmt.Printf("Free list:\n")
    for _, entry := range freeList {
        fmt.Printf("{index: %d, size: %d}\n", entry.start, entry.space)
    }
}

func generateLists(defrag []int) ([]FreeListEntry, []FileListEntry) {
    currId := -3

    var fileList []FileListEntry

    currFileWidth := 1
    for i := len(defrag) - 1; i > 0; i-- {
        isFileBlock := defrag[i] != -1
        isRunningFile := defrag[i] == currId
        
        if (isFileBlock && isRunningFile) {
            currFileWidth++
        }

        if currId > -1 && ((isFileBlock && !isRunningFile) || (!isFileBlock)) {
            fileList = append(fileList, FileListEntry{currId, i+1, currFileWidth})
            currFileWidth = 1
        }

        currId = defrag[i]
    }

    fileList = append(fileList, FileListEntry{0, 0, currFileWidth+1})

    var freeList []FreeListEntry
    freeWidth := 0
    peekBack := false
    for i, val := range defrag {
        isCurrFree := val == -1
        if isCurrFree {
            freeWidth++
        } else if peekBack {
            freeList = append(freeList, FreeListEntry{i-freeWidth, freeWidth})
            freeWidth = 0
        }

        peekBack = isCurrFree
    }

    return freeList, fileList
}

func findSumOfDefrag(defrag []int) int64 {
    var sum int64 = 0
    for i, val := range defrag {
        if val < 0 {
            continue 
        }

        sum += int64(i * val)
    }

    return sum
}

func generateDefrag(diskmap []int) []int {
	sum := 0
	for _, block := range diskmap {
		sum += block
	}

	defrag := make([]int, sum)

    for i := range defrag {
        defrag[i] = -1
    }


    
    diskCursor := 0

    for defragIndex := 0; defragIndex < len(defrag); {
        for range diskmap[diskCursor] {
            if diskCursor % 2 == 0 {
                defrag[defragIndex] = getIdFromIndex(diskCursor)       
            }
            defragIndex++
        }
        diskCursor++
    }


	return defrag
}

func findNextFileSlot(defrag []int) int {
    for i := range defrag {
        if defrag[i]  == -1 { return i }
    }
    return -1
}

func findLastFileSlot(defrag []int) int {
    for i := len(defrag) - 1; i > 0; i-- {
        if defrag[i] != -1 { return i }
    }

    return -1
}


func getIdFromIndex(cursor int) int {
	return ((cursor) / 2)
}

func printDefrag(defrag []int) {
    printable := ""
    for _, val := range defrag {
        if val < 0 {
            printable += "."
        } else {
            printable += strconv.Itoa(val)
        }
    }

    fmt.Printf("[%s]\n", printable)
}

func convertDiskmap(compMap []byte) []int {
	diskMap := make([]int, len(compMap))
	for i, char := range compMap {
        if char == '\n' {
            break
        }
		diskMap[i] = int(char - '0')
	}

	return diskMap
}
