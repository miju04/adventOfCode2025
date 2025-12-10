package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func partOne(ingredientIds []int, freshRanges []Range) int {
	count := 0
	for _, id := range ingredientIds {
		for _, r := range freshRanges {
			if id >= r.start && id <= r.end {
				count++
				break
			}
		}
	}
	return count
}

func aggregateRanges(ranges []Range) []Range {
	var aggregated []Range
	for _, r := range ranges {
		if len(aggregated) == 0 {
			aggregated = append(aggregated, r)
		}
		if r.start > aggregated[len(aggregated)-1].end+1 {
			aggregated = append(aggregated, r)
		} else {
			aggregated[len(aggregated)-1].end = max(aggregated[len(aggregated)-1].end, r.end)
		}
	}
	return aggregated
}

func partTwo(freshRanges []Range) int {
	count := 0
	for _, r := range aggregateRanges(freshRanges) {
		count += r.end - r.start + 1
	}
	return count
}

func main() {
	input, _ := os.Open("input.txt")

	defer input.Close()

	scanner := bufio.NewScanner(input)

	var freshRanges []Range
	var ingredientIds []int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		temp := strings.Split(line, "-")
		start, _ := strconv.Atoi(temp[0])
		end, _ := strconv.Atoi(temp[1])

		freshRanges = append(freshRanges, Range{start, end})
	}

	for scanner.Scan() {
		id, _ := strconv.Atoi(scanner.Text())
		ingredientIds = append(ingredientIds, id)
	}

	slices.SortFunc(freshRanges, func(a, b Range) int {
		return a.start - b.start
	})

	countAvailableFreshIds := partOne(ingredientIds, freshRanges)

	countDistinctFreshIds := partTwo(freshRanges)

	log.Println("Number of ingredients used that are fresh:", countAvailableFreshIds)
	log.Println("Number of distinct fresh ingredients:", countDistinctFreshIds)
}
