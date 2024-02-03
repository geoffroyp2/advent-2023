package day06

import (
	"advent-2023/src/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Dir string

const (
	U Dir = "U"
	R Dir = "R"
	D Dir = "D"
	L Dir = "L"
)

type Color struct {
	r, g, b int
}

type Step struct {
	dir Dir
	len int64
}

type Boundaries struct {
	w, h, x, y int
}

type Coord struct {
	x, y int64
}

func getSteps1(lines []string) []Step {
	steps := make([]Step, 0)

	for idx := range lines {
		valStr := strings.Split(lines[idx], " ")
		dir := Dir(valStr[0])
		if dir != R && dir != D && dir != L && dir != U {
			panic("Invalid direction")
		}
		dist, err := strconv.Atoi(valStr[1])
		if err != nil {
			panic(err)
		}
		steps = append(steps, Step{dir: dir, len: int64(dist)})
	}
	return steps
}

func getSteps2(lines []string) []Step {
	steps := make([]Step, 0)

	for idx := range lines {
		valStr := strings.Split(lines[idx], " ")
		lenVal := valStr[2][2:7]
		dirVal := rune(valStr[2][7])

		dist, err := strconv.ParseInt(lenVal, 16, 64)
		if err != nil {
			panic(err)
		}
		var dir Dir
		switch dirVal {
		case '0':
			dir = R
		case '1':
			dir = D
		case '2':
			dir = L
		case '3':
			dir = U
		default:
			panic("Invalid direction")
		}

		steps = append(steps, Step{dir: dir, len: dist})
	}
	return steps
}

func getVertices(steps *[]Step) []Coord {
	var x int64 = 0
	var y int64 = 0
	var maxx int64 = 0
	var maxy int64 = 0
	var minx int64 = 0
	var miny int64 = 0

	vert := make([]Coord, 0)
	for idx := range *steps {
		vert = append(vert, Coord{x: x, y: y})
		switch (*steps)[idx].dir {
		case U:
			y -= (*steps)[idx].len
		case D:
			y += (*steps)[idx].len
		case L:
			x -= (*steps)[idx].len
		case R:
			x += (*steps)[idx].len
		}
		maxx = max(x, maxx)
		maxy = max(y, maxy)
		minx = min(x, minx)
		miny = min(y, miny)
	}
	return vert
}

func getArea(vert *[]Coord) int {
	// https://web.archive.org/web/20100405070507/http://valis.cs.uiuc.edu/~sariel/research/CG/compgeom/msg00831.html
	// + half of perimeter because of wall thickness
	// + 1 for some reason :D
	n := len(*vert)
	area := 0.
	per := 0.
	for i := range *vert {
		j := (i + 1) % n
		area += float64(((*vert)[i].x * (*vert)[j].y))
		area -= float64(((*vert)[i].y * (*vert)[j].x))
		per += math.Abs(float64(((*vert)[i].x - (*vert)[j].x) + ((*vert)[i].y - (*vert)[j].y)))
	}
	return int(area/2) + int(per/2) + 1
}

func part2(lines []string) {
	steps := getSteps2(lines)
	vert := getVertices(&steps)
	area := getArea(&vert)

	fmt.Println(area)
}

func part1(lines []string) {
	steps := getSteps1(lines)
	vert := getVertices(&steps)
	area := getArea(&vert)

	fmt.Println(area)
}

func Run() {
	input := utils.GetFileContent("./src/day_18/input")
	lines := strings.Split(strings.Trim(input, "\n "), "\n")

	part1(lines)
	part2(lines)
}
