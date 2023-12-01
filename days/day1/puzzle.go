package day1

import (
    "fmt"
    "os"
    "path/filepath"
    "strconv"
    "aoc2022/input"
)

const dayName string = "day1"

func Puzzle1() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    input := input.GetInputSlice(filename)

    for _, currLine := range input {
        fmt.Println(currLine)
    }

    output := "Hello from " + dayName
    fmt.Printf("Output: %v\n", output)
}

func Puzzle2() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    input := input.GetInputSlice(filename)

    for _, currLine := range input {
        fmt.Println(currLine)
    }

    output := "Hello from " + dayName
    fmt.Printf("Output: %v\n", output)
}

func handleLine(line string, nums *[]int) {
    textNums := []string{
        "one",
        "two",
        "three",
        "four",
        "five",
        "six",
        "seven",
        "eight",
        "nine",
    }

    for i := 0; i < len(line); i++ {
        if num, err := strconv.Atoi(line[i:i+1]); err == nil {
            *nums = append(*nums, num)
        } else {
            for j := 0; j < len(textNums); j++ {
                if i + len(textNums[j]) <= len(line) &&
                    line[i:i+len(textNums[j])] == textNums[j] {
                    *nums = append(*nums, j+1)
                    break
                }
            }
        } 
    }
}
