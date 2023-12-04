package day03

import (
	"advent-2023/src/utils"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Game struct {
	id      int
	winning []int
	got     []int
}

func get_game(line *string) Game {
	game_str_re := regexp.MustCompile(": +")
	spaces_re := regexp.MustCompile(" +")
	num_str_re := regexp.MustCompile(" +\\| +")

	game_str := game_str_re.Split(*line, -1)
	id, err := strconv.Atoi(spaces_re.Split(game_str[0], -1)[1])
	if err != nil {
		panic(err)
	}

	num_str := num_str_re.Split(game_str[1], -1)
	win_num_str := spaces_re.Split(num_str[0], -1)
	got_num_str := spaces_re.Split(num_str[1], -1)

	game := Game{id: id}

	for idx := range win_num_str {
		val, errn := strconv.Atoi(win_num_str[idx])
		if errn != nil {
			panic(errn)
		}
		game.winning = append(game.winning, val)
	}
	for idx := range got_num_str {
		val, errn := strconv.Atoi(got_num_str[idx])
		if errn != nil {
			panic(errn)
		}
		game.got = append(game.got, val)
	}
	return game
}

func get_score(game *Game) int {
	score := 0

	for _, val := range game.winning {
		match := slices.IndexFunc(game.got, func(n int) bool { return val == n })
		if match >= 0 {
			score++
		}
	}
	return score
}

func get_score_part1(score int) int {
	if score == 0 {
		return 0
	}
	total := 1
	for i := 1; i < score; i++ {
		total *= 2
	}
	return total
}

func part_1(lines *[]string) {
	total := 0
	for idx := range *lines {
		game := get_game(&(*lines)[idx])
		score := get_score_part1(get_score(&game))
		total += score
	}

	println(total)
}

func part_2(lines *[]string) {
	l := len(*lines)
	card_amounts := make([]int, l)
	for idx := range card_amounts {
		card_amounts[idx] = 1
	}

	for idx := range *lines {
		game := get_game(&(*lines)[idx])
		score := get_score(&game)
		for idx2 := 1; idx2 <= score && (idx2+idx) < l; idx2++ {
			card_amounts[idx+idx2] += card_amounts[idx]
		}
	}

	total := 0
	for idx := range card_amounts {
		total += card_amounts[idx]
	}
	println(total)
}

func Run() {
	input := utils.GetFileContent("./src/day_04/input")
	lines := strings.Split(strings.Trim(input, "\n "), "\n")

	part_1(&lines)
	part_2(&lines)
}
