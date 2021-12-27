package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	testFile      = "test.txt"
	inputFile     = "input1.txt"
	solutionTest1 = 35
	solutionTest2 = 3351
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
	iea, input := readInput(fileName)
	for round := 0; round < 2; round++ {
		input = step(iea, input, round)
	}
	return countOnes(input)
}

func solution2(fileName string) int {
	iea, input := readInput(fileName)
	for round := 0; round < 50; round++ {
		input = step(iea, input, round)
	}
	return countOnes(input)
}

func step(iea string, input [][]int, round int) [][]int {
	output := make([][]int, len(input)+2)
	newInput := make([][]int, len(input)+4)
	for i := 0; i < len(output); i++ {
		output[i] = make([]int, len(input)+2)
		newInput[i] = make([]int, len(input)+4)
	}
	newInput[len(newInput)-2] = make([]int, len(input)+4)
	newInput[len(newInput)-1] = make([]int, len(input)+4)
	if round%2 != 0 && iea[0] == '#' {
		for i := 0; i < len(newInput); i++ {
			newInput[0][i] = 1
			newInput[1][i] = 1
			newInput[i][0] = 1
			newInput[i][1] = 1
			newInput[len(newInput)-2][i] = 1
			newInput[len(newInput)-1][i] = 1
			newInput[i][len(newInput)-2] = 1
			newInput[i][len(newInput)-1] = 1
		}
	}
	for i := 2; i < len(newInput)-2; i++ {
		for j := 2; j < len(newInput)-2; j++ {
			newInput[i][j] = input[i-2][j-2]
		}
	}
	for i := 0; i < len(output); i++ {
		for j := 0; j < len(output); j++ {
			output[i][j] = numberFromGrid(iea, newInput, i+1, j+1)
		}
	}
	return output
}

func numberFromGrid(iea string, grid [][]int, x, y int) int {
	index := 0
	exp := 0
	for i := x + 1; i >= x-1; i-- {
		for j := y + 1; j >= y-1; j-- {
			if grid[i][j] == 1 {
				index += pow2(exp)
			}
			exp++
		}
	}
	if iea[index] == '#' {
		return 1
	}
	return 0
}

func pow2(exponent int) int {
	res := 1
	for i := 0; i < exponent; i++ {
		res = res * 2
	}
	return res
}

func countOnes(input [][]int) int {
	tot := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if input[i][j] == 1 {
				tot++
			}
		}
	}
	return tot
}

func readInput(fileName string) (string, [][]int) {
	dataRaw, _ := os.ReadFile(fileName)
	blocks := strings.Split(string(dataRaw), "\n\n")
	iea := blocks[0]
	inputRaw := strings.Split(blocks[1], "\n")
	input := make([][]int, len(inputRaw))
	for i := 0; i < len(input); i++ {
		input[i] = make([]int, len(inputRaw))
	}
	for i, row := range inputRaw {
		for j, c := range row {
			if c == '#' {
				input[i][j] = 1
			}
		}
	}
	return iea, input
}

func assertEqual(a, b int) {
	if a != b {
		fmt.Println(fmt.Sprintf("Test failed, got %d", b))
		os.Exit(1)
	}
}
