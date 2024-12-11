package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var input string
	var parts []string
	safeCount := 0

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = scanner.Text()
		parts = strings.Fields(input)

		levels := parseDigits(parts)
		if isSafe(levels) {
			safeCount += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(safeCount)
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func isSafe(line []int) bool {
	tolerance := 1
	problems := 0

	reverse := make([]int, len(line))
	copy(reverse, line)
	slices.Reverse(reverse)
	if !slices.IsSorted(line) && !slices.IsSorted(reverse) {
		return false
	}

	for i, num := range line {
		if len(line)-1 != i {
			fluctuation := diff(num, line[i+1])
			if fluctuation < 1 || fluctuation > 3 {
				problems += 1
			}
		}
	}

	return problems <= tolerance
}

func parseDigits(parts []string) []int {
	result := make([]int, len(parts))

	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			log.Fatalf("Could not convert %v to integer: %v", part, err)
		}
		result[i] = num
	}
	return result
}
