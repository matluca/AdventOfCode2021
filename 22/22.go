package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	testFile      = "test.txt"
	testFile2     = "test2.txt"
	inputFile     = "input1.txt"
	solutionTest1 = 590784
	solutionTest2 = 2758514936282235
)

type step struct {
	action     int
	boundaries [6]int
}

type point struct {
	x, y, z int
}

func main() {
	test1 := solution1(testFile)
	fmt.Println(fmt.Sprintf("Part 1 test: %d", solutionTest1))
	assertEqual(solutionTest1, test1)
	sol1 := solution1(inputFile)
	fmt.Println(fmt.Sprintf("Part 1: %d", sol1))

	test2 := solution2(testFile2)
	fmt.Println(fmt.Sprintf("Part 2 test: %d", solutionTest2))
	assertEqual(solutionTest2, test2)
	sol2 := solution2(inputFile)
	fmt.Println(fmt.Sprintf("Part 2: %d", sol2))
}

func solution1(fileName string) int {
	steps := readInput(fileName)
	grid := make(map[point]bool)
	for x := -50; x < 51; x++ {
		for y := -50; y < 51; y++ {
			for z := -50; z < 51; z++ {
				grid[point{x: x, y: y, z: z}] = false
			}
		}
	}
	for _, s := range steps {
		action(grid, s)
	}
	return countOn(grid)
}

func solution2(fileName string) int {
	steps := readInput(fileName)
	var cubes []step
	for _, s := range steps {
		cubes = applyStep(cubes, s)
	}
	return signedVolume(cubes)
}

func signedVolume(cubes []step) int {
	var tot float64
	for _, c := range cubes {
		b := c.boundaries
		a := c.action
		tot += float64(b[1]-b[0]+1) * float64(b[3]-b[2]+1) * float64(b[5]-b[4]+1) * float64(a)
	}
	return int(tot)
}

func applyStep(cubes []step, s step) []step {
	sBoundaries := s.boundaries
	sAction := s.action
	var toAdd []step
	newCubes := cubes
	for _, c := range cubes {
		b := c.boundaries
		i0 := max(sBoundaries[0], b[0])
		i1 := min(sBoundaries[1], b[1])
		i2 := max(sBoundaries[2], b[2])
		i3 := min(sBoundaries[3], b[3])
		i4 := max(sBoundaries[4], b[4])
		i5 := min(sBoundaries[5], b[5])
		if i0 <= i1 && i2 <= i3 && i4 <= i5 {
			toAdd = append(toAdd, step{boundaries: [6]int{i0, i1, i2, i3, i4, i5}, action: -c.action})
		}
	}
	if sAction > 0 {
		newCubes = append(newCubes, step{boundaries: sBoundaries, action: sAction})
	}
	newCubes = append(newCubes, toAdd...)
	return newCubes
}

//func applyStep(cubes map[[6]int]int, s step) map[[6]int]int {
//	sBoundaries := s.boundaries
//	sAction := s.action
//	toAdd := make(map[[6]int]int)
//	newCubes := make(map[[6]int]int)
//	for k, v := range cubes {
//		newCubes[k] = v
//	}
//	for b, a := range cubes {
//		i0 := max(sBoundaries[0], b[0])
//		i1 := min(sBoundaries[1], b[1])
//		i2 := max(sBoundaries[2], b[2])
//		i3 := min(sBoundaries[3], b[3])
//		i4 := max(sBoundaries[4], b[4])
//		i5 := min(sBoundaries[5], b[5])
//		if i0 <= i1 && i2 <= i3 && i4 <= i5 {
//			toAdd[[6]int{i0, i1, i2, i3, i4, i5}] = -a
//		}
//	}
//	if sAction > 0 {
//		newCubes[sBoundaries] = sAction
//	}
//	for b, v := range toAdd {
//		if _, ok := newCubes[b]; !ok {
//			newCubes[b] = 0
//		}
//		newCubes[b] += v
//	}
//	return newCubes
//}

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

func countOn(grid map[point]bool) int {
	tot := 0
	for x := -50; x < 51; x++ {
		for y := -50; y < 51; y++ {
			for z := -50; z < 51; z++ {
				if grid[point{x: x, y: y, z: z}] {
					tot++
				}
			}
		}
	}
	return tot
}

func action(grid map[point]bool, s step) {
	if (s.boundaries[0] < -50) || (s.boundaries[2] < -50) || (s.boundaries[4] < -50) || (s.boundaries[1] > 50) || (s.boundaries[3] > 50) || (s.boundaries[5] > 50) {
		return
	}
	for x := s.boundaries[0]; x < s.boundaries[1]+1; x++ {
		for y := s.boundaries[2]; y < s.boundaries[3]+1; y++ {
			for z := s.boundaries[4]; z < s.boundaries[5]+1; z++ {
				p := point{x: x, y: y, z: z}
				switch s.action {
				case 1:
					grid[p] = true
				case -1:
					grid[p] = false
				}
			}
		}
	}
}

func readInput(fileName string) []step {
	dataRaw, _ := os.ReadFile(fileName)
	rows := strings.Split(string(dataRaw), "\n")
	steps := make([]step, len(rows))
	for i, row := range rows {
		spl := strings.Split(row, " ")
		var ac int
		if spl[0] == "on" {
			ac = 1
		} else {
			ac = -1
		}
		nRaw := strings.Split(spl[1], ",")
		var b [6]int
		for j := 0; j < 6; j++ {
			b[j], _ = strconv.Atoi(nRaw[j])
		}
		steps[i] = step{
			action:     ac,
			boundaries: b,
		}
	}
	return steps
}

func readInput2(fileName string) map[[6]int]int {
	dataRaw, _ := os.ReadFile(fileName)
	rows := strings.Split(string(dataRaw), "\n")
	steps := make(map[[6]int]int)
	for _, row := range rows {
		spl := strings.Split(row, " ")
		var ac int
		if spl[0] == "on" {
			ac = 1
		} else {
			ac = -1
		}
		nRaw := strings.Split(spl[1], ",")
		var b [6]int
		for j := 0; j < 6; j++ {
			b[j], _ = strconv.Atoi(nRaw[j])
		}
		steps[b] = ac
	}
	return steps
}

func assertEqual(a, b int) {
	if a != b {
		fmt.Println(fmt.Sprintf("Test failed, got %d", b))
		os.Exit(1)
	}
}
