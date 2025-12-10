package main

import (
	"bufio"
	"log"
	"os"
	"slices"
)

func countAroundPaper(value int, previous, current, next []int) int {
	count := 0
	for i := value - 1; i <= value+1; i++ {
		if slices.Contains(previous, i) {
			count++
		}
		if slices.Contains(current, i) {
			count++
		}
		if slices.Contains(next, i) {
			count++
		}
	}
	return count
}

func countNumberOfAccessed(previous, current, next []int) int {
	numberOfAccessed := 0
	for _, value := range current {
		count := countAroundPaper(value, previous, current, next)
		if count < 5 {
			numberOfAccessed++
		}
	}
	return numberOfAccessed
}

func partOne(indexedInput [][]int) int {
	var previousLineIndex, nextLineIndex []int
	numberOfAccessible := 0
	for i := range indexedInput {
		if i-1 < 0 {
			previousLineIndex = []int{}
		} else {
			previousLineIndex = indexedInput[i-1]
		}
		if i+1 >= len(indexedInput) {
			nextLineIndex = []int{}
		} else {
			nextLineIndex = indexedInput[i+1]
		}

		numberOfAccessible += countNumberOfAccessed(previousLineIndex, indexedInput[i], nextLineIndex)
	}

	return numberOfAccessible
}

func countAndRemoveAccessed(previous, current, next []int) (int, []int) {
	numberOfAccessed := 0
	for i := 0; i < len(current); i++ {
		value := current[i]
		count := countAroundPaper(value, previous, current, next)
		if count < 5 {
			numberOfAccessed++
			current = slices.Delete(current, i, i+1)
			i--
		}
	}
	return numberOfAccessed, current
}

func partTwo(indexedInput [][]int) int {
	var previousLineIndex, nextLineIndex []int
	numberOfAccessible := 0

	removed := true
	for removed {
		removed = false
		for i := range indexedInput {
			if i-1 < 0 {
				previousLineIndex = []int{}
			} else {
				previousLineIndex = indexedInput[i-1]
			}
			if i+1 >= len(indexedInput) {
				nextLineIndex = []int{}
			} else {
				nextLineIndex = indexedInput[i+1]
			}

			newAccessible, newCurrent := countAndRemoveAccessed(previousLineIndex, indexedInput[i], nextLineIndex)
			numberOfAccessible += newAccessible
			indexedInput[i] = newCurrent
			if newAccessible > 0 {
				removed = true
			}
		}
	}
	return numberOfAccessible
}

func main() {
	input, _ := os.Open("input.txt")

	defer input.Close()

	scanner := bufio.NewScanner(input)

	var indexedInput [][]int
	for scanner.Scan() {
		var indexes []int
		for i, char := range scanner.Text() {
			if char == '@' {
				indexes = append(indexes, i)
			}
		}
		indexedInput = append(indexedInput, indexes)
	}

	numberOfAccessiblePartOne := partOne(indexedInput)

	numberOfAccessiblePartTwo := partTwo(indexedInput)

	log.Printf("Number of accessible (Part 1): %d", numberOfAccessiblePartOne)
	log.Printf("Number of accessible (Part 2): %d", numberOfAccessiblePartTwo)
}
