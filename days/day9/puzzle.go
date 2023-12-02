package day9

import (
    "fmt"
    "os"
    "path/filepath"
    "aoc2022/input"
    "strings"
    "strconv"
    "math"
)

const dayName string = "day9"

var moveCounter int = 0

type Position struct {
    x int
    y int
}

func Puzzle1() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    inputSlice := input.GetInputSlice(filename)

    posH := Position{
        x:0,
        y:0,
    }

    posT := Position{
        x:0,
        y:0,
    }

    posMap := make(map[string]int)

    for _, line := range inputSlice {
        split := strings.Fields(line)
        dir := split[0]
        steps, err := strconv.Atoi(split[1])
        if err != nil {
            panic(err)
        }

        handleInstruction(dir, steps, &posH, &posT, &posMap)
    }

    visitedPos := len(posMap)
    fmt.Printf("Output: %v\n", visitedPos)
    fmt.Printf("moveCounter: %v\n", moveCounter)
}

func Puzzle2() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    inputSlice := input.GetInputSlice(filename)

    posH := Position{
        x:0,
        y:0,
    }

    posSlice := make([]Position, 0) 

    for i := 0; i < 9; i++ {
        posSlice = append(posSlice, Position{x:0, y:0})
    }

    posMap := make(map[string]int)

    for _, line := range inputSlice {
        split := strings.Fields(line)
        dir := split[0]
        steps, err := strconv.Atoi(split[1])
        if err != nil {
            panic(err)
        }

        handleInstructionChain(dir, steps, &posH, &posSlice, &posMap)
    }

    visitedPos := len(posMap)
    fmt.Printf("Output: %v\n", visitedPos)
    fmt.Printf("moveCounter: %v\n", moveCounter)
}

func handleInstruction(dir string, steps int, posH, posT *Position, posMap *map[string]int) {
    // Loop over number of steps
    for i := 0; i < steps; i++ {
        switch dir {
        case "U":
            posH.y = posH.y + 1
        case "D":
            posH.y = posH.y - 1
        case "L":
            posH.x = posH.x - 1
        case "R":
            posH.x = posH.x + 1
        default:
            panic("Unknown direction parsed")
        }

        // Update the position of Tail
        handleMove(posH, posT)

        // Convert position to string and check presence in map
        xString := strconv.Itoa(posT.x)
        yString := strconv.Itoa(posT.y)

        mapCopy := *posMap
        key := xString + "_" + yString
        _, exists := mapCopy[key]
        if !exists {
            mapCopy[key] = 1
            *posMap = mapCopy
        }
    }
}

func handleInstructionChain(dir string, steps int, posH *Position, posSlice *[]Position, posMap *map[string]int) {
    // Loop over number of steps
    for i := 0; i < steps; i++ {
        switch dir {
        case "U":
            posH.y = posH.y + 1
        case "D":
            posH.y = posH.y - 1
        case "L":
            posH.x = posH.x - 1
        case "R":
            posH.x = posH.x + 1
        default:
            panic("Unknown direction parsed")
        }

        // Update the position of first knot after head
        handleMove(posH, &(*posSlice)[0])

        // Handle rest of rope
        for j := 1; j < 9; j++ {
            handleMove(&(*posSlice)[j-1], &(*posSlice)[j])
        }

        // Convert position to string and check presence in map
        xString := strconv.Itoa((*posSlice)[8].x)
        yString := strconv.Itoa((*posSlice)[8].y)

        mapCopy := *posMap
        key := xString + "_" + yString
        _, exists := mapCopy[key]
        if !exists {
            mapCopy[key] = 1
            *posMap = mapCopy
        }
    }
}

func handleMove(posH, posT *Position) {

    moveCounter++
    if isValidPos(*posH, *posT) {
        return
    } else {
        for i := -1; i < 2; i++ {
            for j := -1; j < 2; j++ {
                newPos := *posT
                newPos.x = newPos.x + j
                newPos.y = newPos.y + i

                // Check if new position is valid
                if isValidPosPerp(*posH, newPos) {
                    posT.x = posT.x + j
                    posT.y = posT.y + i
                    return
                }
            }
        }
        // Else check if we can move to be in a diagonal position, the 
        // perpendicular should be exhausted first
        for i := -1; i < 2; i++ {
            for j := -1; j < 2; j++ {
                newPos := *posT
                newPos.x = newPos.x + j
                newPos.y = newPos.y + i

                // Check if new position is valid
                if isValidPos(*posH, newPos) {
                    posT.x = posT.x + j
                    posT.y = posT.y + i
                    return
                }
            }
        }
    } 

    panic("Function execution should not get here")
}

func isValidPos(posH, posT Position) bool {
    if posH.x == posT.x {
        if math.Abs(float64(posH.y - posT.y)) <= float64(1) {
            return true
        } else {
            return false
        }
    } else if posH.y == posT.y {
        if math.Abs(float64(posH.x - posT.x)) <= float64(1) {
            return true
        } else {
            return false
        }
    } else if math.Abs(float64(posH.x - posT.x)) <= float64(1) && math.Abs(float64(posH.y - posT.y)) <= float64(1) {
        return true
    } else {
        return false
    }
}

func isValidPosPerp(posH, posT Position) bool {
    if posH.x == posT.x {
        if math.Abs(float64(posH.y - posT.y)) <= float64(1) {
            return true
        } else {
            return false
        }
    } else if posH.y == posT.y {
        if math.Abs(float64(posH.x - posT.x)) <= float64(1) {
            return true
        } else {
            return false
        }
    } else {
        return false
    }
}
