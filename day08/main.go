package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type vector3d struct {
	x, y, z int
}

type pair struct {
	a, b     int
	distance int
}

func calculateDistance(v1, v2 vector3d) int {
	dx := v1.x - v2.x
	dy := v1.y - v2.y
	dz := v1.z - v2.z
	// dont use sqrt to avoid float operations
	return dx*dx + dy*dy + dz*dz
}

func calculateDistances(junctionBoxes []vector3d) []pair {
	var pairs []pair
	for i := range junctionBoxes {
		for j := i + 1; j < len(junctionBoxes); j++ {
			dist := calculateDistance(junctionBoxes[i], junctionBoxes[j])
			pairs = append(pairs, pair{i, j, dist})
		}
	}
	return pairs
}

func find(parent []int, i int) int {
	if parent[i] != i {
		parent[i] = find(parent, parent[i])
	}
	return parent[i]
}

func union(parent []int, a, b int) bool {
	rootA := find(parent, a)
	rootB := find(parent, b)
	if rootA != rootB {
		parent[rootB] = rootA
		return true
	}
	return false
}

func partOne(junctionBoxes []vector3d) int {
	pairs := calculateDistances(junctionBoxes)
	slices.SortFunc(pairs, func(a, b pair) int {
		return a.distance - b.distance
	})

	// Disjoint Set Union structure
	parent := make([]int, len(junctionBoxes))
	for i := range parent {
		parent[i] = i
	}

	for i := 0; i < 1000 && i < len(pairs); i++ {
		p := pairs[i]
		union(parent, p.a, p.b)
	}

	// calculate sizes of connections
	sizeMap := make(map[int]int)
	for i := range junctionBoxes {
		root := find(parent, i)
		sizeMap[root]++
	}

	var circuitSizes []int
	for _, size := range sizeMap {
		circuitSizes = append(circuitSizes, size)
	}

	slices.SortFunc(circuitSizes, func(a, b int) int {
		return b - a
	})

	result := 1
	for i := 0; i < 3 && i < len(circuitSizes); i++ {
		result *= circuitSizes[i]
	}

	return result
}

func partTwo(junctionBoxes []vector3d) int {
	pairs := calculateDistances(junctionBoxes)
	slices.SortFunc(pairs, func(a, b pair) int {
		return a.distance - b.distance
	})

	// Disjoint Set Union structure
	parent := make([]int, len(junctionBoxes))
	for i := range parent {
		parent[i] = i
	}

	result := 0
	numCircuits := len(junctionBoxes)

	for _, p := range pairs {
		if union(parent, p.a, p.b) {
			numCircuits--
			if numCircuits == 1 {
				v1 := junctionBoxes[p.a]
				v2 := junctionBoxes[p.b]
				result = v1.x * v2.x
				break
			}
		}
	}

	return result
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var junctionBoxes []vector3d
	for scanner.Scan() {
		line := scanner.Text()
		var v vector3d
		_, _ = fmt.Sscanf(line, "%d,%d,%d", &v.x, &v.y, &v.z)
		junctionBoxes = append(junctionBoxes, v)
	}

	resultPartOne := partOne(junctionBoxes)
	reviewResultPartTwo := partTwo(junctionBoxes)
	log.Printf("Part One: %d", resultPartOne)
	log.Printf("Part Two: %d", reviewResultPartTwo)
}
