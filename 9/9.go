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
	solutionTest1 = 15
	solutionTest2 = 1134
)

type point struct {
	x int
	y int
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
	if (input[0][0] < input[0][1]) && (input[0][0] < input[1][0]) {
		tot += input[0][0] + 1
	}
	lenCol := len(input)
	lenRow := len(input[0])
	if (input[0][lenRow-1] < input[0][lenRow-2]) && (input[0][lenRow-1] < input[1][lenRow-1]) {
		tot += input[0][lenRow-1] + 1
	}
	if (input[lenCol-1][0] < input[lenCol-2][0]) && (input[lenCol-1][0] < input[lenCol-1][1]) {
		tot += input[lenCol-1][0] + 1
	}
	if (input[lenCol-1][lenRow-1] < input[lenCol-2][lenRow-1]) && (input[lenCol-1][lenRow-1] < input[lenCol-1][lenRow-2]) {
		tot += input[lenCol-1][lenRow-1] + 1
	}
	for j := 1; j < lenRow-1; j++ {
		if (input[0][j] < input[0][j-1]) && (input[0][j] < input[0][j+1]) && (input[0][j] < input[1][j]) {
			tot += input[0][j] + 1
		}
	}
	for j := 1; j < lenRow-1; j++ {
		if (input[lenCol-1][j] < input[lenCol-1][j-1]) && (input[lenCol-1][j] < input[lenCol-1][j+1]) && (input[lenCol-1][j] < input[lenCol-2][j]) {
			tot += input[lenCol-1][j] + 1
		}
	}
	for i := 1; i < lenCol-1; i++ {
		if (input[i][0] < input[i+1][0]) && (input[i][0] < input[i-1][0]) && (input[i][0] < input[i][1]) {
			tot += input[i][0] + 1
		}
	}
	for i := 1; i < lenCol-1; i++ {
		if (input[i][lenRow-1] < input[i+1][lenRow-1]) && (input[i][lenRow-1] < input[i-1][lenRow-1]) && (input[i][lenRow-1] < input[i][lenRow-2]) {
			tot += input[i][lenRow-1] + 1
		}
	}
	for i := 1; i < lenCol-1; i++ {
		for j := 1; j < lenRow-1; j++ {
			if (input[i][j] < input[i-1][j]) && (input[i][j] < input[i+1][j]) && (input[i][j] < input[i][j-1]) && (input[i][j] < input[i][j+1]) {
				tot += input[i][j] + 1
			}
		}
	}
	return tot
}

func solution2(fileName string) int {
	input := readInput(fileName)
	minima := findMinima(input)
	replaced := replaceMinima(input, minima)
	for i := 0; i < max(len(input), len(input[0])); i++ {
		replaced = step(replaced)
	}
	sizes := make([]int, len(minima))
	for i := 0; i < len(minima); i++ {
		sizes[i] = size(replaced, i)
	}
	sort.Ints(sizes)
	length := len(sizes) - 1
	return sizes[length] * sizes[length-1] * sizes[length-2]
}

func size(d [][]string, n int) int {
	count := 0
	nString := fmt.Sprintf("%d", n)
	for i := 0; i < len(d); i++ {
		for j := 0; j < len(d[i]); j++ {
			if d[i][j] == nString {
				count++
			}
		}
	}
	return count
}

