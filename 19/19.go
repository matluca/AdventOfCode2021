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
	solutionTest1 = 79
	solutionTest2 = 3621
)

type point struct {
	x int
	y int
	z int
}

type vector struct {
	x int
	y int
	z int
}

type scanner struct {
	id      int
	pos     point
	beacons []point
}

func add(a point, v vector) point {
	return point{
		x: a.x + v.x,
		y: a.y + v.y,
		z: a.z + v.z,
	}
}

func sub(a, b point) vector {
	return vector{
		x: a.x - b.x,
		y: a.y - b.y,
		z: a.z - b.z,
	}
}

type scannerQueueElement struct {
	sc   *scanner
	next *scannerQueueElement
}

type scannerQueue struct {
	first *scannerQueueElement
	last  *scannerQueueElement
}

func (scq *scannerQueue) append(sc *scanner) {
	newLast := &scannerQueueElement{sc: sc}
	scq.last.next = newLast
	scq.last = newLast
}

func main() {
	test1, test2 := solution(testFile)
	fmt.Println(fmt.Sprintf("Part 1 test: %v", solutionTest1))
	fmt.Println(fmt.Sprintf("Part 2 test: %v", solutionTest2))
	assertEqual(solutionTest1, test1)
	assertEqual(solutionTest2, test2)
	sol1, sol2 := solution(inputFile)
	fmt.Println(fmt.Sprintf("Part 1: %d", sol1))
	fmt.Println(fmt.Sprintf("Part 2: %d", sol2))
}

func solution(fileName string) (int, int) {
	scanners := readInput(fileName)
	sc0 := copy(scanners[0])
	sc0Element := scannerQueueElement{sc: &sc0}
	queue := scannerQueue{first: &sc0Element, last: &sc0Element}
	currentScanner := queue.first
	found := make(map[int]bool)
	found[0] = true
	for len(found) < len(scanners) {
		for j := 0; j < len(scanners); j++ {
			if _, ok := found[j]; ok {
				continue
			}
			scB := copy(scanners[j])
			for _, iso := range isometries {
				transformedScanner := apply(scB, iso)
				match, _, newScanner := haveNInCommonWithTranslation(*currentScanner.sc, transformedScanner, 12)
				if match {
					found[j] = true
					queue.append(&newScanner)
					continue
				}
			}
		}
		currentScanner = currentScanner.next
	}

	foundBeacons := make(map[point]bool)
	currentScanner = queue.first
	for currentScanner != nil {
		for _, b := range currentScanner.sc.beacons {
			foundBeacons[b] = true
		}
		currentScanner = currentScanner.next
	}

	maxDistance := 0
	base := queue.first
	for base != nil {
		target := base.next
		for target != nil {
			d := manhattanDistance(base.sc.pos, target.sc.pos)
			if d > maxDistance {
				maxDistance = d
			}
			target = target.next
		}
		base = base.next
	}
	return len(foundBeacons), maxDistance
}

func manhattanDistance(a, b point) int {
	d := 0
	if a.x > b.x {
		d += a.x - b.x
	} else {
		d += b.x - a.x
	}
	if a.y > b.y {
		d += a.y - b.y
	} else {
		d += b.y - a.y
	}
	if a.z > b.z {
		d += a.z - b.z
	} else {
		d += b.z - a.z
	}
	return d
}

// checks if two scanners have at least n beacons in common up to translation of the second scanner.
// Also translates the second scanner to match the first
func haveNInCommonWithTranslation(s1, s2 scanner, n int) (bool, vector, scanner) {
	for _, b1 := range s1.beacons {
		for _, b2 := range s2.beacons {
			v := sub(b1, b2)
			movedScanner2 := translate(s2, v)
			if haveNInCommon(s1, movedScanner2, n) {
				return true, v, movedScanner2
			}
		}
	}
	return false, vector{}, s2
}

func translate(a scanner, v vector) scanner {
	pos := add(a.pos, v)
	b := make([]point, len(a.beacons))
	for i := 0; i < len(a.beacons); i++ {
		b[i] = add(a.beacons[i], v)
	}
	return scanner{
		id:      a.id,
		pos:     pos,
		beacons: b,
	}
}

