package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	testFile      = "test.txt"
	inputFile     = "input1.txt"
	solutionTest1 = 58
)

func main() {
	test1 := solution1(testFile)
	fmt.Println(fmt.Sprintf("Part 1 test: %d", solutionTest1))
	assertEqual(solutionTest1, test1)
	sol1 := solution1(inputFile)
	fmt.Println(fmt.Sprintf("Part 1: %d", sol1))
}

func solution1(fileName string) int {
	grid := readInput(fileName)
	stepCounter := 0
	for {
		stepCounter++
		newGrid, counter := step(grid)
		grid = newGrid
		if counter == 0 {
			break
		}
	}
	return stepCounter
}

func step(in [][]string) ([][]string, int) {
	afterEast, c1 := stepEast(in)
	afterSouth, c2 := stepSouth(afterEast)
	return afterSouth, c1 + c2
}

func stepEast(in [][]string) ([][]string, int) {
	out := make([][]string, len(in))
	for i := 0; i < len(out); i++ {
		out[i] = make([]string, len(in[i]))
		for j := 0; j < len(out[i]); j++ {
			out[i][j] = "."
		}
	}
	moveCounter := 0
	// move east
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[i])-1; j++ {
			if in[i][j] == ">" {
				if in[i][j+1] == "." {
					out[i][j+1] = ">"
					moveCounter++
				} else {
					out[i][j] = ">"
				}
			}
		}
		if in[i][len(in[i])-1] == ">" {
			if in[i][0] == "." {
				out[i][0] = ">"
				moveCounter++
			} else {
				out[i][len(in[i])-1] = ">"
			}
		}
	}
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[i]); j++ {
			if in[i][j] == "v" {
				out[i][j] = "v"
			}
		}
	}
	return out, moveCounter
}

func stepSouth(in [][]string) ([][]string, int) {
	out := make([][]string, len(in))
	for i := 0; i < len(out); i++ {
		out[i] = make([]string, len(in[i]))
		for j := 0; j < len(out[i]); j++ {
			out[i][j] = "."
		}
	}
	moveCounter := 0
	// move south
	for i := 0; i < len(in)-1; i++ {
		for j := 0; j < len(in[i]); j++ {
			if in[i][j] == "v" {
				if in[i+1][j] == "." {
					out[i+1][j] = "v"
					moveCounter++
				} else {
					out[i][j] = "v"
				}
			}
		}
	}
	for j := 0; j < len(in[len(in)-1]); j++ {
		if in[len(in)-1][j] == "v" {
			if in[0][j] == "." {
				out[0][j] = "v"
				moveCounter++
			} else {
				out[len(in)-1][j] = "v"
			}
		}
	}
	for i := 0; i < len(in); i++ {
		for j := 0; j < len(in[i]); j++ {
			if in[i][j] == ">" {
				out[i][j] = ">"
			}
		}
	}
	return out, moveCounter
}

//func printGrid(grid [][]string) {
//	for i := 0; i < len(grid); i++ {
//		for j := 0; j < len(grid[i]); j++ {
//			fmt.Print(grid[i][j])
//		}
//		fmt.Println()
//	}
//}

func readInput(fileName string) [][]string {
	dataRaw, _ := os.ReadFile(fileName)
	rows := strings.Split(string(dataRaw), "\n")
	grid := make([][]string, len(rows))
	for i, row := range rows {
		grid[i] = make([]string, len(row))
		for j, c := range row {
			grid[i][j] = string(c)
		}
	}
	return grid
}

func assertEqual(a, b int) {
	if a != b {
		fmt.Println(fmt.Sprintf("Test failed, got %d", b))
		os.Exit(1)
	}
}
