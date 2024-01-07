package day06

import (
	"advent-2023/src/utils"
	"fmt"
	"math"
	"strings"
)

type Coord struct {
	x int
	y int
}

type Cell struct {
	val bool
}

type Grid struct {
	lines        []string
	emptyColumns []int
	emptyRows    []int
}

func getDist(v1 int, v2 int, empty []int, expansion int) int {
	dist := int(math.Abs(float64(v1 - v2)))
	for _, n := range empty {
		if (v1 > n && v2 < n) || (v2 > n && v1 < n) {
			dist += expansion
		}
	}
	return dist
}

func getPairsSum(grid Grid, expansion int) int {
	h := len(grid.lines)
	w := len(grid.lines[0])
	points := make([]Coord, 0)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if grid.lines[y][x] == '#' {
				points = append(points, Coord{x: x, y: y})
			}
		}
	}

	var total int = 0
	for idx1 := range points {
		for idx2 := range points {
			if idx1 == idx2 {
				continue
			}
			dx := getDist(points[idx1].x, points[idx2].x, grid.emptyColumns, expansion)
			dy := getDist(points[idx1].y, points[idx2].y, grid.emptyRows, expansion)
			total += dx + dy
		}
	}

	return total / 2 // Remove duplicates
}

func getGrid(lines []string) Grid {
	h := len(lines)
	w := len(lines[0])

	emptyRows := make([]int, 0)
	for y := 0; y < h; y++ {
		isEmpty := true
		for x := 0; x < w; x++ {
			if lines[y][x] == '#' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyRows = append(emptyRows, y)
		}
	}

	emptyLines := make([]int, 0)
	for x := 0; x < w; x++ {
		isEmpty := true
		for y := 0; y < h; y++ {
			if lines[y][x] == '#' {
				isEmpty = false
				break
			}
		}
		if isEmpty {
			emptyLines = append(emptyLines, x)
		}
	}

	return Grid{lines: lines, emptyColumns: emptyLines, emptyRows: emptyRows}
}

func Run() {
	input := utils.GetFileContent("./src/day_11/input")
	lines := strings.Split(strings.Trim(input, "\n "), "\n")

	grid := getGrid(lines)

	fmt.Println(getPairsSum(grid, 1))      // part 1
	fmt.Println(getPairsSum(grid, 999999)) // part 2
}
