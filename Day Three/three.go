package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    file, err := ioutil.ReadFile("input")
    if err != nil {
        fmt.Println("File reading error:", err)
        return
    }
    input := string(file)

    inputSlice := strings.Split(strings.TrimSpace(input), "\n")

    var forest [][]rune
    for _, line := range inputSlice {
        forest = append(forest, []rune(line))
    }

    collisions := findCollisions(forest, 3, 1)

    fmt.Println("Tree collisions (p1):", collisions)

    offsetPairs := [][]int{
        {3, 1},
        {1, 1},
        {5, 1},
        {7, 1},
        {1, 2},
    }
    treeProduct := 1

    for _, pair := range offsetPairs {
        encounteredTrees := findCollisions(forest, pair[0], pair[1])
        treeProduct *= encounteredTrees
    }

    fmt.Println("Tree colisions (p2):", treeProduct)
}

var tree_char = []rune("#")[0]

func findCollisions(forest [][]rune, xOffset, yOffset int) int {
    var encounteredTrees int
    var xPointer int
    var yPointer int

    for yPointer < len(forest) {
        row := forest[yPointer]
        targetIndex := xPointer % len(row)
        if row[targetIndex] == tree_char {
            encounteredTrees += 1
        }
        xPointer += xOffset
        yPointer += yOffset
    }

    return encounteredTrees
}
