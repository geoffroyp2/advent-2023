package day08

import (
	"fmt"
	"strings"
)

func getNode(line *string) Node {
	nStr := strings.Split(*line, " = ")
	childStr := strings.Split(nStr[1], ", ")
	left := strings.Split(childStr[0], "(")[1]
	right := strings.Split(childStr[1], ")")[0]

	return Node{name: nStr[0], left: left, right: right}
}

func part1(lines *[]string) {

	nodes := map[string]Node{}

	for idx := range *lines {
		if idx < 2 {
			continue
		}
		node := getNode(&(*lines)[idx])
		nodes[node.name] = node
	}

	current := nodes["AAA"]
	steps := 0
	instructions := (*lines)[0]
	for {
		// fmt.Println(current.name)
		nextI := instructions[steps%len(instructions)]
		if nextI == 'L' {
			current = nodes[current.left]
		} else {
			current = nodes[current.right]
		}
		if current.name == "ZZZ" {
			break
		}
		steps++
	}

	fmt.Println(steps + 1)
}