func step(d [][]string) [][]string {
	o := d
	for i := 0; i < len(d); i++ {
		for j := 0; j < len(d[0]); j++ {
			if d[i][j] == "." {
				if d[max(0, i-1)][j] != "." && d[max(0, i-1)][j] != "X" {
					o[i][j] = d[max(0, i-1)][j]
				}
				if d[i][max(0, j-1)] != "." && d[i][max(0, j-1)] != "X" {
					o[i][j] = d[i][max(0, j-1)]
				}
				if d[min(len(d)-1, i+1)][j] != "." && d[min(len(d)-1, i+1)][j] != "X" {
					o[i][j] = d[min(len(d)-1, i+1)][j]
				}
				if d[i][min(len(d[0])-1, j+1)] != "." && d[i][min(len(d[0])-1, j+1)] != "X" {
					o[i][j] = d[i][min(len(d[0])-1, j+1)]
				}
			}
		}
	}
	return o
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printDasboard(d [][]string) {
	for i := 0; i < len(d); i++ {
		for j := 0; j < len(d[i]); j++ {
			fmt.Print(d[i][j])
		}
		fmt.Println()
	}
}

func replaceMinima(input [][]int, minima []point) [][]string {
	output := make([][]string, len(input))
	for i := 0; i < len(output); i++ {
		output[i] = make([]string, len(input[i]))
	}
	for i, p := range minima {
		output[p.x][p.y] = fmt.Sprintf("%d", i)
	}
	for i := 0; i < len(output); i++ {
		for j := 0; j < len(output[0]); j++ {
			if input[i][j] == 9 {
				output[i][j] = "X"
			}
			if output[i][j] == "" {
				output[i][j] = "."
			}
		}
	}
	return output
}

func findMinima(input [][]int) []point {
	var points []point
	if (input[0][0] < input[0][1]) && (input[0][0] < input[1][0]) {
		points = append(points, point{0, 0})
	}
	lenCol := len(input)
	lenRow := len(input[0])
	if (input[0][lenRow-1] < input[0][lenRow-2]) && (input[0][lenRow-1] < input[1][lenRow-1]) {
		points = append(points, point{0, lenRow - 1})
	}
	if (input[lenCol-1][0] < input[lenCol-2][0]) && (input[lenCol-1][0] < input[lenCol-1][1]) {
		points = append(points, point{lenCol - 1, 0})
	}
	if (input[lenCol-1][lenRow-1] < input[lenCol-2][lenRow-1]) && (input[lenCol-1][lenRow-1] < input[lenCol-1][lenRow-2]) {
		points = append(points, point{lenCol - 1, lenRow - 1})
	}
	for j := 1; j < lenRow-1; j++ {
		if (input[0][j] < input[0][j-1]) && (input[0][j] < input[0][j+1]) && (input[0][j] < input[1][j]) {
			points = append(points, point{0, j})
		}
	}
	for j := 1; j < lenRow-1; j++ {
		if (input[lenCol-1][j] < input[lenCol-1][j-1]) && (input[lenCol-1][j] < input[lenCol-1][j+1]) && (input[lenCol-1][j] < input[lenCol-2][j]) {
			points = append(points, point{lenCol - 1, j})
		}
	}
	for i := 1; i < lenCol-1; i++ {
		if (input[i][0] < input[i+1][0]) && (input[i][0] < input[i-1][0]) && (input[i][0] < input[i][1]) {
			points = append(points, point{i, 0})
		}
	}
	for i := 1; i < lenCol-1; i++ {
		if (input[i][lenRow-1] < input[i+1][lenRow-1]) && (input[i][lenRow-1] < input[i-1][lenRow-1]) && (input[i][lenRow-1] < input[i][lenRow-2]) {
			points = append(points, point{i, lenRow - 1})
		}
	}
	for i := 1; i < lenCol-1; i++ {
		for j := 1; j < lenRow-1; j++ {
			if (input[i][j] < input[i-1][j]) && (input[i][j] < input[i+1][j]) && (input[i][j] < input[i][j-1]) && (input[i][j] < input[i][j+1]) {
				points = append(points, point{i, j})
			}
		}
	}
	return points
}

func readInput(fileName string) [][]int {
	dataRaw, _ := os.ReadFile(fileName)
	rows := strings.Split(string(dataRaw), "\n")
	output := make([][]int, len(rows))
	for i := 0; i < len(output); i++ {
		output[i] = make([]int, len(rows[i]))
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
