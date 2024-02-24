package day08

import (
	"advent-2023/src/utils"
	"fmt"
	"slices"
)

type Step struct {
	currentLoop int
	sinceStart  int
	node        string
}

func getLoopLength(start string, instructions string, nodes map[string]Node) int {
	steps := make([]Step, 0)
	totalLength := 0
	currentNode := start

	for {
		currentIndex := totalLength % len(instructions)
		currentInstruction := instructions[currentIndex]

		stepIdx := slices.IndexFunc(steps, func(s Step) bool { return s.node == currentNode && s.currentLoop == currentIndex })
		if stepIdx >= 0 {
			// Loop found (this only works like that because all paths loops cross exactly 1 exit)
			return totalLength - steps[stepIdx].sinceStart
		}
		steps = append(steps, Step{currentLoop: currentIndex, node: currentNode, sinceStart: totalLength})

		if currentInstruction == 'L' {
			currentNode = nodes[currentNode].left
		} else {
			currentNode = nodes[currentNode].right
		}
		totalLength++
	}
}

func part2(lines *[]string) {
	nodes := map[string]Node{}
	simNodes := make([]string, 0)

	for idx := range *lines {
		if idx < 2 {
			continue
		}
		node := getNode(&(*lines)[idx])
		nodes[node.name] = node
		if node.name[2] == 'A' {
			simNodes = append(simNodes, node.name)
		}
	}

	loopLengths := make([]int, 0)
	instructions := (*lines)[0]
	for idx := range simNodes {
		length := getLoopLength(simNodes[idx], instructions, nodes)
		loopLengths = append(loopLengths, length)
	}

	result := utils.LCMForList(loopLengths)
	fmt.Println(result)
}
