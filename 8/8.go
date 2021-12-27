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
	solutionTest1 = 26
	solutionTest2 = 61229
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
	_, outputs := readInput(fileName)
	tot := 0
	for _, p := range outputs {
		tot += countUnique(p)
	}
	return tot
}

func solution2(fileName string) int {
	patterns, outputs := readInput(fileName)
	sum := 0
	for i:=0; i<len(patterns); i++ {
		mapping := findMapping(patterns[i])
		sum += computeOutput(mapping, outputs[i])
	}
	return sum
}

func computeOutput(mapping map[string]int, output []string) int {
	return 1000*mapping[output[0]] + 100*mapping[output[1]] + 10*mapping[output[2]] + mapping[output[3]]
}

func findMapping(pattern []string) map[string]int {
	out := make(map[string]int, 10)
	numbers := make(map[int]string, 10)
	for _, p := range pattern {
		switch len(p) {
		case 2:
			out[p] = 1
			numbers[1] = p
		case 3:
			out[p] = 7
			numbers[7] = p
		case 4:
			out[p] = 4
			numbers[4] = p
		case 7:
			out[p] = 8
			numbers[8] = p
		}
	}
	numbers[3] = find3(pattern, numbers)
	out[numbers[3]] = 3
	numbers[9] = find9(pattern, numbers)
	out[numbers[9]] = 9
	numbers[5] = find5(pattern, numbers)
	out[numbers[5]] = 5
	numbers[2] = find2(pattern, numbers)
	out[numbers[2]] = 2
	numbers[0] = find0(pattern, numbers)
	out[numbers[0]] = 0
	numbers[6] = find6(pattern, numbers)
	out[numbers[6]] = 6
	return out
}

func find3(pattern []string, numbers map[int]string) string {
	for _, p := range pattern {
		if (len(p) == 5) && (contains(p, numbers[7])) {
			return p
		}
	}
	return ""
}

func find9(pattern []string, numbers map[int]string) string {
	for _, p := range pattern {
		if (len(p) == 6) && (contains(p, numbers[3])) {
			return p
		}
	}
	return ""
}

func find5(pattern []string, numbers map[int]string) string {
	for _, p := range pattern {
		if (len(p) == 5) && (contains(numbers[9], p)) && !(p == numbers[3]) {
			return p
		}
	}
	return ""
}

func find2(pattern []string, numbers map[int]string) string {
	for _, p := range pattern {
		if (len(p) == 5) && !(p == numbers[3]) && !(p == numbers[5]) {
			return p
		}
	}
	return ""
}

func find0(pattern []string, numbers map[int]string) string {
	for _, p := range pattern {
		if (len(p) == 6) && contains(p, numbers[1]) && !(p == numbers[9]) {
			return p
		}
	}
	return ""
}

func find6(pattern []string, numbers map[int]string) string {
	for _, p := range pattern {
		if (len(p) == 6) && !(p == numbers[0]) && !(p == numbers[9]) {
			return p
		}
	}
	return ""
}

func countUnique(patterns []string) int {
	sum := 0
	uniqueLengths := []int{2, 3, 4, 7}
	for _, p := range patterns {
		if valueIn(len(p), uniqueLengths) {
			sum++
		}
	}
	return sum
}

func valueIn(n int, list []int) bool {
	for _, l := range list {
		if n == l {
			return true
		}
	}
	return false
}

func contains(a string, b string) bool {
	for _, c := range b {
		if !strings.Contains(a, string(c)) {
			return false
		}
	}
	return true
}

func readInput(fileName string) ([][]string, [][]string) {
	dataRaw, _ := os.ReadFile(fileName)
	rows := strings.Split(string(dataRaw), "\n")
	patterns := make([][]string, len(rows))
	outputs := make([][]string, len(rows))
	for i, row := range rows {
		dataStr := strings.Split(row, " | ")
		patterns[i] = strings.Split(dataStr[0], " ")
		outputs[i] = strings.Split(dataStr[1], " ")
		for j := 0; j < len(patterns[i]); j++ {
			patterns[i][j] = sortString(patterns[i][j])
		}
		for j := 0; j < len(outputs[i]); j++ {
			outputs[i][j] = sortString(outputs[i][j])
		}
	}
	return patterns, outputs
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func assertEqual(a, b int) {
	if a != b {
		fmt.Println(fmt.Sprintf("Test failed, got %d", b))
		os.Exit(1)
	}
}
