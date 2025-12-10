package main

import (
	"bufio"
	"log"
	"os"
)

func partOne(grid [][]rune) int {
	splitCount := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 'S' || grid[row][col] == '|' {
				if row+1 < len(grid) {
					if grid[row+1][col] != '^' {
						grid[row+1][col] = '|'
					} else {
						splitCount++
						if col-1 >= 0 {
							grid[row+1][col-1] = '|'
						}
						if col+1 < len(grid[row]) {
							grid[row+1][col+1] = '|'
						}
					}
				}
			}
		}
	}
	return splitCount
}

func partTwo(grid [][]rune) int {
	numRows := len(grid)
	numCols := len(grid[0])

	sRow, sCol := -1, -1
	for r, row := range grid {
		for c, char := range row {
			if char == 'S' {
				sRow, sCol = r, c
				break
			}
		}
		if sRow != -1 {
			break
		}
	}

	currentPositions := make(map[int]int)
	currentPositions[sCol] = 1

	for r := sRow; r < numRows; r++ {
		nextPositions := make(map[int]int)
		for currentCol, count := range currentPositions {
			if currentCol < 0 || currentCol >= numCols {
				continue // Path went off the grid horizontally.
			}

			if grid[r][currentCol] == '^' {
				nextColLeft := currentCol - 1
				if nextColLeft >= 0 {
					nextPositions[nextColLeft] += count
				}
				nextColRight := currentCol + 1
				if nextColRight < numCols {
					nextPositions[nextColRight] += count
				}
			} else {
				nextPositions[currentCol] += count
			}
		}
		currentPositions = nextPositions
	}

	var totalTimelines int
	for _, count := range currentPositions {
		totalTimelines += count
	}

	return totalTimelines
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var gridForPartOne, gridForPartTwo [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gridForPartOne = append(gridForPartOne, []rune(scanner.Text()))
		gridForPartTwo = append(gridForPartTwo, []rune(scanner.Text()))
	}

	resultPartOne := partOne(gridForPartOne)
	resultPartTwo := partTwo(gridForPartTwo)

	log.Printf("Part One: %d", resultPartOne)
	log.Printf("Part Two: %d", resultPartTwo)
}
