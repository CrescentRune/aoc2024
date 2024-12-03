package main

import (
    "fmt"
)

func main() {

    full := []byte{0, 0, 0, 0, 0, 0, 0}
    part := []byte{1, 1, 1}

    newFull := append([]byte{}, full...)
    copy(newFull[2:], part)
    fmt.Println("newFull:      ", newFull)
    fmt.Println("original full:", full)

}
