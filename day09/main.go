package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type vector2d struct {
	x, y int
}

type pair struct {
	a, b vector2d
	size int
}

func calculateFieldSize(v1, v2 vector2d) int {
	var width, height int
	if v1.x > v2.x {
		width = v1.x - v2.x
	} else {
		width = v2.x - v1.x
	}
	if v1.y > v2.y {
		height = v1.y - v2.y
	} else {
		height = v2.y - v1.y
	}
	return (width + 1) * (height + 1)
}

func getPolygonEdges(vectors []vector2d) []pair {
	if len(vectors) < 2 {
		return []pair{}
	}
	var borders []pair
	for i := 0; i < len(vectors)-1; i++ {
		borders = append(borders, pair{a: vectors[i], b: vectors[i+1]})
	}
	borders = append(borders, pair{a: vectors[len(vectors)-1], b: vectors[0]})
	return borders
}

func intersectsWithInterior(rect pair, polygonEdges []pair) bool {
	minX, maxX := rect.a.x, rect.b.x
	if minX > maxX {
		minX, maxX = maxX, minX
	}
	minX++
	maxX--
	minY, maxY := rect.a.y, rect.b.y
	if minY > maxY {
		minY, maxY = maxY, minY
	}
	minY++
	maxY--

	for _, edge := range polygonEdges {
		// Check for intersection between a vertical polygon edge and a inner horizontal rectangle side
		if edge.a.x == edge.b.x {
			polyX := edge.a.x
			if polyX > minX && polyX < maxX {
				polyMinY, polyMaxY := edge.a.y, edge.b.y
				if polyMinY > polyMaxY {
					polyMinY, polyMaxY = polyMaxY, polyMinY
				}

				if (minY > polyMinY && minY < polyMaxY) || (maxY > polyMinY && maxY < polyMaxY) {
					return true
				}
			}
		}

		// Check for intersection between a horizontal polygon edge and a inner vertical rectangle side
		if edge.a.y == edge.b.y {
			polyY := edge.a.y
			if polyY > minY && polyY < maxY {
				polyMinX, polyMaxX := edge.a.x, edge.b.x
				if polyMinX > polyMaxX {
					polyMinX, polyMaxX = polyMaxX, polyMinX
				}

				if (minX > polyMinX && minX < polyMaxX) || (maxX > polyMinX && maxX < polyMaxX) {
					return true
				}
			}
		}
	}
	return false
}

func partOne(vector2ds []vector2d) int {
	var largest int
	for i := range vector2ds {
		for j := i + 1; j < len(vector2ds); j++ {
			area := calculateFieldSize(vector2ds[i], vector2ds[j])
			if area > largest {
				largest = area
			}
		}
	}
	return largest
}

func partTwo(vector2ds []vector2d) int {
	polygonEdges := getPolygonEdges(vector2ds)

	var pairs []pair
	for i := range vector2ds {
		for j := i + 1; j < len(vector2ds); j++ {
			area := calculateFieldSize(vector2ds[i], vector2ds[j])
			pairs = append(pairs, pair{vector2ds[i], vector2ds[j], area})
		}
	}

	slices.SortFunc(pairs, func(a, b pair) int {
		return b.size - a.size
	})

	var largestContained int
	for _, p := range pairs {
		intersected := intersectsWithInterior(p, polygonEdges)
		if !intersected {
			largestContained = p.size
			break
		}
	}

	return largestContained
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var vector2ds []vector2d

	for scanner.Scan() {
		line := scanner.Text()
		var v vector2d
		_, _ = fmt.Sscanf(line, "%d,%d", &v.x, &v.y)
		vector2ds = append(vector2ds, v)
	}

	resultPartOne := partOne(vector2ds)
	resultPartTwo := partTwo(vector2ds)
	log.Println("Part One Result:", resultPartOne)
	log.Println("Part Two Result:", resultPartTwo)
}
