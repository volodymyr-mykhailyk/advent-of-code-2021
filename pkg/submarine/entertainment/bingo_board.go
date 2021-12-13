package entertainment

import (
	"strconv"
	"strings"
)

type BingoBoard [][]int

func (board BingoBoard) refresh(refresher func(v int) int) {
	for _, line := range board {
		for x, value := range line {
			line[x] = refresher(value)
		}
	}
}

func (board BingoBoard) PlayNumber(number int) bool {
	board.refresh(func(v int) int {
		if v == number {
			return -1
		} else {
			return v
		}
	})
	return board.IsWinning()
}

func (board BingoBoard) Score() int {
	sum := 0
	board.refresh(func(v int) int {
		if v != -1 {
			sum += v
		}
		return v
	})
	return sum
}

func (board BingoBoard) IsWinning() bool {
	for i, _ := range board {
		if isWinningLine(board.HLine(i)) || isWinningLine(board.VLine(i)) {
			return true
		}
	}
	return false
}

func (board BingoBoard) VLine(x int) []int {
	slice := make([]int, len(board))
	for y, line := range board {
		slice[y] = line[x]
	}
	return slice
}

func (board BingoBoard) HLine(y int) []int {
	return board[y]
}

func isWinningLine(line []int) bool {
	for _, v := range line {
		if v != -1 {
			return false
		}
	}
	return true
}

func BingoFromString(input []string) BingoBoard {
	maxY := len(input)
	board := make(BingoBoard, maxY)
	for y, line := range input {
		numbers := strings.Fields(line)
		board[y] = make([]int, len(numbers))
		for x, value := range numbers {
			v, _ := strconv.ParseInt(value, 10, 0)
			board[y][x] = int(v)
		}
	}
	return board
}
