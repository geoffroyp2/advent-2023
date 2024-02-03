package day06

import (
	"advent-2023/src/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

type Dir int

const (
	U Dir = iota
	D Dir = iota
	L Dir = iota
	R Dir = iota
)

type Visit struct {
	val int
	dir Dir
}

type Cell struct {
	val  int
	dist float64
	vis  []Visit
}

type Position struct {
	x   int
	y   int
	dir Dir
	val int
}

type Grid [][]Cell

func getGrid(lines []string) Grid {
	grid := make([][]Cell, len(lines))
	for y := range lines {
		row := make([]Cell, len(lines[y]))
		for x := range lines[y] {
			row[x] = Cell{val: int(lines[y][x] - 48), dist: math.Sqrt(float64(x*x + y*y))}
		}
		grid[y] = row
	}
	return grid
}

func canL(pos Position, amount int) bool {
	return pos.x >= amount
}
func canR(pos Position, amount, w int) bool {
	return pos.x <= w-1-amount
}
func canU(pos Position, amount int) bool {
	return pos.y >= amount
}
func canD(pos Position, amount, h int) bool {
	return pos.y <= h-1-amount
}

func getNextPositions(grid *Grid, pos Position) []Position {
	h := len(*grid)
	w := len((*grid)[0])

	positions := make([]Position, 0)
	switch pos.dir {
	case U:
		val := 0
		for dy := 0; dy < 3; dy++ {
			if !canU(pos, dy) {
				break
			}
			val += (*grid)[pos.y-dy][pos.x].val
			if canR(pos, 1, w) {
				positions = append(positions, Position{x: pos.x + 1, y: pos.y - dy, val: pos.val + val, dir: R})
			}
			if canL(pos, 1) {
				positions = append(positions, Position{x: pos.x - 1, y: pos.y - dy, val: pos.val + val, dir: L})
			}
		}
	case D:
		val := 0
		for dy := 0; dy < 3; dy++ {
			if !canD(pos, dy, h) {
				break
			}
			val += (*grid)[pos.y+dy][pos.x].val
			if canR(pos, 1, w) {
				positions = append(positions, Position{x: pos.x + 1, y: pos.y + dy, val: pos.val + val, dir: R})
			}
			if canL(pos, 1) {
				positions = append(positions, Position{x: pos.x - 1, y: pos.y + dy, val: pos.val + val, dir: L})
			}
		}
	case L:
		val := 0
		for dx := 0; dx < 3; dx++ {
			if !canL(pos, dx) {
				break
			}
			val += (*grid)[pos.y][pos.x-dx].val
			if canU(pos, 1) {
				positions = append(positions, Position{x: pos.x - dx, y: pos.y - 1, val: pos.val + val, dir: U})
			}
			if canD(pos, 1, h) {
				positions = append(positions, Position{x: pos.x - dx, y: pos.y + 1, val: pos.val + val, dir: D})
			}
		}
	case R:
		val := 0
		for dx := 0; dx < 3; dx++ {
			if !canR(pos, dx, w) {
				break
			}
			val += (*grid)[pos.y][pos.x+dx].val
			if canU(pos, 1) {
				positions = append(positions, Position{x: pos.x + dx, y: pos.y - 1, val: pos.val + val, dir: U})
			}
			if canD(pos, 1, h) {
				positions = append(positions, Position{x: pos.x + dx, y: pos.y + 1, val: pos.val + val, dir: D})
			}
		}
	}
	return positions
}

func part1(lines []string) {
	grid := getGrid(lines)

	for y := range grid {
		for x := range grid[y] {
			fmt.Print(grid[y][x].val)
		}
		fmt.Print("\n")
	}

	queue := make([]Position, 0)
	queue = append(queue, Position{x: 0, y: 0, dir: R, val: -grid[0][0].val}, Position{x: 0, y: 0, dir: D, val: -grid[0][0].val})

	it := 0
	for len(queue) != 0 {
		it++

		pos := queue[0]
		queue = queue[1:]

		if (it % 1000) == 0 {
			fmt.Printf("%v %v %+v\n", it, len(queue), pos)
		}

		if pos.x == len(grid[0])-1 && pos.y == len(grid)-1 {
			break
		}

		cell := &grid[pos.y][pos.x]

		// Mark cell as visited
		visited := false
		shouldStop := false
		for vidx := range cell.vis {
			if cell.vis[vidx].dir == pos.dir {
				if cell.vis[vidx].val == pos.val {
					shouldStop = true
				}
				cell.vis[vidx].val = min(cell.vis[vidx].val, pos.val)
				visited = true
				break
			}
		}
		if shouldStop {
			continue
		}
		if !visited {
			cell.vis = append(cell.vis, Visit{val: pos.val, dir: pos.dir})
		}

		// Add next cells
		next := getNextPositions(&grid, pos)
		for nidx := range next {
			isDifferent := true
			for qidx := range queue {
				if next[nidx].x == queue[qidx].x && next[nidx].y == queue[qidx].y && next[nidx].dir == queue[qidx].dir {
					isDifferent = false
					break
				}
			}
			if isDifferent {
				queue = append(queue, next[nidx])
			} else {
				// panic("ALREADY IN QUEUE")
			}
		}
		// fmt.Printf("%+v\n", len(queue))

		sort.SliceStable(queue, func(i, j int) bool {
			if queue[i].val == queue[j].val {
				return grid[queue[i].y][queue[i].x].dist > grid[queue[j].y][queue[j].x].dist
			}
			return queue[i].val < queue[j].val
		})
	}

}

func Run() {
	input := utils.GetFileContent("./src/day_17/input")
	lines := strings.Split(strings.Trim(input, "\n "), "\n")

	part1(lines)
}
