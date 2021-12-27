package main

import (
	"fmt"
	"os"
)

const (
	solutionTest1 = 739785
	solutionTest2 = 444356092776315
)

var testInput = [2]int{4, 8}
var input = [2]int{10, 1}

type status struct {
	pos   int
	score int
}

func main() {
	test1 := solution1(testInput)
	fmt.Println(fmt.Sprintf("Part 1 test: %d", solutionTest1))
	assertEqual(solutionTest1, uint64(test1))
	sol1 := solution1(input)
	fmt.Println(fmt.Sprintf("Part 1: %d", sol1))

	test2 := solution2(testInput)
	fmt.Println(fmt.Sprintf("Part 2 test: %d", solutionTest2))
	assertEqual(solutionTest2, test2)
	sol2 := solution2(input)
	fmt.Println(fmt.Sprintf("Part 2: %d", sol2))
}

func solution2(input [2]int) uint64 {
	var universes [2]float64
	statusMap0 := map[status]float64{
		status{
			pos:   input[0],
			score: 0,
		}: 1,
	}
	statusMap1 := map[status]float64{
		status{
			pos:   input[1],
			score: 0,
		}: 1,
	}
	statusMap := [2]map[status]float64{statusMap0, statusMap1}
	turn := 0
	for i := 0; i < 30; i++ {
		statusMap[turn] = statusMapStep(statusMap[turn])
		newStatusMap, winUniverses, notWinUniverses := checkWins(statusMap[turn])
		statusMap[turn] = newStatusMap
		universes[turn] += winUniverses
		if notWinUniverses == 0 {
			break
		}
		statusMap[(turn+1)%2] = updateOtherPlayer(statusMap[(turn+1)%2], winUniverses, notWinUniverses)
		turn = (turn + 1) % 2
	}
	if universes[turn] > universes[(turn+1)%2] {
		return uint64(universes[turn])
	}
	return uint64(universes[(turn+1)%2])
}

func updateOtherPlayer(this map[status]float64, winUniverses, notWinUniverses float64) map[status]float64 {
	newStatusMap := make(map[status]float64)
	for key, value := range this {
		newStatusMap[key] = value * 27 * notWinUniverses / (winUniverses + notWinUniverses)
	}
	return newStatusMap
}

func checkWins(statusMap map[status]float64) (map[status]float64, float64, float64) {
	var universesWin, universesNotWin float64
	newStatusMap := make(map[status]float64)
	for key, value := range statusMap {
		if key.score >= 21 {
			universesWin += value
		} else {
			newStatusMap[key] = value
			universesNotWin += value
		}
	}
	return newStatusMap, universesWin, universesNotWin
}

func statusMapStep(statusMap map[status]float64) map[status]float64 {
	newStatusMap := make(map[status]float64)
	for d1 := 1; d1 < 4; d1++ {
		for d2 := 1; d2 < 4; d2++ {
			for d3 := 1; d3 < 4; d3++ {
				for key, value := range statusMap {
					newStatus := nextStatus(key, d1+d2+d3)
					if _, ok := newStatusMap[newStatus]; !ok {
						newStatusMap[newStatus] = 0
					}
					newStatusMap[newStatus] += value
				}
			}
		}
	}
	return newStatusMap
}

func nextStatus(st status, d int) status {
	return status{
		pos:   move(st.pos, d),
		score: st.score + move(st.pos, d),
	}
}

func solution1(input [2]int) int {
	die := 0
	var score [2]int
	pos := input
	turn := 0
	rolls := 0
	for max(score) < 1000 {
		sum := 3*die + 6
		pos[turn] = move(pos[turn], sum)
		score[turn] += pos[turn]
		die = (die+2)%100 + 1
		rolls += 3
		turn = (turn + 1) % 2
	}
	return rolls * min(score)
}

func move(pos, sum int) int {
	return (pos-1+sum)%10 + 1
}

func max(in [2]int) int {
	if in[0] > in[1] {
		return in[0]
	}
	return in[1]
}

func min(in [2]int) int {
	if in[0] < in[1] {
		return in[0]
	}
	return in[1]
}

func assertEqual(a, b uint64) {
	if a != b {
		fmt.Println(fmt.Sprintf("Test failed, got %d", b))
		os.Exit(1)
	}
}
