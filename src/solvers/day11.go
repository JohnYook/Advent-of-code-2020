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
	seatMap := make([][]byte, _ydim)
	for i, line := range lines {
		seatMap[i] = make([]byte, _xdim)
		for j, char := range line {
			seatMap[i][j] = byte(char)
		}
	}

	findStableSeatArrangement(seatMap)
	fmt.Printf("Number of occupied seats: %d\n", countOccupiedSeats(seatMap))
}

func findStableSeatArrangement(startingMap [][]byte) {
	iterations := 0
	for {
		iterations += 1
		changes := seatPeople(startingMap)
		if changes == false {
			fmt.Printf("Map unchanged at iteration #%d\n", iterations)
			return
		}
		if iterations > 100 {
			fmt.Printf("Unable to reach stable state after %d iterations. Exiting....\n", iterations)
			return
		}
	}
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
