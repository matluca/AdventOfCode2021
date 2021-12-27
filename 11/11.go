package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	testFile      = "test.txt"
	inputFile     = "input1.txt"
	solutionTest1 = 1656
	solutionTest2 = 195
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
	lights := readInput(fileName)
	flashes := 0
	for i := 0; i < 100; i++ {
		newLights, f := step(lights)
		lights = newLights
		flashes += f
	}
	return flashes
}

func solution2(fileName string) int {
	i := 1
	lights := readInput(fileName)
	for {
		newLights, f := step(lights)
		lights = newLights
		if f==100 {
			return i
		}
		i++
	}
	return 0
}

func step(in [10][10]int) ([10][10]int, int) {
	out := in
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			out[i][j]++
		}
	}
	var flashMap [10][10]bool
	for {
		newFlashMap, combinedFlashMap := buildFlashMap(out, flashMap)
		if !stillGoing(newFlashMap) {
			return resetFlashes(out)
		}
		out = flash(out, newFlashMap)
		flashMap = combinedFlashMap
	}
	return resetFlashes(out)
}

func resetFlashes(in [10][10]int) ([10][10]int, int) {
	f := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if in[i][j] > 9 {
				in[i][j] = 0
				f++
			}
		}
	}
	return in, f
}

func buildFlashMap(in [10][10]int, inFlashMap [10][10]bool) ([10][10]bool, [10][10]bool) {
	var newFlashMap [10][10]bool
	combinedFlashMap := inFlashMap
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if in[i][j] > 9 {
				combinedFlashMap[i][j] = true
				if !inFlashMap[i][j] {
					newFlashMap[i][j] = true
				}
			}
		}
	}
	return newFlashMap, combinedFlashMap
}

func flash(in [10][10]int, flashMap [10][10]bool) [10][10]int {
	out := in
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if flashMap[i][j] {
				if i > 0 {
					out[i-1][j]++
					if j > 0 {
						out[i-1][j-1]++
					}
					if j < 9 {
						out[i-1][j+1]++
					}
				}
				if i < 9 {
					out[i+1][j]++
					if j > 0 {
						out[i+1][j-1]++
					}
					if j < 9 {
						out[i+1][j+1]++
					}
				}
				if j > 0 {
					out[i][j-1]++
				}
				if j < 9 {
					out[i][j+1]++
				}
			}
		}
	}
	return out
}

func stillGoing(flashMap [10][10]bool) bool {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if flashMap[i][j] {
				return true
			}
		}
	}
	return false
}

func readInput(fileName string) [10][10]int {
	dataRaw, _ := os.ReadFile(fileName)
	rows := strings.Split(string(dataRaw), "\n")
	var output [10][10]int
	for i := 0; i < len(output); i++ {
		for j := 0; j < len(rows[i]); j++ {
			output[i][j] = int(rows[i][j] - '0')
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
