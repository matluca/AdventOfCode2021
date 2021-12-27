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
	solutionTest1 = 17
	solutionTest2 = 195
)

func main() {
	test1 := solution1(testFile)
	fmt.Println(fmt.Sprintf("Part 1 test: %d", solutionTest1))
	assertEqual(solutionTest1, test1)
	sol1 := solution1(inputFile)
	fmt.Println(fmt.Sprintf("Part 1: %d", sol1))

	solution2(testFile, 10, 10)
	solution2(inputFile, 50, 10)
}

type point struct {
	x int
	y int
}

type fold struct {
	dir   string
	value int
}

func solution1(fileName string) int {
	points, folds := readInput(fileName)
	afterOneFold := foldOnce(points, folds[0])
	simplified := simplifyMap(afterOneFold)
	return len(simplified)
}

func solution2(fileName string, sizeX, sizeY int) {
	points, folds := readInput(fileName)
	for _, f := range folds {
		points = foldOnce(points, f)
		points = simplifyMap(points)
	}
	printMap(points, sizeX, sizeY)
}

func foldOnce(points []point, f fold) []point {
	switch f.dir {
	case "x":
		for i, p := range points {
			if p.x > f.value {
				points[i] = point{x: 2*f.value-p.x, y: p.y}
			}
		}
	case "y":
		for i, p := range points {
			if p.y > f.value {
				points[i] = point{x: p.x, y: 2*f.value-p.y}
			}
		}
	}
	return points
}

func simplifyMap (points []point) []point {
	found := make(map[point]bool)
	var newPoints []point
	for _, p := range points {
		if !found[p] {
			found[p] = true
			newPoints = append(newPoints, p)
		}
	}
	return newPoints
}

func printMap(points []point, sizeX, sizeY int) {
	for y:=0; y<sizeY; y++ {
		for x:=0; x<sizeX; x++ {
			if isIn(point{x: x, y: y}, points) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func isIn(p point, points []point) bool {
	for _, q := range points {
		if p == q {
			return true
		}
	}
	return false
}

func readInput(fileName string) ([]point, []fold) {
	dataRaw, _ := os.ReadFile(fileName)
	rows := strings.Split(string(dataRaw), "\n\n")
	pointsRaw := strings.Split(rows[0], "\n")
	foldsRaw := strings.Split(rows[1], "\n")
	points := make([]point, len(pointsRaw))
	for i, pR := range pointsRaw {
		x, _ := strconv.Atoi(strings.Split(pR, ",")[0])
		y, _ := strconv.Atoi(strings.Split(pR, ",")[1])
		points[i] = point{x: x, y: y}
	}
	folds := make([]fold, len(foldsRaw))
	for i, fR := range foldsRaw {
		pacified := strings.Split(fR, "fold along ")[1]
		dir := strings.Split(pacified, "=")[0]
		v, _ := strconv.Atoi(strings.Split(pacified, "=")[1])
		folds[i] = fold{dir: dir, value: v}
	}
	return points, folds
}

func assertEqual(a, b int) {
	if a != b {
		fmt.Println(fmt.Sprintf("Test failed, got %d", b))
		os.Exit(1)
	}
}
