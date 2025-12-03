package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var input string
	var sum int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = scanner.Text()

		sum += maxBankJoltage(input)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	fmt.Printf("%d", sum)
}

func maxBankJoltage(bank string) int {
	var first int
	var second int
	for i, c := range bank {
		n, err := strconv.Atoi(string(c))

		if err != nil {
			log.Fatalf("Invalid joltage")
		}

		if first == 0 || n > first && len(bank) != i+1 {
			first = n
			second = 0
			continue
		}

		if second == 0 || n > second {
			second = n
			continue
		}
	}
	joltage, err := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(second))

	if err != nil {
		log.Fatalf("Could not calculate bank joltage")
	}
	return joltage
}
