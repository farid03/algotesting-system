package dataprocessing

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func SplitTests(file *os.File) []string {
	var testCases []string
	lines, err := readLines(file)

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var currentTestCase strings.Builder
	for _, line := range lines {
		if line != "$\n" {
			currentTestCase.WriteString(line)
		} else {
			testCases = append(testCases, currentTestCase.String())
			currentTestCase = strings.Builder{}
		}
	}
	return testCases
}

func readLines(file *os.File) ([]string, error) {
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text()+"\n")
	}
	return lines, scanner.Err()
}
