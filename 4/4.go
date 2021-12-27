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
	solutionTest1 = 4512
	solutionTest2 = 1924
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
	num, boards := readData(fileName)
	pos := positions(num)
	wins := make([]int, len(boards))
	for i := 0; i < len(boards); i++ {
		wins[i] = winRound(boards[i], pos)
	}
	win, _ := MinMax(wins)
	var winBoardIndex int
	for i := 0; i < len(wins); i++ {
		if wins[i] == win {
			winBoardIndex = i
			break
		}
	}
	winBoard := boards[winBoardIndex]
	score := boardScore(winBoard, pos, num[win])
	return score
}

func solution2(fileName string) int {
	num, boards := readData(fileName)
	pos := positions(num)
	wins := make([]int, len(boards))
	for i := 0; i < len(boards); i++ {
		wins[i] = winRound(boards[i], pos)
	}
	_, last := MinMax(wins)
	var lastBoardIndex int
	for i := 0; i < len(wins); i++ {
		if wins[i] == last {
			lastBoardIndex = i
			break
		}
	}
	lastBoard := boards[lastBoardIndex]
	score := boardScore(lastBoard, pos, num[last])
	return score
}

func readData(fileName string) ([]int, [][][]int) {
	dataRaw, _ := os.ReadFile(fileName)
	dataStr := strings.Split(string(dataRaw), "\n\n")
	numbersRaw := strings.Split(dataStr[0], ",")
	numbers := make([]int, len(numbersRaw))
	for i := 0; i < len(numbersRaw); i++ {
		numbers[i], _ = strconv.Atoi(numbersRaw[i])
	}
	boards := make([][][]int, len(dataStr)-1)
	for i := 0; i < len(boards); i++ {
		boards[i] = make([][]int, 5)
		for j := 0; j < 5; j++ {
			boards[i][j] = make([]int, 5)
		}
	}
	for b := 0; b < len(boards); b++ {
		boardRaw := dataStr[b+1]
		rows := strings.Split(boardRaw, "\n")
		for i := 0; i < 5; i++ {
			entries := strings.Split(rows[i], " ")
			for j := 0; j < 5; j++ {
				n, _ := strconv.Atoi(entries[j])
				boards[b][i][j] = n
			}
		}
	}
	return numbers, boards
}

func positions(numbers []int) map[int]int {
	pos := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		pos[numbers[i]] = i
	}
	return pos
}

func winRound(board [][]int, positions map[int]int) int {
	rows := make([]int, 5)
	cols := make([]int, 5)
	for r := 0; r < 5; r++ {
		_, rows[r] = MinMax(posBoard(board, positions)[r])
		_, cols[r] = MinMax(boardColumn(posBoard(board, positions), r))
	}
	minRows, _ := MinMax(rows)
	minCols, _ := MinMax(cols)
	if minRows < minCols {
		return minRows
	}
	return minCols
}

func posBoard(board [][]int, positions map[int]int) [][]int {
	posBoard := make([][]int, 5)
	for i := 0; i < 5; i++ {
		posBoard[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			posBoard[i][j] = positions[board[i][j]]
		}
	}
	return posBoard
}

func boardColumn(board [][]int, columnIndex int) (column []int) {
	column = make([]int, 0)
	for _, row := range board {
		column = append(column, row[columnIndex])
	}
	return
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func boardScore(board [][]int, positions map[int]int, number int) int {
	var sum int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if positions[board[i][j]] > positions[number] {
				sum = sum + board[i][j]
			}
		}
	}
	return sum * number
}

func assertEqual(a, b int) {
	if a != b {
		fmt.Println(fmt.Sprintf("Test failed, got %d", b))
		os.Exit(1)
	}
}
