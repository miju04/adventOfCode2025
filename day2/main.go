package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func isInvalidIDPartOne(idStr string) bool {
	length := len(idStr)
	if length%2 != 0 {
		return false
	}
	return idStr[:length/2] == idStr[length/2:]
}

func isInvalidIDPartTwo(idStr string) bool {
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

	var invalidIDValuePartOne, invalidIDValuePartTwo int64

	var wg sync.WaitGroup

	for _, valueStr := range values {
		wg.Add(1)
		go func(valueStr string) {
			defer wg.Done()
			ranges := strings.Split(valueStr, "-")

			begin, _ := strconv.Atoi(ranges[0])
			end, _ := strconv.Atoi(ranges[1])

			for i := begin; i <= end; i++ {
				idStr := strconv.Itoa(i)
				if isInvalidIDPartOne(idStr) {
					atomic.AddInt64(&invalidIDValuePartOne, int64(i))
				}
				if isInvalidIDPartTwo(idStr) {
					atomic.AddInt64(&invalidIDValuePartTwo, int64(i))
				}
			}
		}(valueStr)
	}

	wg.Wait()

	log.Printf("The value of the invalid IDs Part One is: %d", invalidIDValuePartOne)
	log.Printf("The value of the invalid IDs Part Two is: %d", invalidIDValuePartTwo)
}
