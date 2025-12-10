package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func calculateLargestJoltagePartOne(batteries []int) int {
	position := checkLargestInIntArray(batteries, 0, len(batteries)-2)
	joltage := batteries[position]
	position = checkLargestInIntArray(batteries, position+1, len(batteries)-1)
	joltage = (joltage * 10) + batteries[position]

	return joltage
}

func calculateLargestJoltagePartTwo(batteries []int) int {
	joltage, position := 0, 0
	for i := range 12 {
		end := len(batteries) - (12 - i)
		position = checkLargestInIntArray(batteries, position, end)
		joltage = (joltage * 10) + batteries[position]
		position++
	}

	return joltage
}

func checkLargestInIntArray(batteries []int, start, end int) int {
	position := start
	for i := start; i <= end; i++ {
		if batteries[i] == 9 {
			return i
		}
		if batteries[i] > batteries[position] {
			position = i
		}
	}

	return position
}

func main() {
	input, _ := os.Open("input.txt")

	defer input.Close()
	scanner := bufio.NewScanner(input)
	var accJoltagePartOne, accJoltagePartTwo int64

	var wg sync.WaitGroup
	for scanner.Scan() {
		wg.Add(1)
		line := scanner.Text()
		go func(line string) {
			defer wg.Done()
			var batteries []int
			for s := range strings.SplitSeq(line, "") {
				battery, _ := strconv.Atoi(strings.TrimSpace(s))
				batteries = append(batteries, battery)
			}

			atomic.AddInt64(&accJoltagePartOne, int64(calculateLargestJoltagePartOne(batteries)))
			atomic.AddInt64(&accJoltagePartTwo, int64(calculateLargestJoltagePartTwo(batteries)))
		}(line)
	}

	wg.Wait()

	log.Printf("accumulated largest joltage part one: %v\n", accJoltagePartOne)
	log.Printf("accumulated largest joltage part two: %v\n", accJoltagePartTwo)
}
