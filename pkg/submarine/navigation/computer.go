package navigation

import (
	"github.com/vmykhailyk/advent-of-code-2021/pkg/submarine"
	"strconv"
	"strings"
)

func PredictPosition(commands []string, coordinates submarine.Coordinates) submarine.Coordinates {
	coords := coordinates
	for _, command := range commands {
		moveForward(command, &coords)
		moveDown(command, &coords)
		moveUp(command, &coords)
	}
	return coords
}

func moveUp(command string, coordinates *submarine.Coordinates) {
	distance := getDistance(command, "up")
	if distance > 0 {
		coordinates.Depth -= distance
	}
}

func moveDown(command string, coordinates *submarine.Coordinates) {
	distance := getDistance(command, "down")
	if distance > 0 {
		coordinates.Depth += distance
	}
}

func moveForward(command string, coordinates *submarine.Coordinates) {
	distance := getDistance(command, "forward")
	if distance > 0 {
		coordinates.Position += distance
	}
}

func getDistance(command string, instruction string) int {
	if strings.HasPrefix(command, instruction) {
		distanceStr := strings.TrimSpace(strings.TrimPrefix(command, instruction))
		distance, _ := strconv.Atoi(distanceStr)
		return distance
	} else {
		return 0
	}
}
