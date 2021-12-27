package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	testFile      = "test.txt"
	inputFile     = "input1.txt"
	solutionTest1 = 5
	solutionTest2 = 12
)

type point struct {
	x int
	y int
}

type line struct {
	a *point
	b *point
}

var grid [1000][1000]int

func main() {
	test1 := solution1(testFile)
	fmt.Println(fmt.Sprintf("Part 1 test: %d", solutionTest1))
	assertEqual(solutionTest1, test1)
	sol1 := solution1(inputFile)
	fmt.Println(fmt.Sprintf("Part 1: %d", sol1))

	test2 := solution2(testFile)
	fmt.Println(fmt.Sprintf("Part 2 test: %d", solutionTest2))
	assertEqual(solutionTest2, test2)
	sol2 := solution2(inputFile)
	fmt.Println(fmt.Sprintf("Part 2: %d", sol2))
}

func solution1(fileName string) int {
	reset()
	lines := readData(fileName)
	for _, l := range(lines) {
		coverLine1(l)
	}
	return countBiggerThan1()
}

func solution2(fileName string) int {
	reset()
	lines := readData(fileName)
	for _, l := range(lines) {
		coverLine2(l)
	}
	return countBiggerThan1()
}

func readData(fileName string) ([]line) {
	dataRaw, _ := os.ReadFile(fileName)
	dataStr := strings.Split(string(dataRaw), "\n")
	lines := make([]line, len(dataStr))
	for i:=0; i<len(lines); i++ {
		pointsRaw := strings.Split(dataStr[i], " -> ")
		var points [2]point
		for j:=0; j<2; j++ {
			coord := strings.Split(pointsRaw[j], ",")
			x, _ := strconv.Atoi(coord[0])
			y, _ := strconv.Atoi(coord[1])
			points[j] = point{x, y}
		}
		lines[i] = line{&points[0], &points[1]}
	}
	return lines
}

func coverLine1 (line line) {
	if line.a.x == line.b.x {
		start := min(line.a.y, line.b.y)
		end := max(line.a.y, line.b.y)
		for j:=start; j<=end; j++ {
			grid[line.a.x][j]++
		}
		return
	}
	if line.a.y == line.b.y {
		start := min(line.a.x, line.b.x)
		end := max(line.a.x, line.b.x)
		for j:=start; j<=end; j++ {
			grid[j][line.a.y]++
		}
		return
	}
	return
}

func coverLine2 (line line) {
	coverLine1(line)
	if line.a.x < line.b.x {
		if line.a.y < line.b.y {
			for i:=0; i<=line.b.x-line.a.x; i++ {
				grid[line.a.x+i][line.a.y+i]++
			}
		}
		if line.a.y > line.b.y {
			for i:=0; i<=line.b.x-line.a.x; i++ {
				grid[line.a.x+i][line.a.y-i]++
			}
		}
	}
	if line.a.x > line.b.x {
		if line.a.y < line.b.y {
			for i:=0; i<=line.a.x-line.b.x; i++ {
				grid[line.a.x-i][line.a.y+i]++
			}
		}
		if line.a.y > line.b.y {
			for i:=0; i<=line.a.x-line.b.x; i++ {
				grid[line.a.x-i][line.a.y-i]++
			}
		}
	}
}

func countBiggerThan1() int {
	size := len(grid)
	tot := 0
	for i:=0; i<size; i++ {
		for j:=0; j<size; j++ {
			if grid[i][j] > 1 {
				tot++
			}
		}
	}
	return tot
}

func reset() {
	grid = [1000][1000]int{}
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func assertEqual(a, b int) {
	if a != b {
		fmt.Println(fmt.Sprintf("Test failed, got %d", b))
		os.Exit(1)
	}
}
