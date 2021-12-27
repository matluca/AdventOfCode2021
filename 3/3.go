package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	testFile      = "test.txt"
	inputFile     = "input1.txt"
	solutionTest1 = 198
	solutionTest2 = 230
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
	dataRaw, _ := os.ReadFile(fileName)
	dataStr := strings.Split(string(dataRaw), "\n")
	lenRow := len(dataStr[0])
	n0 := make([]int, lenRow)
	n1 := make([]int, lenRow)
	for i := 0; i < len(dataStr); i++ {
		for j := 0; j < lenRow; j++ {
			if dataStr[i][j] == '0' {
				n0[j]++
			} else {
				n1[j]++
			}
		}
	}
	common := make([]int, lenRow)
	uncommon := make([]int, lenRow)
	for j := 0; j < lenRow; j++ {
		if n0[j] > n1[j] {
			uncommon[j] = 1
		} else {
			common[j] = 1
		}
	}
	return toDecimal(common) * toDecimal(uncommon)
}

func solution2(fileName string) int {
	dataRaw, _ := os.ReadFile(fileName)
	dataStr := strings.Split(string(dataRaw), "\n")
	dataCopy := dataStr
	idx := 0
	for {
		mc := mostCommon(dataCopy, idx)
		var dataNew []string
		for i := 0; i < len(dataCopy); i++ {
			if dataCopy[i][idx] == mc {
				dataNew = append(dataNew, dataCopy[i])
			}
		}
		dataCopy = dataNew
		if len(dataCopy) == 1 {
			break
		}
		idx++
	}
	oxygen := dataCopy[0]
	dataCopy = dataStr
	idx = 0
	for {
		lc := leastCommon(dataCopy, idx)
		var dataNew []string
		for i := 0; i < len(dataCopy); i++ {
			if dataCopy[i][idx] == lc {
				dataNew = append(dataNew, dataCopy[i])
			}
		}
		dataCopy = dataNew
		if len(dataCopy) == 1 {
			break
		}
		idx++
	}
	co2 := dataCopy[0]
	return toDecimalString(oxygen)*toDecimalString(co2)
}

func toDecimal(bin []int) int {
	var n int
	for i := len(bin) - 1; i >= 0; i-- {
		n = n + bin[i]*int(math.Pow(2, float64(len(bin)-i-1)))
	}
	return n
}

func toDecimalString(bin string) int {
	var n int
	for i := len(bin) - 1; i >= 0; i-- {
		if bin[i] == '1' {
			n = n + int(math.Pow(2, float64(len(bin)-i-1)))
		}
	}
	return n
}

func mostCommon(data []string, idx int) uint8 {
	var n0, n1 int
	for i := 0; i < len(data); i++ {
		if data[i][idx] == '0' {
			n0++
		} else {
			n1++
		}
	}
	if n0 > n1 {
		return '0'
	}
	return '1'
}

func leastCommon(data []string, idx int) uint8 {
	var n0, n1 int
	for i := 0; i < len(data); i++ {
		if data[i][idx] == '0' {
			n0++
		} else {
			n1++
		}
	}
	if n0 > n1 {
		return '1'
	}
	return '0'
}

func assertEqual(a, b int) {
	if a != b {
		fmt.Println(fmt.Sprintf("Test failed, got %d", b))
		os.Exit(1)
	}
}
