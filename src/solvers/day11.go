package solvers

import (
	"fmt"
	"helpers"
)

var _xdim int
var _ydim int

func SolveDay11() {
	lines, err := helpers.ReadInputFile()

	if err != nil {
		fmt.Println("ReadInputFile returned error(s). Exiting.")
		fmt.Println(lines)
		return
	}

	_ydim = len(lines)
	_xdim = len(lines[0])
	initialMap := make([][]byte, _ydim)
	for i, line := range lines {
		initialMap[i] = make([]byte, _xdim)
		for j, char := range line {
			initialMap[i][j] = byte(char)
		}
	}

	seatMap := findStableSeatArrangement(initialMap)
	fmt.Printf("Number of occupied seats: %d\n", countOccupiedSeats(seatMap))
	seatMapPart2 := findStableSeatArrangementPart2(initialMap)
	fmt.Printf("Number of occupied seats for part 2: %d\n", countOccupiedSeats(seatMapPart2))
}

func findStableSeatArrangement(startingMap [][]byte) [][]byte{
	currentMap := copyMap(startingMap)
	iterations := 0
	for {
		iterations += 1
		changes := seatPeople(currentMap)
		if changes == false {
			fmt.Printf("Map unchanged at iteration #%d\n", iterations)
			return currentMap
		}
		if iterations > 100 {
			fmt.Printf("Unable to reach stable state after %d iterations. Exiting....\n", iterations)
			return currentMap
		}
	}
}

func copyMap(src [][]byte) [][]byte {
	srcYLen := len(src)
	newMap := make([][]byte, srcYLen)
	for i := 0; i < srcYLen; i++ {
		srcXLen := len(src[i])
		newMap[i] = make([]byte, srcXLen)
		for j := 0; j < srcXLen; j++ {
			newMap[i][j] = src[i][j]
		}
	}
	return newMap
}

func seatPeople(startingMap [][]byte) bool {
	newMap := make([][]byte, _ydim)
	for i := 0; i < _ydim; i++ {
		newMap[i] = make([]byte, _xdim)
		for j := 0; j < _xdim; j++ {
			if isFloor(startingMap, i, j) {
				newMap[i][j] = startingMap[i][j]
				continue
			}
			isOccupied := isOccupied(startingMap, i, j)
			adjacentOccupied := adjacentOccupiedSeats(startingMap, i,j)
			if !isOccupied && adjacentOccupied == 0 {
				newMap[i][j] = byte('#')
			} else if isOccupied && adjacentOccupied >= 4 {
				newMap[i][j] = byte('L')
			} else {
				newMap[i][j] = startingMap[i][j]
			}
		}
	}

	if mapChanged(startingMap, newMap) {
		overwriteMap(startingMap, newMap)
		return true
	}
	return false
}

func adjacentOccupiedSeats(seatMap [][]byte, y int, x int) int {
	count := 0
	yStart := y - 1
	if yStart < 0 {
		yStart = 0
	}
	yEnd := y + 2
	if yEnd > _ydim {
		yEnd = _ydim
	}
	xStart := x - 1
	if xStart < 0 {
		xStart = 0
	}
	xEnd := x + 2
	if xEnd > _xdim {
		xEnd = _xdim
	}
	for i := yStart; i < yEnd; i++ {
		for j := xStart; j < xEnd; j++ {
			if i == y && j == x {
				continue
			}
			if seatMap[i][j] == byte('#') {
				count += 1
			}
		}
	}
	return count
}

func isOccupied(seatMap [][]byte, y int, x int) bool {
	return seatMap[y][x] == byte('#')
}

func isFloor(seatMap [][]byte,y int, x int) bool {
	return seatMap[y][x] == byte('.')
}

func countOccupiedSeats(seatMap [][]byte,) int {
	count := 0
	for i := 0; i < _ydim; i++ {
		for j := 0; j < _xdim; j++ {
			if seatMap[i][j] == byte('#') {
				count += 1
			}
		}
	}
	return count
}

func mapChanged(initialMap [][]byte, newMap [][]byte) bool {
	for i := 0; i < _ydim; i++ {
		for j := 0; j < _xdim; j++ {
			if initialMap[i][j] != newMap[i][j] {
				return true
			}
		}
	}
	return false
}

func overwriteMap(dest [][]byte, src [][]byte) {
	for i := 0; i < _ydim; i++ {
		for j := 0; j < _xdim; j++ {
			dest[i][j] = src[i][j]
		}
	}
}

func findStableSeatArrangementPart2(startingMap [][]byte) [][]byte{
	currentMap := copyMap(startingMap)
	iterations := 0
	for {
		iterations += 1
		changes := seatPeoplePart2(currentMap)
		if changes == false {
			fmt.Printf("Part 2: map unchanged at iteration #%d\n", iterations)
			return currentMap
		}
		if iterations > 100 {
			fmt.Printf("Part 2: Unable to reach stable state after %d iterations. Exiting....\n", iterations)
			return currentMap
		}
	}
}

func seatPeoplePart2(startingMap [][]byte) bool {
	newMap := make([][]byte, _ydim)
	for i := 0; i < _ydim; i++ {
		newMap[i] = make([]byte, _xdim)
		for j := 0; j < _xdim; j++ {
			if isFloor(startingMap, i, j) {
				newMap[i][j] = startingMap[i][j]
				continue
			}
			isOccupied := isOccupied(startingMap, i, j)
			occupiedSeatsIn8dirs := occupiedSeatsIn8dirs(startingMap, i, j)
			if !isOccupied && occupiedSeatsIn8dirs == 0 {
				newMap[i][j] = byte('#')
			} else if isOccupied && occupiedSeatsIn8dirs >= 5 {
				newMap[i][j] = byte('L')
			} else {
				newMap[i][j] = startingMap[i][j]
			}
		}
	}

	if mapChanged(startingMap, newMap) {
		overwriteMap(startingMap, newMap)
		return true
	}
	return false
}

func occupiedSeatsIn8dirs(seatMap [][]byte, y int, x int) int {
	occupiedSeats := 0
	for rise := -1; rise < 2; rise++ {
		for run := -1; run < 2; run++ {
			if rise == 0 && run == 0 {
				continue
			}
			if seeOccupiedSeatInDirection(seatMap, y, x, rise, run) {
				occupiedSeats += 1
			}
		}
	}
	return occupiedSeats
}

func seeOccupiedSeatInDirection(seatMap [][]byte, ypos int, xpos int, ystep int, xstep int) bool {
	currentY := ypos
	currentX := xpos
	for {
		currentY += ystep
		currentX += xstep
		if inBounds(currentY, currentX) {
			if isFloor(seatMap, currentY, currentX) {
				continue
 			} else {
 				return isOccupied(seatMap, currentY, currentX)
			}
		} else {
			return false
		}
	}
}

func inBounds(y int, x int) bool {
	return y >= 0 &&  y < _ydim && x >= 0 && x < _xdim
}