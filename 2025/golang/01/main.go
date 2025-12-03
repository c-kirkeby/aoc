package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type dial struct {
	arrow   int
	counter int
}

func main() {
	var input string
	line := 0
	d := dial{arrow: 50, counter: 0}

	log.Printf("The dial starts by pointing at %d", d.arrow)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = scanner.Text()
		d.rotate(input)
		line = line + 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	fmt.Print(d.counter)
}

func (d *dial) rotate(instruction string) {
	if instruction == "" {
		return
	}
	runes := []rune(instruction)
	direction := runes[0]
	steps, err := strconv.Atoi(string(runes[1:]))

	if err != nil {
		log.Fatalf("Could not parse steps adjustment")
	}

	switch direction {
	case 'R':
		d.arrow += steps
	case 'L':
		d.arrow -= steps
	default:
		log.Printf("Unrecognised direction")
	}

	d.arrow = d.arrow % 100

	if d.arrow < 0 {
		d.arrow = 100 + d.arrow
	} else if d.arrow > 99 {
		d.arrow = d.arrow - 100
	}

	log.Printf("The dial is rotated %s to point at %d", instruction, d.arrow)

	if d.arrow == 0 {
		d.counter = d.counter + 1
	}
}
