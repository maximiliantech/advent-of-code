package common

import (
	"bufio"
	"log"
	"os"
)

// ReadAllLines reads all lines from a file and returns them as a slice of strings.
func ReadAllLines(filePath string) []string {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}
