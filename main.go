package main

import (
	"fmt"
	"log"
)

func main() {

	var hSize, vSize, x, y int
	var position, path string

	fmt.Println("Test Input")
	fmt.Scanln(&hSize, &vSize)
	fmt.Scanln(&x, &y, &position)
	fmt.Scanln(&path)

	//Align with the requirement. Increasing by 1.
	hSize = hSize + 1
	vSize = vSize + 1

	plane := make([][]string, hSize)
	for i := range plane {
		plane[i] = make([]string, vSize)
	}

	if x >= len(plane) || y >= len(plane[x]) {
		log.Fatal("Robot starting position is not within the rectangular plane")
	}
	plane[x][y] = "*"
	chars := []rune(path)
	for i := 0; i < len(chars); i++ {
		nextStep := string(chars[i])
		if nextStep == "R" || nextStep == "L" {
			changeDirection(nextStep, &position)
			continue
		} else if nextStep == "M" {
			moveRobot(position, plane, &x, &y)
		} else {
			log.Fatal("Command is invalid")
		}
	}
	fmt.Println("Test Output:")
	//fmt.Println(plane)
	fmt.Println(y, x, position)
}

func moveRobot(instruction string, plane [][]string, x *int, y *int) {
	switch {
	case instruction == "N":
		*x = *x + 1
	case instruction == "S":
		*x = *x - 1
	case instruction == "E":
		*y = *y + 1
	case instruction == "W":
		*y = *y - 1
	}
	if *x >= len(plane) || *y >= len(plane[*x]) {
		log.Fatal("Robot is trying to walk away from the rectangular plane")
	} else if plane[*x][*y] == "*" {
		log.Fatal("Robot is trying to walk where it already travelled")
	} else if plane[*x][*y] != "" {
		log.Fatal("Robot has particle on its way")
	} else {
		plane[*x][*y] = "*"
	}
}

func changeDirection(direction string, pos *string) {
	if direction == "R" {
		switch {
		case *pos == "N":
			*pos = "E"
		case *pos == "S":
			*pos = "W"
		case *pos == "E":
			*pos = "S"
		case *pos == "W":
			*pos = "N"
		}
	} else {
		switch {
		case *pos == "N":
			*pos = "W"
		case *pos == "S":
			*pos = "E"
		case *pos == "E":
			*pos = "N"
		case *pos == "W":
			*pos = "S"
		}
	}
}
