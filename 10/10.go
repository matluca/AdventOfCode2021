package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	testFile      = "test.txt"
	inputFile     = "input1.txt"
	solutionTest1 = 26397
	solutionTest2 = 288957
)

type point struct {
	x int
	y int
}

var opposite = map[string]string{
	")": "(",
	"]": "[",
	">": "<",
	"}": "{",
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var points = map[string]int{
	")": 3,
	"]": 57,
	">": 25137,
	"}": 1197,
}

var mult = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

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
	input := readInput(fileName)
	tot := 0
	for _, r := range input {
		tot += checkLine(r)
	}
	return tot
}

func solution2(fileName string) int {
	input := readInput(fileName)
	var scores []int
	for _, r := range input {
		score := checkLine2(r)
		if score != 0 {
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	l := len(scores)
	return scores[(l-1)/2]
}

func checkLine2(line []string) int {
	var opened []string
	for i := 0; i < len(line); i++ {
		switch line[i] {
		case "(", "<", "[", "{":
			opened = append(opened, line[i])
		case ")", ">", "]", "}":
			l := len(opened)
			if opened[l-1] == opposite[line[i]] {
				opened = opened[:l-1]
			} else {
				return 0
			}
		}
	}
	score := 0
	for i:=len(opened)-1; i>=0; i-- {
		opp := opposite[opened[i]]
		score *= 5
		score += mult[opp]
	}
	return score
}

func checkLine(line []string) int {
	var opened []string
	for i := 0; i < len(line); i++ {
		switch line[i] {
		case "(", "<", "[", "{":
			opened = append(opened, line[i])
		case ")", ">", "]", "}":
			l := len(opened)
			if opened[l-1] == opposite[line[i]] {
				opened = opened[:l-1]
			} else {
				return points[line[i]]
			}
		}
	}
	return 0
}

func readInput(fileName string) [][]string {
	dataRaw, _ := os.ReadFile(fileName)
	rows := strings.Split(string(dataRaw), "\n")
	output := make([][]string, len(rows))
	for i := 0; i < len(output); i++ {
		output[i] = make([]string, len(rows[i]))
		for j := 0; j < len(rows[i]); j++ {
			output[i][j] = string(rows[i][j])
		}
	}
	return output
}

func assertEqual(a, b int) {
	if a != b {
		fmt.Println(fmt.Sprintf("Test failed, got %d", b))
		os.Exit(1)
	}
}
