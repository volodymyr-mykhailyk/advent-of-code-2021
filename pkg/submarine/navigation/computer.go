package navigation

import (
	"github.com/vmykhailyk/advent-of-code-2021/pkg/submarine"
	"strconv"
	"strings"
)

func PredictLinearPosition(commands []string, coordinates submarine.Coordinates) submarine.Coordinates {
	for _, command := range commands {
		coordinates.Position += getDistance(command, "forward")
		coordinates.Depth += getDistance(command, "down")
		coordinates.Depth -= getDistance(command, "up")
	}
	return coordinates
}

func PredictAimedPosition(commands []string, coordinates submarine.Coordinates) submarine.Coordinates {
	for _, command := range commands {
		coordinates.Position += getDistance(command, "forward")
		coordinates.Depth += getDistance(command, "forward") * coordinates.Aim
		coordinates.Aim += getDistance(command, "down")
		coordinates.Aim -= getDistance(command, "up")
	}
	return coordinates
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
