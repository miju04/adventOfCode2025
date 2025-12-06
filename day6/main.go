package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func calculateCompleteValue(numbers [][]int, operations []string) int {
	addedValues := 0
	for column := range operations {
		columnValue := numbers[column][0]
		op := operations[column]
		for row := 1; row < len(numbers[column]); row++ {
			num := numbers[column][row]
			switch op {
			case "+":
				columnValue += num
			case "*":
				columnValue *= num
			}
		}
		addedValues += columnValue
	}
	return addedValues
}

func partOne(lines []string) int {
	var numbers [][]int
	var operations []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		i := 0
		for _, num := range strings.Split(line, " ") {
			if num == "" {
				continue
			}
			n, err := strconv.Atoi(num)
			if err != nil {
				operations = append(operations, num)
				continue
			}
			if len(numbers) <= i {
				numbers = append(numbers, []int{n})
			} else {
				numbers[i] = append(numbers[i], n)
			}
			i++
		}
	}
	return calculateCompleteValue(numbers, operations)
}

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func parseProblem(problem [][]string) ([]int, string) {
	var numbers []int
	var op string

	for _, pColumn := range problem {
		numStr := ""
		for _, char := range pColumn {
			if char >= "0" && char <= "9" {
				numStr += char
			}
		}
		if numStr != "" {
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
		}
		if len(pColumn) > 0 {
			lastChar := pColumn[len(pColumn)-1]
			if lastChar == "*" || lastChar == "+" {
				op = string(lastChar)
			}
		}
	}

	return numbers, op
}

func partTwo(lines []string) int {
	var grid [][]string
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	transposedGrid := transpose(grid)

	var problems [][][]string
	var currentProblem [][]string

	for i := len(transposedGrid) - 1; i >= 0; i-- {
		column := transposedGrid[i]
		isSeparator := true
		for _, char := range column {
			if char != " " {
				isSeparator = false
				break
			}
		}

		if isSeparator {
			if len(currentProblem) > 0 {
				problems = append(problems, currentProblem)
				currentProblem = [][]string{}
			}
		} else {
			currentProblem = append([][]string{column}, currentProblem...)
		}
	}
	if len(currentProblem) > 0 {
		problems = append(problems, currentProblem)
	}

	var allProblemsNumbers [][]int
	var allProblemsOperations []string

	for _, problem := range problems {
		numbers, op := parseProblem(problem)
		if len(numbers) > 0 {
			allProblemsNumbers = append(allProblemsNumbers, numbers)
			allProblemsOperations = append(allProblemsOperations, op)
		}
	}

	return calculateCompleteValue(allProblemsNumbers, allProblemsOperations)
}

func main() {
	input, _ := os.Open("test.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	resultPartOne := partOne(lines)
	resultPartTwo := partTwo(lines)
	log.Printf("Result part one: %d", resultPartOne)
	log.Printf("Result part two: %d", resultPartTwo)
}
