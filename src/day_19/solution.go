package day06

import (
	"advent-2023/src/utils"
	"fmt"
	"strconv"
	"strings"
)

type Condition struct {
	letter    byte
	isGreater bool
	value     int
}

type Rule struct {
	condition Condition
	target    string
}

type Pipeline struct {
	name  string
	rules []Rule
}

type Item struct {
	x, m, a, s int
}

type ItemBound struct {
	x, m, a, s,
	X, M, A, S int
}

func getRule(str string) Rule {
	split := strings.Split(str, ":")
	if len(split) > 1 {
		comp := split[0][1]
		letter := split[0][0]
		value, err := strconv.Atoi(split[0][2:])
		if err != nil {
			panic(err)
		}
		target := split[1]
		return Rule{target: target, condition: Condition{value: value, letter: letter, isGreater: comp == '>'}}
	} else {
		return Rule{target: split[0]}
	}
}

func getPipelines(lines []string) []Pipeline {
	pipelines := make([]Pipeline, 0)
	for idx := range lines {
		split := strings.Split(lines[idx], "{")
		name := split[0]
		rulestr := strings.Split(split[1][:len(split[1])-1], ",")
		rules := make([]Rule, len(rulestr))
		for ridx := range rulestr {
			rules[ridx] = getRule(rulestr[ridx])
		}
		pipelines = append(pipelines, Pipeline{name: name, rules: rules})
	}
	return pipelines
}

func getItems(lines []string) []Item {
	items := make([]Item, len(lines))
	for idx := range lines {
		item := Item{}
		itemstr := strings.Split(lines[idx][1:len(lines[idx])-1], ",")
		for iidx := range itemstr {
			split := strings.Split(itemstr[iidx], "=")
			val, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			switch split[0][0] {
			case 'x':
				item.x = val
			case 'm':
				item.m = val
			case 'a':
				item.a = val
			case 's':
				item.s = val
			}
		}
		items[idx] = item
	}
	return items
}

func isConditionTrue(item Item, condition Condition) bool {
	switch condition.letter {
	case 'x':
		if condition.isGreater {
			return item.x > condition.value
		} else {
			return item.x < condition.value
		}
	case 'm':
		if condition.isGreater {
			return item.m > condition.value
		} else {
			return item.m < condition.value
		}
	case 'a':
		if condition.isGreater {
			return item.a > condition.value
		} else {
			return item.a < condition.value
		}
	case 's':
		if condition.isGreater {
			return item.s > condition.value
		} else {
			return item.s < condition.value
		}
	}
	panic("Invalid condition")
}

func isAccepted(pipelines []Pipeline, item Item) bool {
	EMPTY := Condition{}
	pname := "in"
	for {
		pipeline := pipelines[0]
		// Find current pipeline
		for pidx := range pipelines {
			if pipelines[pidx].name == pname {
				pipeline = pipelines[pidx]
			}
		}
		// Follow rules
		shouldGoNext := false
		for ridx := range pipeline.rules {
			// End of pipeline
			if pipeline.rules[ridx].condition == EMPTY {
				pname = pipeline.rules[ridx].target // Go next pipeline
				shouldGoNext = true
			} else {
				// Check condition
				if isConditionTrue(item, pipeline.rules[ridx].condition) {
					pname = pipeline.rules[ridx].target // Go next pipeline
					shouldGoNext = true
				} else {
				}
			}
			if shouldGoNext {
				switch pname {
				case "R":
					return false // Rejected
				case "A":
					return true // Accepted
				}
				break
			}
		}
	}
}

func findValidPaths(pipelines *[]Pipeline, paths *[]ItemBound, item ItemBound) {
	fmt.Println(item)

	EMPTY := Condition{}
	pname := "in"
	for {
		pipeline := (*pipelines)[0]
		// Find current pipeline
		for pidx := range *pipelines {
			if (*pipelines)[pidx].name == pname {
				pipeline = (*pipelines)[pidx]
				break
			}
		}
		// Follow rules
		shouldGoNext := false
		for ridx := range pipeline.rules {
			cond := pipeline.rules[ridx].condition
			// End of pipeline
			if cond == EMPTY {
				pname = pipeline.rules[ridx].target // Go next pipeline
				shouldGoNext = true
			} else {
				// Check condition
				switch cond.letter {
				case 'x':
					next1 := item
					next2 := item
					*paths = append(*paths, next1, next2)
				case 'm':
				case 'a':
				case 's':
				}
				// if isConditionTrue(item, pipeline.rules[ridx].condition) {
				// 	pname = pipeline.rules[ridx].target // Go next pipeline
				// 	shouldGoNext = true
				// }
			}
			if shouldGoNext {
				if pname == "R" || pname == "A" {
					return
				}
				break
			}
		}
	}

}

func part2(pipelines []Pipeline) {

}

func part1(pipelines []Pipeline, items []Item) {
	total := 0
	for idx := range items {
		if isAccepted(pipelines, items[idx]) {
			total += items[idx].x
			total += items[idx].m
			total += items[idx].a
			total += items[idx].s

		}
	}
	fmt.Println(total)
}

func Run() {
	input := utils.GetFileContent("./src/day_19/input")
	blocks := strings.Split(strings.Trim(input, "\n "), "\n\n")
	block1 := strings.Split(blocks[0], "\n")
	block2 := strings.Split(blocks[1], "\n")

	part1(getPipelines(block1), getItems(block2))
}
