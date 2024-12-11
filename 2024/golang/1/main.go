package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input string
	var parts []string
	left, right := []int{}, []int{}
	var tuple []int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = scanner.Text()
		parts = strings.Fields(input)

		tuple = parseDigits(parts)
		left, right = append(left, tuple[0]), append(right, tuple[1])
	}

	sort.Ints(left)
	sort.Ints(right)

	sum := 0

	for i := 0; i < len(left); i++ {
		line := diff(left[i], right[i])
		sum += line
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print(sum)
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
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
