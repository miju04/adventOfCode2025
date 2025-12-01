package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const dialSize = 100

func mod(a, b int) int {
	return ((a % b) + b) % b
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}

	defer input.Close()
	scanner := bufio.NewScanner(input)

	dial := 50
	counterOnZero := 0
	counterOverZero := 0

	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)

		value, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		if line[0] == 'L' {
			value *= -1
		}

		counterOverZero += calculateTimesOverZero(dial, value)

		dial = mod(dial+value, dialSize)

		if dial == 0 {
			counterOnZero++
			log.Printf("Dial hit exactly zero! Current count: %d\n", counterOnZero)
		}
		log.Printf("Dial position: %d\n", dial)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Final number of dial hits exactly zero: %d\n", counterOnZero)
	log.Printf("Final number of dial over zero: %d\n", counterOverZero)
}

func calculateTimesOverZero(currentDial, value int) int {
	rotations := value / dialSize
	if rotations < 0 {
		rotations *= -1
	}

	value = value % dialSize

	// right over zero
	if currentDial+value >= dialSize {
		rotations++
	}

	// left over zero
	if currentDial > 0 && currentDial+value <= 0 {
		rotations++
	}

	log.Printf("Calculated times over zero: %d\n", rotations)
	return rotations
}
