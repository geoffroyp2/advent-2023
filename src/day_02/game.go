package day02

import (
	"strconv"
	"strings"
)

type Draw struct {
	blue  int
	green int
	red   int
}

type Game struct {
	id    int
	draws []Draw
}

func newDraw(str *string) Draw {
	blue := 0
	red := 0
	green := 0
	values := strings.Split(*str, ", ")
	for idx := range values {
		val := strings.Split(values[idx], " ")
		amount, err := strconv.Atoi(val[0])
		if err != nil {
			panic(err)
		}
		switch val[1] {
		case "blue":
			blue = amount
		case "red":
			red = amount
		case "green":
			green = amount
		}
	}
	return *&Draw{blue: blue, red: red, green: green}
}

func NewGame(line *string) Game {
	s1 := strings.Split(*line, ": ")
	id, err := strconv.Atoi(strings.Split(s1[0], " ")[1])
	if err != nil {
		panic(err)
	}
	drawStr := strings.Split(s1[1], "; ")
	draws := make([]Draw, len(drawStr))
	for idx := range drawStr {
		draws[idx] = newDraw(&drawStr[idx])
	}
	return *&Game{id: id, draws: draws}
}

func (self *Game) GetMaxAmount(color string) int {
	max := 0
	for idx := range self.draws {
		switch color {
		case "red":
			{
				if max < self.draws[idx].red {
					max = self.draws[idx].red
				}
			}
		case "blue":
			{
				if max < self.draws[idx].blue {
					max = self.draws[idx].blue
				}
			}
		case "green":
			{
				if max < self.draws[idx].green {
					max = self.draws[idx].green
				}
			}
		default:
			panic("Invalid color")
		}

	}
	return max
}
