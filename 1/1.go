package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	testFile = "test.txt"
	inputFile = "input1.txt"
	solutionTest1 = 7
	solutionTest2 = 5
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
	depths := readData(fileName)
	return countIncreased(depths)
}

func solution2(fileName string) int {
	depths := readData(fileName)
	windowSums := make([]int, len(depths)-2)
	for i:=0; i<len(windowSums); i++ {
		windowSums[i] = depths[i]+depths[i+1]+depths[i+2]
	}
	return countIncreased(windowSums)
}


func readData(fileName string) []int {
	dataRaw, _ := os.ReadFile(fileName)
	dataStr := strings.Split(string(dataRaw), "\n")
	dataInt := make([]int, len(dataStr))
	for i:=0; i<len(dataStr); i++ {
		n, _ := strconv.Atoi(dataStr[i])
		dataInt[i] = n
	}
	return dataInt
}

func countIncreased(input []int) int {
	counter := 0
	for i:=1; i<len(input); i++ {
		if input[i] > input[i-1] {
			counter++
		}
	}
	return counter
}

func assertEqual(a, b int) {
	if a != b {
		fmt.Println(fmt.Sprintf("Test failed, got %d", b))
		os.Exit(1)
	}
}