package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func onComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := range data {
		if data[i] == ',' {
			return i + 1, data[:i], nil
		}
	}
	if !atEOF {
		return 0, nil, nil
	}
	// Return the final token if there's data remaining at EOF
	if len(data) > 0 {
		return len(data), data, nil
	}
	// No more data
	return 0, nil, bufio.ErrFinalToken
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(onComma)

	var total int64 = 0

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())

		if text == "" {
			continue
		}

		total += sumInvalid(text)
	}
	fmt.Printf("%d", total)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}

func sumInvalid(text string) int64 {
	var sum int64 = 0
	parts := strings.Split(text, "-")

	if len(parts) != 2 {
		log.Fatalf("Invalid range")
	}

	first, err := strconv.ParseInt(parts[0], 10, 64)

	if err != nil {
		log.Fatalf("Could not parse start index")
	}

	last, err := strconv.ParseInt(parts[1], 10, 64)

	if err != nil {
		log.Fatalf("Could not parse end index")
	}

	for i := first; i <= last; i++ {
		n := strconv.FormatInt(i, 10)

		halfLen := len(n) / 2

		if halfLen == 0 {
			continue
		}

		left, right := n[:halfLen], n[halfLen:]

		if left == right {
			sum += i
		}
	}
	return sum
}
