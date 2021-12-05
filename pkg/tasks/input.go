package tasks

import (
	"log"
	"os"
	"path/filepath"
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

func Iterate(inputs []string, iterator func(input string, index int)) {
	for i, input := range inputs {
		iterator(input, i)
	}
}