package day07

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func getCardValue1(c rune) int {
	if c >= '2' && c <= '9' {
		return int(c - '0')
	}
	switch c {
	case 'T':
		return 10
	case 'J':
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	}
	panic(fmt.Sprintf("unkown card %c", c))
}

func getHType1(cards []int) HType {
	amounts := make([]int, 15)
	for _, val := range cards {
		amounts[val]++
	}

	bestDup := -1
	for _, val := range amounts {
		if val > bestDup {
			bestDup = val
		}
	}

	switch bestDup {
	case 1:
		return High
	case 2:
		{
			pairAmount := 0
			for _, val := range amounts {
				if val == 2 {
					pairAmount++
				}
			}
			if pairAmount == 1 {
				return OnePair
			} else if pairAmount == 2 {
				return TwoPairs
			}
			panic("Invalid combination for pairs")
		}
	case 3:
		{
			pairAmount := 0
			for _, val := range amounts {
				if val == 2 {
					pairAmount++
				}
			}
			if pairAmount == 0 {
				return ThreeKind
			} else if pairAmount == 1 {
				return FullHouse
			}
			panic("Invalid combination for threes")
		}
	case 4:
		return FourKind
	case 5:
		return FiveKind
	}
	panic("No combination found")
}

func getHand1(line *string) Hand {
	hStr := strings.Split(*line, " ")
	hand := Hand{}

	for _, c := range hStr[0] {
		hand.cards = append(hand.cards, getCardValue1(c))
	}
	hand.htype = getHType1(hand.cards)

	bet, err := strconv.Atoi(hStr[1])
	if err != nil {
		panic(err)
	}
	hand.bet = bet

	return hand
}

func part1(lines *[]string) {
	hands := make([]Hand, len(*lines))
	for idx := range *lines {
		hand := getHand1(&(*lines)[idx])
		hands[idx] = hand
	}

	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].htype == hands[j].htype {
			for idx := range hands[i].cards {
				if hands[i].cards[idx] == hands[j].cards[idx] {
					continue
				}
				return hands[i].cards[idx] < hands[j].cards[idx]
			}
		}
		return hands[i].htype < hands[j].htype
	})

	total := 0
	for idx := range hands {
		total += hands[idx].bet * (idx + 1)
	}
	fmt.Println(total)
}
