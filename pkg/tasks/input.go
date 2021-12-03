package tasks

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func PathToFile(file string) string {
	basePath := os.Getenv("DATA_DIR")
	if basePath == "" {
		ex, _ := os.Getwd()
		basePath = filepath.Dir(ex)
	}
	return filepath.Join(basePath, file)
}

func ReadLines(file string) []string {
	body, err := os.ReadFile(PathToFile(file))
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return strings.Split(string(body), "\n")
}

func ParseIntegers(input []string) []int {
	var result []int
	for _, v := range input {
		parsed, err := strconv.Atoi(v)
		result = append(result, parsed)
		if err != nil {
			log.Fatalf("unable to parse int: %v", v)
		}
	}
	return result
}

func ParseBits(input []string) []int {
	var result []int
	for _, v := range input {
		parsed, err := strconv.ParseInt(v, 2, 0)
		result = append(result, int(parsed))
		if err != nil {
			log.Fatalf("unable to parse int: %v", v)
		}
	}
	return result
}
