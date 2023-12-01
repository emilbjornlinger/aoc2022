package day9

import (
    "fmt"
    "os"
    "path/filepath"
    "aoc2022/input"
)

const dayName string = "day9"

func Puzzle1() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    inputSlice := input.GetInputSlice(filename)

    // Implementation
    for _, line := range inputSlice {
        fmt.Println(line)
    }

    // Have a struct for position containing x and y

    // One pos for H and one for T

    // Move head according to input, start both at 0, 0

    // Perform motion of T according to logic -> one updateTail() function

    // Convert position into strings and concatenate (-1, 5 -> "-15") and 
    // put into a positionMap

    // Count the number of entries when the simulation is done

    output := "Hello from " + dayName
    fmt.Printf("Output: %v\n", output)
}

func Puzzle2() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, dayName, "input.txt")
    inputSlice := input.GetInputSlice(filename)

    // Implementation
    for _, line := range inputSlice {
        fmt.Println(line)
    }

    output := "Hello from " + dayName
    fmt.Printf("Output: %v\n", output)
}
