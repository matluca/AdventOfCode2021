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
	solutionTest1 = 150
	solutionTest2 = 900
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
	keys, values := readData(fileName)
	var x, y int
	for i := 0; i < len(keys); i++ {
		switch keys[i] {
		case "forward":
			x = x + values[i]
		case "down":
			y = y + values[i]
		case "up":
			y = y - values[i]
		}
	}
	return x * y
}

func solution2(fileName string) int {
	keys, values := readData(fileName)
	var x, y, a int
	for i := 0; i < len(keys); i++ {
		switch keys[i] {
		case "forward":
			x = x + values[i]
			y = y + values[i]*a
		case "down":
			a = a + values[i]
		case "up":
			a = a - values[i]
		}
	}
	return x*y
}

func readData(fileName string) ([]string, []int) {
	dataRaw, _ := os.ReadFile(fileName)
	dataStr := strings.Split(string(dataRaw), "\n")
	keys := make([]string, len(dataStr))
	values := make([]int, len(dataStr))
	for i := 0; i < len(dataStr); i++ {
		s := strings.Split(dataStr[i], " ")
		keys[i] = s[0]
		n, _ := strconv.Atoi(s[1])
		values[i] = n
	}
	return keys, values
}

func assertEqual(a, b int) {
	if a != b {
		fmt.Println(fmt.Sprintf("Test failed, got %d", b))
		os.Exit(1)
	}
}
