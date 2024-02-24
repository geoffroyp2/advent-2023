package day06

import (
	"advent-2023/src/utils"
	"fmt"
	"strings"
)

type ModuleType int

const broadcaster = "broadcaster"

const (
	Broadcaster ModuleType = iota
	FlipFlop    ModuleType = iota
	Conjunction ModuleType = iota
	Empty       ModuleType = iota
)

type Module struct {
	id      string
	input   []*Module
	output  []*Module
	modType ModuleType
	on      bool            // for FlipFlop
	last    map[string]bool // for Conjunction
	loop    int
}

type Signal struct {
	val  bool
	from string
	to   *Module
}

func getModule(line string) (Module, []string) {
	str1 := strings.Split(line, " -> ")
	modType := Broadcaster
	id := broadcaster

	if str1[0] != broadcaster {
		tchar := str1[0][0]
		switch tchar {
		case '%':
			modType = FlipFlop
		case '&':
			modType = Conjunction
		default:
			panic("Invalid module type")
		}
		id = str1[0][1:]
	}

	dest := strings.Split(str1[1], ", ")
	module := Module{id: id, input: make([]*Module, 0), output: make([]*Module, 0), modType: modType, on: false, last: make(map[string]bool)}

	return module, dest
}

func buildGraph(lines []string) *map[string]*Module {
	modules := make(map[string]*Module, 0)
	destinations := make(map[string][]string)

	for i := range lines {
		module, dest := getModule(lines[i])
		modules[module.id] = &module
		destinations[module.id] = dest
	}

	for id, dest := range destinations {
		startMod := modules[id]
		for _, d := range dest {
			destMod, exists := modules[d]
			if !exists {
				emptyMod := Module{id: d, input: make([]*Module, 0), output: make([]*Module, 0), modType: Empty, on: false, last: make(map[string]bool)}
				modules[d] = &emptyMod
				destMod = &emptyMod
			}

			startMod.output = append(startMod.output, destMod)
			destMod.input = append(destMod.input, startMod)
		}
	}

	for _, mod := range modules {
		if mod.modType == Conjunction {
			for _, dest := range mod.input {
				mod.last[dest.id] = false
			}
		}
	}

	return &modules
}

func runSignal(root *Module, it int) (int, int) {
	queue := make([]Signal, 0)
	queue = append(queue, Signal{to: root, val: false})
	lowAmount := 0  // Part 1
	highAmount := 0 // Part 1

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		if current.val {
			highAmount++
		} else {
			lowAmount++
		}

		switch current.to.modType {
		case Broadcaster:
			for _, m := range current.to.output {
				queue = append(queue, Signal{val: current.val, from: current.to.id, to: m})
			}
		case FlipFlop:
			if current.val {
				break
			}
			current.to.on = !current.to.on
			for _, m := range current.to.output {
				queue = append(queue, Signal{val: current.to.on, from: current.to.id, to: m})
			}
		case Conjunction:
			current.to.last[current.from] = current.val
			isAllOn := true
			for _, val := range current.to.last {
				if !val {
					isAllOn = false
				}
			}
			if isAllOn && current.to.loop == 0 {
				current.to.loop = it
			}
			for _, m := range current.to.output {
				queue = append(queue, Signal{val: !isAllOn, from: current.to.id, to: m})
			}
		}
	}

	return lowAmount, highAmount
}

func getLoopLengthR(node *Module, lengths *[]int) {
	if node.loop > 1 {
		*lengths = append(*lengths, node.loop)
		return
	}
	for _, n := range node.input {
		getLoopLengthR(n, lengths)
	}
}

func getLoopLength(end *Module) []int {
	lengths := make([]int, 0)
	getLoopLengthR(end, &lengths)
	return lengths
}

func part2(lines []string) {
	modules := buildGraph(lines)
	root := (*modules)[broadcaster]

	for i := 0; i < 10000; i++ {
		runSignal(root, i)
	}

	rx := (*modules)["rx"]
	loopLengths := getLoopLength(rx)
	lcm := utils.LCMForList(loopLengths)
	fmt.Println(lcm)
}

func part1(lines []string) {
	modules := buildGraph(lines)
	root := (*modules)[broadcaster]

	lowTotal := 0
	highTotal := 0
	for i := 0; i < 1000; i++ {
		low, high := runSignal(root, i)
		lowTotal += low
		highTotal += high
	}

	fmt.Println(lowTotal * highTotal)
}

func Run() {
	input := utils.GetFileContent("../src/day_20/input")
	lines := strings.Split(strings.Trim(input, "\n "), "\n")
	part1(lines)
	part2(lines)
}