func copy(a scanner) scanner {
	pos := a.pos
	b := make([]point, len(a.beacons))
	for i := 0; i < len(a.beacons); i++ {
		b[i] = a.beacons[i]
	}
	return scanner{
		id:      a.id,
		pos:     pos,
		beacons: b,
	}
}

// checks if two scanners have at least n beacons in common
func haveNInCommon(s1, s2 scanner, n int) bool {
	totCommon := 0
	for _, b1 := range s1.beacons {
		for _, b2 := range s2.beacons {
			if b1 == b2 {
				totCommon++
			}
		}
	}
	return totCommon >= n
}

func readInput(fileName string) []scanner {
	dataRaw, _ := os.ReadFile(fileName)
	scannersRaw := strings.Split(string(dataRaw), "\n\n")
	scanners := make([]scanner, len(scannersRaw))
	for i, scannerRaw := range scannersRaw {
		beaconsRaw := strings.Split(scannerRaw, "\n")
		beacons := make([]point, len(beaconsRaw)-1)
		for j := 1; j < len(beaconsRaw); j++ {
			coords := strings.Split(beaconsRaw[j], ",")
			x, _ := strconv.Atoi(coords[0])
			y, _ := strconv.Atoi(coords[1])
			z, _ := strconv.Atoi(coords[2])
			beacons[j-1] = point{x: x, y: y, z: z}
		}
		scanners[i] = scanner{
			id:      i,
			beacons: beacons,
		}
	}
	return scanners
}

func assertEqual(a, b int) {
	if a != b {
		fmt.Println(fmt.Sprintf("Test failed, got %v", b))
		os.Exit(1)
	}
}

type isometry [3][3]int

func apply(sc scanner, iso isometry) scanner {
	s := copy(sc)
	pos := s.pos
	b := make([]point, len(s.beacons))
	for i := 0; i < len(s.beacons); i++ {
		b[i] = matrixMult(iso, s.beacons[i])
	}
	return scanner{
		id:      s.id,
		pos:     pos,
		beacons: b,
	}
}

func matrixMult(iso isometry, p point) point {
	in := [3]int{p.x, p.y, p.z}
	out := [3]int{0, 0, 0}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out[i] += iso[i][j] * in[j]
		}
	}
	return point{x: out[0], y: out[1], z: out[2]}
}

