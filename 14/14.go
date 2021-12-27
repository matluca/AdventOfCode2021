package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	testFile      = "test.txt"
	inputFile     = "input1.txt"
	solutionTest1 = 1588
	solutionTest2 = 2188189693529
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
	template, pairs := readInput(fileName)
	for i:=0; i<10; i++ {
		template = step(template, pairs)
	}
	return MinMax(occurrences(template))
}

func solution2(fileName string) int {
	template, pairs := readInput(fileName)
	startLetter := string(template[0])
	endLetter := string(template[len(template)-1])
	pairCount := buildInitialPairCount(template)
	for i:=0; i<40; i++ {
		pairCount = step2(pairCount, pairs)
	}
	occ := occurrences2(pairCount, startLetter, endLetter)
	return MinMax(occ)
}

func buildInitialPairCount(in string) map[string]int {
	pairCount := make(map[string]int)
	for i:=0; i<len(in)-1; i++ {
		if _, ok := pairCount[in[i:i+2]]; !ok {
			pairCount[in[i:i+2]] = 0
		}
		pairCount[in[i:i+2]]++
	}
	return pairCount
}

func occurrences2(pairCount map[string]int, startLetter, endLetter string) map[string]int {
	occ := make(map[string]int)
	for key, value := range pairCount {
		a := string(key[0])
		b := string(key[1])
		if _, ok := occ[a]; !ok {
			occ[a] = 0
		}
		if _, ok := occ[b]; !ok {
			occ[b] = 0
		}
		occ[a]+=value
		occ[b]+=value
	}
	occ[startLetter]++
	occ[endLetter]++
	for key, _ := range occ {
		occ[key] = occ[key]/2
	}
	return occ
}

func step2(inPairCount map[string]int, pairs map[string]string) map[string]int {
	outPairCount := make(map[string]int)
	for key, value := range inPairCount {
		if letter, ok := pairs[key]; ok {
			start := string(key[0]) + letter
			end := letter + string(key[1])
			if _, ok2 := outPairCount[start]; !ok2 {
				outPairCount[start] = 0
			}
			if _, ok2 := outPairCount[end]; !ok2 {
				outPairCount[end] = 0
			}
			outPairCount[start] += value
			outPairCount[end] += value
		} else {
			if _, ok2 := outPairCount[key]; !ok2 {
				outPairCount[key] = 0
			}
			outPairCount[key] += value
		}
	}
	return outPairCount
}

func step(in string, pairs map[string]string) string {
	out := string(in[0])
	for i:=0; i<len(in)-1; i++ {
		if insert, ok := pairs[in[i:i+2]]; ok {
			out = out + insert + string(in[i+1])
		}
	}
	return out
}

func occurrences(in string) map[string]int {
	occ := make(map[string]int)
	for _, l := range in {
		if _, ok := occ[string(l)]; !ok {
			occ[string(l)] = 0
		}
		occ[string(l)]++
	}
	return occ
}

func MinMax(occ map[string]int) int {
	min := occ["B"]
	max := occ["B"]
	for _, v := range occ {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max-min
}

func readInput(fileName string) (string, map[string]string) {
	dataRaw, _ := os.ReadFile(fileName)
	rows := strings.Split(string(dataRaw), "\n\n")
	template := rows[0]
	pairsRaw := strings.Split(rows[1], "\n")
	pairs := make(map[string]string, len(pairsRaw))
	for _, pR := range pairsRaw {
		a := strings.Split(pR, " -> ")[0]
		b := strings.Split(pR, " -> ")[1]
		pairs[a] = b
	}
	return template, pairs
}

func assertEqual(a, b int) {
	if a != b {
		fmt.Println(fmt.Sprintf("Test failed, got %d", b))
		os.Exit(1)
	}
}
