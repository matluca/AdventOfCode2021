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
	solutionTest1 = 40
	solutionTest2 = 315
)

func main() {
	test1 := solution1(testFile)
	fmt.Println(fmt.Sprintf("Part 1 test: %d", solutionTest1))
	assertEqual(solutionTest1, test1)
	sol1 := solution1(inputFile)
	fmt.Println(fmt.Sprintf("Part 1: %d", sol1))

	//test2 := solution2(testFile)
	//fmt.Println(fmt.Sprintf("Part 2 test: %d", solutionTest2))
	//assertEqual(solutionTest2, test2)
	//sol2 := solution2(inputFile)
	//fmt.Println(fmt.Sprintf("Part 2: %d", sol2))
}

func solution1(fileName string) int {
	cavern := readInput(fileName)
	size := len(cavern)
	return minPathScore(cavern, size)
}

//func solution2(fileName string) int {
//	cavern := readInput(fileName)
//	size := len(cavern)
//	newCavern := buildLargerMap(cavern, size)
//	return minPathScore(newCavern, 5*size)
//}

func buildLargerMap(cavern [][]int, size int) [][]int {
	newCavern := make([][]int, 5*size)
	for i := 0; i < len(newCavern); i++ {
		newCavern[i] = make([]int, 5*size)
	}
	for i:=0; i<len(newCavern); i++ {
		for j:=0; j<len(newCavern); j++ {
			offset := i/size + j/size
			newCavern[i][j] = (cavern[i%size][j%size])+offset
			if newCavern[i][j]>9 {
				newCavern[i][j] = newCavern[i][j]-9
			}
		}
	}
	return newCavern
}

func printCavern(cavern [][]int, size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(cavern[i][j])
		}
		fmt.Println()
	}
}

func minPathScore(cavern [][]int, size int) int {
	score := cavern
	score[0][0] = 0
	for j := 1; j < size; j++ {
		score[0][j] = score[0][j] + score[0][j-1]
		score[j][0] = score[j][0] + score[j-1][0]
	}
	for i := 1; i < size; i++ {
		for j := 1; j < size; j++ {
			score[i][j] = score[i][j] + min(score[i-1][j], score[i][j-1])
		}
	}
	return score[size-1][size-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func readInput(fileName string) [][]int {
	dataRaw, _ := os.ReadFile(fileName)
	rows := strings.Split(string(dataRaw), "\n")
	cavern := make([][]int, len(rows))
	for i, row := range rows {
		cavern[i] = make([]int, len(row))
		for j, c := range row {
			n, _ := strconv.Atoi(string(c))
			cavern[i][j] = n
		}
	}
	return cavern
}

func assertEqual(a, b int) {
	if a != b {
		fmt.Println(fmt.Sprintf("Test failed, got %d", b))
		os.Exit(1)
	}
}
