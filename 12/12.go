package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

const (
	testFile       = "test.txt"
	testFileb      = "testb.txt"
	testFilec      = "testc.txt"
	inputFile      = "input1.txt"
	solutionTest1  = 10
	solutionTest1b = 19
	solutionTest1c = 226
	solutionTest2  = 36
	solutionTest2b = 103
	solutionTest2c = 3509
)

func main() {
	test1 := solution1(testFile)
	fmt.Println(fmt.Sprintf("Part 1 test: %d", solutionTest1))
	assertEqual(solutionTest1, test1)
	test1b := solution1(testFileb)
	fmt.Println(fmt.Sprintf("Part 1 test b: %d", solutionTest1b))
	assertEqual(solutionTest1b, test1b)
	test1c := solution1(testFilec)
	fmt.Println(fmt.Sprintf("Part 1 test c: %d", solutionTest1c))
	assertEqual(solutionTest1c, test1c)
	sol1 := solution1(inputFile)
	fmt.Println(fmt.Sprintf("Part 1: %d", sol1))

	test2 := solution2(testFile)
	fmt.Println(fmt.Sprintf("Part 2 test: %d", solutionTest2))
	assertEqual(solutionTest2, test2)
	test2b := solution2(testFileb)
	fmt.Println(fmt.Sprintf("Part 2 test b: %d", solutionTest2b))
	assertEqual(solutionTest2b, test2b)
	test2c := solution2(testFilec)
	fmt.Println(fmt.Sprintf("Part 2 test c: %d", solutionTest2c))
	assertEqual(solutionTest2c, test2c)
	sol2 := solution2(inputFile)
	fmt.Println(fmt.Sprintf("Part 2: %d", sol2))
}

func solution1(fileName string) int {
	pairs := readData(fileName)
	linkMap := buildMap(pairs)
	paths := findPaths("start", []string{"start"}, linkMap)
	return len(paths)
}

func solution2(fileName string) int {
	pairs := readData(fileName)
	linkMap := buildMap(pairs)
	paths := findPaths2("start", []string{"start"}, linkMap, false)
	return len(paths)
}

func findPaths2(startNode string, explored []string, linkMap map[string][]string, doubleVisitedLowercase bool) [][]string {
	var paths [][]string
	for _, nextNode := range linkMap[startNode] {
		hereExplored := explored
		hereDoubleVisitedLowercase := doubleVisitedLowercase
		if (isIn(nextNode, explored)) && (IsLower(nextNode)) {
			if hereDoubleVisitedLowercase || (nextNode == "start") {
				continue
			}
			hereDoubleVisitedLowercase = true
		}
		hereExplored = append(hereExplored, nextNode)
		if nextNode == "end" {
			paths = append(paths, hereExplored)
			continue
		}
		paths = append(paths, findPaths2(nextNode, hereExplored, linkMap, hereDoubleVisitedLowercase)...)
	}
	return paths
}

func findPaths(startNode string, explored []string, linkMap map[string][]string) [][]string {
	var paths [][]string
	for _, nextNode := range linkMap[startNode] {
		hereExplored := explored
		if (isIn(nextNode, explored)) && (IsLower(nextNode)) {
			continue
		}
		hereExplored = append(hereExplored, nextNode)
		if nextNode == "end" {
			paths = append(paths, hereExplored)
			continue
		}
		paths = append(paths, findPaths(nextNode, hereExplored, linkMap)...)
	}
	return paths
}

func buildMap(pairs [][2]string) map[string][]string {
	linkMap := make(map[string][]string)
	for _, pair := range pairs {
		if _, ok := linkMap[pair[0]]; !ok {
			linkMap[pair[0]] = findLinks(pairs, pair[0])
		}
		if _, ok := linkMap[pair[1]]; !ok {
			linkMap[pair[1]] = findLinks(pairs, pair[1])
		}
	}
	return linkMap
}

func findLinks(pairs [][2]string, a string) []string {
	var links []string
	for _, pair := range pairs {
		if pair[0] == a {
			links = append(links, pair[1])
		}
		if pair[1] == a {
			links = append(links, pair[0])
		}
	}
	return links
}

func readData(fileName string) ([][2]string) {
	dataRaw, _ := os.ReadFile(fileName)
	dataStr := strings.Split(string(dataRaw), "\n")
	pairs := make([][2]string, len(dataStr))
	for i, str := range dataStr {
		pair := strings.Split(str, "-")
		pairs[i][0] = pair[0]
		pairs[i][1] = pair[1]
	}
	return pairs
}

func isIn(a string, b []string) bool {
	for _, c := range b {
		if a==c {
			return true
		}
	}
	return false
}

func IsLower(s string) bool {
    for _, r := range s {
        if !unicode.IsLower(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func assertEqual(a, b int) {
	if a != b {
		fmt.Println(fmt.Sprintf("Test failed, got %d", b))
		os.Exit(1)
	}
}