var isometries = []isometry{
	[3][3]int{ //0
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	},
	[3][3]int{ //1
		{-1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	},
	[3][3]int{ //2
		{1, 0, 0},
		{0, -1, 0},
		{0, 0, 1},
	},
	[3][3]int{ //3
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, -1},
	},
	[3][3]int{ //4
		{-1, 0, 0},
		{0, -1, 0},
		{0, 0, 1},
	},
	[3][3]int{ //5
		{-1, 0, 0},
		{0, 1, 0},
		{0, 0, -1},
	},
	[3][3]int{ //6
		{1, 0, 0},
		{0, -1, 0},
		{0, 0, -1},
	},
	[3][3]int{ //7
		{-1, 0, 0},
		{0, -1, 0},
		{0, 0, -1},
	},
	[3][3]int{ //8
		{1, 0, 0},
		{0, 0, 1},
		{0, 1, 0},
	},
	[3][3]int{ //9
		{-1, 0, 0},
		{0, 0, 1},
		{0, 1, 0},
	},
	[3][3]int{ //10
		{1, 0, 0},
		{0, 0, 1},
		{0, -1, 0},
	},
	[3][3]int{ //11
		{1, 0, 0},
		{0, 0, -1},
		{0, 1, 0},
	},
	[3][3]int{ //12
		{-1, 0, 0},
		{0, 0, 1},
		{0, -1, 0},
	},
	[3][3]int{ //13
		{-1, 0, 0},
		{0, 0, -1},
		{0, 1, 0},
	},
	[3][3]int{ //14
		{1, 0, 0},
		{0, 0, -1},
		{0, -1, 0},
	},
	[3][3]int{ //15
		{-1, 0, 0},
		{0, 0, -1},
		{0, -1, 0},
	},
	[3][3]int{ //16
		{0, 0, 1},
		{1, 0, 0},
		{0, 1, 0},
	},
	[3][3]int{ //17
		{0, 0, 1},
		{-1, 0, 0},
		{0, 1, 0},
	},
	[3][3]int{ //18
		{0, 0, 1},
		{1, 0, 0},
		{0, -1, 0},
	},
	[3][3]int{ //19
		{0, 0, -1},
		{1, 0, 0},
		{0, 1, 0},
	},
	[3][3]int{ //20
		{0, 0, 1},
		{-1, 0, 0},
		{0, -1, 0},
	},
	[3][3]int{ //21
		{0, 0, -1},
		{-1, 0, 0},
		{0, 1, 0},
	},
	[3][3]int{ //22
		{0, 0, -1},
		{1, 0, 0},
		{0, -1, 0},
	},
	[3][3]int{ //23
		{0, 0, -1},
		{-1, 0, 0},
		{0, -1, 0},
	},
	[3][3]int{ //24
		{0, 1, 0},
		{1, 0, 0},
		{0, 0, 1},
	},
	[3][3]int{ //25
		{0, 1, 0},
		{-1, 0, 0},
		{0, 0, 1},
	},
	[3][3]int{ //26
		{0, -1, 0},
		{1, 0, 0},
		{0, 0, 1},
	},
	[3][3]int{ //27
		{0, 1, 0},
		{1, 0, 0},
		{0, 0, -1},
	},
	[3][3]int{ //28
		{0, -1, 0},
		{-1, 0, 0},
		{0, 0, 1},
	},
	[3][3]int{ //29
		{0, 1, 0},
		{-1, 0, 0},
		{0, 0, -1},
	},
	[3][3]int{ //30
		{0, -1, 0},
		{1, 0, 0},
		{0, 0, -1},
	},
	[3][3]int{ //31
		{0, -1, 0},
		{-1, 0, 0},
		{0, 0, -1},
	},
	[3][3]int{ //32
		{0, 0, 1},
		{0, 1, 0},
		{1, 0, 0},
	},
	[3][3]int{ //33
		{0, 0, 1},
		{0, 1, 0},
		{-1, 0, 0},
	},
	[3][3]int{ //34
		{0, 0, 1},
		{0, -1, 0},
		{1, 0, 0},
	},
	[3][3]int{ //35
		{0, 0, -1},
		{0, 1, 0},
		{1, 0, 0},
	},
	[3][3]int{ //36
		{0, 0, 1},
		{0, -1, 0},
		{-1, 0, 0},
	},
	[3][3]int{ //37
		{0, 0, -1},
		{0, 1, 0},
		{-1, 0, 0},
	},
	[3][3]int{ //38
		{0, 0, -1},
		{0, -1, 0},
		{1, 0, 0},
	},
	[3][3]int{ //39
		{0, 0, -1},
		{0, -1, 0},
		{-1, 0, 0},
	},
	[3][3]int{ //40
		{0, 1, 0},
		{0, 0, 1},
		{1, 0, 0},
	},
	[3][3]int{ //41
		{0, 1, 0},
		{0, 0, 1},
		{-1, 0, 0},
	},
	[3][3]int{ //42
		{0, -1, 0},
		{0, 0, 1},
		{1, 0, 0},
	},
	[3][3]int{ //43
		{0, 1, 0},
		{0, 0, -1},
		{1, 0, 0},
	},
	[3][3]int{ //44
		{0, -1, 0},
		{0, 0, 1},
		{-1, 0, 0},
	},
	[3][3]int{ //45
		{0, 1, 0},
		{0, 0, -1},
		{-1, 0, 0},
	},
	[3][3]int{ //46
		{0, -1, 0},
		{0, 0, -1},
		{1, 0, 0},
	},
	[3][3]int{ //47
		{0, -1, 0},
		{0, 0, -1},
		{-1, 0, 0},
	},
}
