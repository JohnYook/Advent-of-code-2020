package solvers

import (
	"fmt"
	"helpers"
	"strconv"
)

var _xPos int
var _yPos int
var _heading int
var _shipX int
var _shipY int

func SolveDay12() {
	lines, err := helpers.ReadInputFile()

	if err != nil {
		fmt.Println("ReadInputFile returned error(s). Exiting.")
		fmt.Println(lines)
		return
	}

	_xPos = 0
	_yPos = 0
	_heading = 90 // (rise, run): E=90 0,1 W=270 0,-1 N=0 1,0 S=180 -1,0

	moveBoat(lines)

	fmt.Printf("Final position: %d, %d. Manhattan distance from origin: %d\n", _xPos, _yPos, manhattanDistance(_xPos, _yPos))

	// part 2
	_xPos = 10
	_yPos = 1
	_shipX = 0
	_shipY = 0

	moveBoatPt2(lines)
	fmt.Printf("Final position Pt2: %d, %d. Manhattan distance from origin: %d\n", _shipX, _shipY, manhattanDistance(_shipX, _shipY))
}

func moveBoat(lines []string) {
	for _, line := range lines {
		processLine(line)
	}
}

func processLine(line string) {
	action := line[0]
	value, ok := strconv.Atoi(line[1:])
	if ok != nil {
		panic(fmt.Sprintf("Error during processing line: %s\n", line))
	}
	switch action {
	case 'N':
		_yPos += value
	case 'S':
		_yPos -= value
	case 'E':
		_xPos += value
	case 'W':
		_xPos -= value
	case 'L':
		turnLeft(value)
	case 'R':
		turnRight(value)
	case 'F':
		moveForward(value)
	default:
		panic(fmt.Sprintf("Unhandled action in line: %s\n", line))
	}
}

func turnLeft(degrees int) {
	_heading = (_heading - degrees) % 360
	if _heading < 0 {
		_heading = convertToPositiveDegrees(_heading)
	}
}

func turnRight(degrees int) {
	_heading = (_heading + degrees) % 360
	if _heading < 0 {
		_heading = convertToPositiveDegrees(_heading)
	}
}

func convertToPositiveDegrees(heading int) int {
	for {
		heading += 360
		if heading >= 0 {
			return heading % 360
		}
	}
}

func moveForward(distance int) {
	switch _heading {
	case 0:
		_yPos += distance
	case 90:
		_xPos += distance
	case 180:
		_yPos -= distance
	case 270:
		_xPos -= distance
	default:
		panic(fmt.Sprintf("Unexpected boat heading: %d\n", _heading))
	}
}

func manhattanDistance(x int, y int) int {
	return helpers.Abs(x) + helpers.Abs(y)
}

func moveBoatPt2(lines []string) {
	for _, line := range lines {
		processLinePt2(line)
	}
}

func processLinePt2(line string) {
	action := line[0]
	value, ok := strconv.Atoi(line[1:])
	if ok != nil {
		panic(fmt.Sprintf("Error during processing line: %s\n", line))
	}
	switch action {
	case 'N':
		_yPos += value
	case 'S':
		_yPos -= value
	case 'E':
		_xPos += value
	case 'W':
		_xPos -= value
	case 'L':
		if value < 0 {
			turnWaypointRight(-value)
		} else {
			turnWaypointLeft(value)
		}
	case 'R':
		if value < 0 {
			turnWaypointLeft(-value)
		} else {
			turnWaypointRight(value)
		}
	case 'F':
		moveShipForward(value)
	default:
		panic(fmt.Sprintf("Unhandled action in line: %s\n", line))
	}
}

func turnWaypointLeft(value int) {
	degrees := value % 360
	initX := _xPos
	initY := _yPos
	switch degrees {
	case 0:
		return
	case 90:
		_xPos = -initY
		_yPos = initX
	case 180:
		_xPos = -initX
		_yPos = -initY
	case 270:
		_xPos = initY
		_yPos = -initX
	default:
		panic(fmt.Sprintf("Unexpected boat heading: %d\n", _heading))
	}
}

func turnWaypointRight(value int) {
	degrees := value % 360
	initX := _xPos
	initY := _yPos
	switch degrees {
	case 0:
		return
	case 90:
		_xPos = initY
		_yPos = -initX
	case 180:
		_xPos = -initX
		_yPos = -initY
	case 270:
		_xPos = -initY
		_yPos = initX
	default:
		panic(fmt.Sprintf("Unexpected boat heading: %d\n", _heading))
	}
}

func moveShipForward(units int) {
	_shipX += _xPos * units
	_shipY += _yPos * units
}