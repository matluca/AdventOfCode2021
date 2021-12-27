package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	testFile       = "test.txt"
	inputFile      = "input1.txt"
	solutionTest1  = 26
	solutionTest1b = 5934
	solutionTest2 = 26984457539
)

func main() {
	test1 := solution1(testFile, 18)
	fmt.Println(fmt.Sprintf("Part 1 test: %d", solutionTest1))
	assertEqual(solutionTest1, test1)
	test1b := solution1(testFile, 80)
	fmt.Println(fmt.Sprintf("Part 1 test b: %d", solutionTest1b))
	assertEqual(solutionTest1b, test1b)
	sol1 := solution1(inputFile, 80)
	fmt.Println(fmt.Sprintf("Part 1: %d", sol1))

	test2 := solution1(testFile, 256)
	fmt.Println(fmt.Sprintf("Part 2 test: %d", solutionTest2))
	assertEqual(solutionTest2, test2)
	sol2 := solution1(inputFile, 256)
	fmt.Println(fmt.Sprintf("Part 2: %d", sol2))
}

func population(input[]int) [9]int {
	var pop [9]int
	for _, v := range input {
		pop[v]++
	}
	return pop
}

func solution1(fileName string, nDays int) int {
	input := readInput(fileName)
	pop := population(input)
	for d:=1; d<=nDays; d++ {
		pop = nextStep(pop)
	}
	return sum(pop)
}

func nextStep(pop [9]int) [9]int {
	var newPop [9]int
	for i:=0; i<8; i++ {
		newPop[i] = pop[i+1]
	}
	newPop[8] = pop[0]
	newPop[6] += pop[0]
	return newPop
}

func sum(pop [9]int) int {
	sum :=0
	for _, v := range pop {
		sum += v
	}
	return sum
}

func readInput(fileName string) []int {
	dataRaw, _ := os.ReadFile(fileName)
	dataStr := strings.Split(string(dataRaw), ",")
	input := make([]int, len(dataStr))
	for i, str := range dataStr {
		input[i], _ = strconv.Atoi(str)
	}
	return input
}

func assertEqual(a, b int) {
	if a != b {
		fmt.Println(fmt.Sprintf("Test failed, got %d", b))
		os.Exit(1)
	}
}
