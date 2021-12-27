package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	testFile       = "test.txt"
	inputFile      = "input1.txt"
	solutionTest1  = 37
	solutionTest2 = 168
)

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
	positions := readInput(fileName)
	return cost1(positions, median(positions))
}

func solution1NotNice(fileName string) int {
	positions := readInput(fileName)
	sort.Ints(positions)
	v := positions[0]
	currentCost := cost1(positions, v)
	for {
		nextCost := cost1(positions, v+1)
		if nextCost > currentCost {
			break
		}
		v++
		currentCost = nextCost
	}
	return currentCost
}

func solution2(fileName string) int {
	positions := readInput(fileName)
	sort.Ints(positions)
	v := positions[0]
	currentCost := cost2(positions, v)
	for {
		nextCost := cost2(positions, v+1)
		if nextCost > currentCost {
			break
		}
		v++
		currentCost = nextCost
	}
	return currentCost
}

func median(positions []int) int {
	sort.Ints(positions)
	if len(positions) % 2 ==1 {
		return positions[(len(positions)-1)/2]
	}
	return (positions[len(positions)/2-1] + positions[len(positions)/2]) / 2
}

func cost1(positions []int, center int) int {
	sum := 0
	for _, pos := range positions {
		sum = sum + int(math.Abs(float64(pos-center)))
	}
	return sum
}

func cost2(positions []int, center int) int {
	sum := 0
	for _, pos := range positions {
		d := int(math.Abs(float64(pos-center)))
		sum = sum + d*(d+1)/2
	}
	return sum
}

func readInput(fileName string) []int {
	dataRaw, _ := os.ReadFile(fileName)
	dataStr := strings.Split(string(dataRaw), ",")
	input := make([]int, len(dataStr))
	for i, str := range dataStr {
		input[i] , _ = strconv.Atoi(str)
	}
	return input
}

func assertEqual(a, b int) {
	if a != b {
		fmt.Println(fmt.Sprintf("Test failed, got %d", b))
		os.Exit(1)
	}
}
