package day8

import (
    "fmt"
    "os"
    "path/filepath"
    "aoc2022/input"
    "strconv"
)

const dayName string = "day8"

func Puzzle1() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    inputSlice := input.GetInputSlice(filename)

    treeMap := make([][]int, 0)
    visible := 0

    for i, line := range inputSlice {
        // Populate treeMap with integers
        treeMap = append(treeMap, make([]int, 0))
        for j := 0; j < len(line); j++ {
            num, err := strconv.Atoi(line[j:j+1]) 
            if err != nil {
                panic(err)
            }

            treeMap[i] = append(treeMap[i], num)
        }
    }

    for i, treeRow := range treeMap {
        for j, height := range treeRow {
            up, down, left, right := true, true, true, true

            // Look up
            dirStep := 1
            for {
                if i - dirStep < 0 {
                    // Visible up
                    break
                } else if treeMap[i-dirStep][j] >= height {
                    // Not visible up
                    up = false
                    break
                } else {
                    // Still potentially visible, increment dirStep
                    dirStep++
                }  
            }

            // Look down
            dirStep = 1
            for {
                if i + dirStep >= len(treeMap) {
                    // Visible down
                    break
                } else if treeMap[i+dirStep][j] >= height {
                    // Not visible down
                    down = false
                    break
                } else {
                    // Still potentially visible, increment dirStep
                    dirStep++
                }  
            }

            // Look left
            dirStep = 1
            for {
                if j - dirStep < 0 {
                    // Visible left
                    break
                } else if treeMap[i][j-dirStep] >= height {
                    // Not visible left
                    left = false
                    break
                } else {
                    // Still potentially visible, increment dirStep
                    dirStep++
                }  
            }

            // Look right
            dirStep = 1
            for {
                if j + dirStep >= len(treeRow) {
                    // Visible left
                    break
                } else if treeMap[i][j+dirStep] >= height {
                    // Not visible right
                    right = false
                    break
                } else {
                    // Still potentially visible, increment dirStep
                    dirStep++
                }  
            }

            if up || down || left || right {
                visible += 1
            }

            fmt.Printf("\ni: %v\nj: %v\n", i, j)
        }
    }

    fmt.Printf("Output: %v\n", visible)
}

func Puzzle2() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    inputSlice := input.GetInputSlice(filename)

    treeMap := make([][]int, 0)
    score := -1

    for i, line := range inputSlice {
        // Populate treeMap with integers
        treeMap = append(treeMap, make([]int, 0))
        for j := 0; j < len(line); j++ {
            num, err := strconv.Atoi(line[j:j+1]) 
            if err != nil {
                panic(err)
            }

            treeMap[i] = append(treeMap[i], num)
        }
    }

    for i, treeRow := range treeMap {
        for j, height := range treeRow {
            up, down, left, right := 0, 0, 0, 0

            // Look up
            dirStep := 1
            for {
                if i - dirStep < 0 {
                    up = dirStep - 1
                    break
                } else if treeMap[i-dirStep][j] >= height {
                    up = dirStep
                    break
                } else {
                    dirStep++
                }  
            }

            // Look down
            dirStep = 1
            for {
                if i + dirStep >= len(treeMap) {
                    down = dirStep-1
                    break
                } else if treeMap[i+dirStep][j] >= height {
                    down = dirStep
                    break
                } else {
                    dirStep++
                }  
            }

            // Look left
            dirStep = 1
            for {
                if j - dirStep < 0 {
                    left = dirStep - 1
                    break
                } else if treeMap[i][j-dirStep] >= height {
                    left = dirStep
                    break
                } else {
                    dirStep++
                }  
            }

            // Look right
            dirStep = 1
            for {
                if j + dirStep >= len(treeRow) {
                    right = dirStep - 1
                    break
                } else if treeMap[i][j+dirStep] >= height {
                    right = dirStep
                    break
                } else {
                    dirStep++
                }  
            }

            product := up * down * left * right
            if product > score {
                score = product
            }

            fmt.Printf("\ni: %v\nj: %v\n", i, j)
        }
    }

    fmt.Printf("Output: %v\n", score)
}
