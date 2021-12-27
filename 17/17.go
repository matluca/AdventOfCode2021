package main

import (
	"fmt"
	"os"
)

const (
	solutionTest1 = 45
	solutionTest2 = 112
)

type target struct {
	xmin int
	ymin int
	xmax int
	ymax int
}

type status struct {
	x  int
	y  int
	vx int
	vy int
}

var testTarget = target{
	xmin: 20,
	ymin: -10,
	xmax: 30,
	ymax: -5,
}

var realTarget = target{
	xmin: 25,
	ymin: -260,
	xmax: 67,
	ymax: -200,
}

func main() {
	test1 := solution1(testTarget)
	fmt.Println(fmt.Sprintf("Part 1 test: %v", solutionTest1))
	assertEqual(solutionTest1, test1)
	sol1 := solution1(realTarget)
	fmt.Println(fmt.Sprintf("Part 1: %d", sol1))

	test2 := solution2(testTarget)
	fmt.Println(fmt.Sprintf("Part 2 test: %d", solutionTest2))
	assertEqual(solutionTest2, test2)
	sol2 := solution2(realTarget)
	fmt.Println(fmt.Sprintf("Part 2: %d", sol2))
}

func solution1(t target) int {
	max := 0
	for vx := 1; vx < 1000; vx++ {
		for vy := 1; vy < 1000; vy++ {
			st := status{
				x:  0,
				y:  0,
				vx: vx,
				vy: vy,
			}
			hmax := 0
			for {
				st = step(st)
				if st.y > hmax {
					hmax = st.y
				}
				if checkIn(st, t) {
					if hmax > max {
						max = hmax
					}
					break
				}
				if checkOvershoot(st, t) {
					break
				}
			}
		}
	}
	return max
}

func solution2(t target) int {
	tot := 0
	for vx := 1; vx < 1000; vx++ {
		for vy := -1000; vy < 1000; vy++ {
			st := status{
				x:  0,
				y:  0,
				vx: vx,
				vy: vy,
			}
			for {
				st = step(st)
				if checkIn(st, t) {
					tot++
					break
				}
				if checkOvershoot(st, t) {
					break
				}
			}
		}
	}
	return tot
}

func step(in status) status {
	out := status{}
	out.x = in.x + in.vx
	out.y = in.y + in.vy
	if in.vx > 0 {
		out.vx = in.vx - 1
	}
	out.vy = in.vy - 1
	return out
}

func checkIn(st status, tg target) bool {
	if st.x <= tg.xmax && st.x >= tg.xmin && st.y >= tg.ymin && st.y <= tg.ymax {
		return true
	}
	return false
}

func checkOvershoot(st status, tg target) bool {
	if st.x > tg.xmax {
		return true
	}
	if st.y < tg.ymin {
		return true
	}
	return false
}

func assertEqual(a, b int) {
	if a != b {
		fmt.Println(fmt.Sprintf("Test failed, got %v", b))
		os.Exit(1)
	}
}
