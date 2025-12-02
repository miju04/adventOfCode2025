package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func isInvalidIDPartOne(id int) bool {
	idStr := strconv.Itoa(id)
	length := len(idStr)
	if length%2 != 0 {
		return false
	}
	return idStr[:length/2] == idStr[length/2:]
}

func isInvalidIDPartTwo(id int) bool {
	idStr := strconv.Itoa(id)
	length := len(idStr)
	for i := 1; i <= length/2; i++ {
		if length%i != 0 {
			continue
		}
		repeatingPart := idStr[:i]
		if strings.Repeat(repeatingPart, length/i) == idStr {
			return true
		}
	}
	return false
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	inputStr := strings.TrimSuffix(string(input), "\n")
	values := strings.Split(inputStr, ",")

	invalidIDValuePartOne, invalidIDValuePartTwo := 0, 0
	for value := range values {
		ranges := strings.Split(values[value], "-")
		begin, _ := strconv.Atoi(ranges[0])
		end, _ := strconv.Atoi(ranges[1])

		for i := begin; i <= end; i++ {
			if isInvalidIDPartOne(i) {
				invalidIDValuePartOne += i
			}
			if isInvalidIDPartTwo(i) {
				invalidIDValuePartTwo += i
			}
		}
	}

	log.Printf("The value of the invalid IDs Part One is: %d", invalidIDValuePartOne)
	log.Printf("The value of the invalid IDs Part Two is: %d", invalidIDValuePartTwo)
}
