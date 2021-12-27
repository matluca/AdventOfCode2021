package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	testFile  = "test.txt"
	inputFile = "input1.txt"
	testFile2 = "test2.txt"
)

var solutionTest1 = []int{16, 12, 23, 31}
var solutionTest2 = []int{3, 54, 7, 9, 1, 0, 0, 1}

var hexToBin = map[rune]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

func main() {
	test1 := solution1(testFile)
	fmt.Println(fmt.Sprintf("Part 1 tests: %v", solutionTest1))
	assertEqual(solutionTest1, test1)
	sol1 := solution1(inputFile)
	fmt.Println(fmt.Sprintf("Part 1: %d", sol1))

	test2 := solution2(testFile2)
	fmt.Println(fmt.Sprintf("Part 2 test: %d", solutionTest2))
	assertEqual(solutionTest2, test2)
	sol2 := solution2(inputFile)
	fmt.Println(fmt.Sprintf("Part 2: %d", sol2))
}

func solution1(fileName string) []int {
	packets := convertHexToBin(fileName)
	versions := make([]int, len(packets))
	for i, packet := range packets {
		v, _ := readPacketVersion(packet, 0, 1, 20000)
		versions[i] = v
	}
	return versions
}

func solution2(fileName string) []int {
	packets := convertHexToBin(fileName)
	values := make([]int, len(packets))
	for i, packet := range packets {
		v, _ := readPacketValue(packet, 0, 1, 20000)
		values[i] = v[0]
	}
	return values
}

func readPacketVersion(packet string, startIndex, packetNumber, packetLength int) (int, int) {
	index := startIndex
	totVersion := 0
	for p := 0; p < packetNumber; p++ {
		if index-startIndex >= packetLength {
			break
		}
		totVersion += binToDecimal(packet[index : index+3])
		typeID := packet[index+3 : index+6]
		if typeID == "100" {
			number := ""
			i := 0
			for {
				number += packet[index+7+5*i : index+11+5*i]
				if string(packet[index+6+5*i]) == "0" {
					index = index + 11 + 5*i
					break
				}
				i++
			}
		} else {
			lengthTypeID := string(packet[index+6])
			if lengthTypeID == "0" {
				subpacketLengthString := packet[index+7 : index+22]
				subPacketLength := binToDecimal(subpacketLengthString)
				spv, _ := readPacketVersion(packet, index+22, 20000, subPacketLength)
				totVersion += spv
				index = index + 22 + subPacketLength
			} else {
				subpacketNumberString := packet[index+7 : index+18]
				subpacketNumber := binToDecimal(subpacketNumberString)
				spv, endIndex := readPacketVersion(packet, index+18, subpacketNumber, 20000)
				totVersion += spv
				index = endIndex
			}
		}
	}
	return totVersion, index
}

func readPacketValue(packet string, startIndex, packetNumber, packetLength int) ([]int, int) {
	index := startIndex
	var returnValues []int
	for p := 0; p < packetNumber; p++ {
		if index-startIndex >= packetLength {
			break
		}
		typeID := packet[index+3 : index+6]
		if typeID == "100" {
			number := ""
			i := 0
			for {
				number += packet[index+7+5*i : index+11+5*i]
				if string(packet[index+6+5*i]) == "0" {
					index = index + 11 + 5*i
					break
				}
				i++
			}
			returnValues = append(returnValues, binToDecimal(number))
		} else {
			lengthTypeID := string(packet[index+6])
			var subPacketValues []int
			if lengthTypeID == "0" {
				subpacketLengthString := packet[index+7 : index+22]
				subPacketLength := binToDecimal(subpacketLengthString)
				subPacketValues, _ = readPacketValue(packet, index+22, 20000, subPacketLength)
				index = index + 22 + subPacketLength
			} else {
				subpacketNumberString := packet[index+7 : index+18]
				subpacketNumber := binToDecimal(subpacketNumberString)
				spv, endIndex := readPacketValue(packet, index+18, subpacketNumber, 20000)
				subPacketValues = spv
				index = endIndex
			}
			switch typeID {
			case "000":
				returnValues = append(returnValues, sum(subPacketValues))
			case "001":
				returnValues = append(returnValues, product(subPacketValues))
			case "010":
				returnValues = append(returnValues, min(subPacketValues))
			case "011":
				returnValues = append(returnValues, max(subPacketValues))
			case "101":
				if subPacketValues[0] > subPacketValues[1] {
					returnValues = append(returnValues, 1)
				} else {
					returnValues = append(returnValues, 0)
				}
			case "110":
				if subPacketValues[0] < subPacketValues[1] {
					returnValues = append(returnValues, 1)
				} else {
					returnValues = append(returnValues, 0)
				}
			case "111":
				if subPacketValues[0] == subPacketValues[1] {
					returnValues = append(returnValues, 1)
				} else {
					returnValues = append(returnValues, 0)
				}
			}
		}
	}
	return returnValues, index
}

func sum(in []int) int {
	s := 0
	for _, n := range in {
		s += n
	}
	return s
}

func product(in []int) int {
	s := 1
	for _, n := range in {
		s *= n
	}
	return s
}

func min(in []int) int {
	m := in[0]
	for _, n := range in {
		if n < m {
			m = n
		}
	}
	return m
}

func max(in []int) int {
	m := in[0]
	for _, n := range in {
		if n > m {
			m = n
		}
	}
	return m
}

func binToDecimal(bin string) int {
	n := 0
	for i := len(bin) - 1; i >= 0; i-- {
		if bin[i] == '1' {
			n = n + int(math.Pow(2, float64(len(bin)-1-i)))
		}
	}
	return n
}

func convertHexToBin(fileName string) []string {
	dataRaw, _ := os.ReadFile(fileName)
	rows := strings.Split(string(dataRaw), "\n")
	bins := make([]string, len(rows))
	for i, row := range rows {
		for _, c := range row {
			bins[i] += hexToBin[c]
		}
	}

	return bins
}

func assertEqual(a, b []int) {
	equal := true
	for i, ai := range a {
		if ai != b[i] {
			equal = false
		}
	}
	if !equal {
		fmt.Println(fmt.Sprintf("Test failed, got %v", b))
		os.Exit(1)
	}
}
