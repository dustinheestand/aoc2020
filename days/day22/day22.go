package day22

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var input [][]int

type hashes struct {
	p0, p1 string
}

func init() {
	file, err := os.Open("input/day22.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	input = [][]int{{}, {}}
	idx := 0
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.HasPrefix(txt, "P") {
			continue
		}
		if len(txt) == 0 {
			idx++
			continue
		}
		i, err := strconv.Atoi(txt)
		if err != nil {
			log.Fatal(err)
		}
		input[idx] = append(input[idx], i)
	}
}

// Solve1 solves.
func Solve1() string {
	decks := [][]int{make([]int, len(input[0])), make([]int, len(input[1]))}
	copy(decks[0], input[0])
	copy(decks[1], input[1])
	for len(decks[0]) != 0 && len(decks[1]) != 0 {
		a, b := decks[0][0], decks[1][0]
		decks[0], decks[1] = decks[0][1:], decks[1][1:]
		if a > b {
			decks[0] = append(decks[0], []int{a, b}...)
		} else {
			decks[1] = append(decks[1], []int{b, a}...)
		}
	}

	var d []int
	if len(decks[0]) == 0 {
		d = decks[1]
	} else {
		d = decks[0]
	}
	return fmt.Sprint(score(d))
}

func score(cs []int) int {
	var score int
	for i, c := range cs {
		score += c * (len(cs) - i)
	}
	return score
}

// Solve2 solves.
func Solve2() string {
	decks := [][]int{make([]int, len(input[0])), make([]int, len(input[1]))}
	copy(decks[0], input[0])
	copy(decks[1], input[1])

	winner := play(decks, make(map[hashes]struct{}))
	var d []int
	if winner {
		d = decks[0]
	} else {
		d = decks[1]
	}
	return fmt.Sprint(score(d))
}

func play(decks [][]int, seen map[hashes]struct{}) bool {
	var gameWinner bool
	for len(decks[0]) != 0 && len(decks[1]) != 0 {
		strs := hashes{hash(decks[0]), hash(decks[1])}
		if _, ok := seen[strs]; ok {
			gameWinner = true
			break
		}
		seen[strs] = struct{}{}

		var winner bool
		a, b := decks[0][0], decks[1][0]
		strs.p0, strs.p1 = strs.p0[2:], strs.p1[2:]
		decks[0], decks[1] = decks[0][1:], decks[1][1:]
		if a > len(decks[0]) || b > len(decks[1]) {
			winner = a > b
		} else {
			newDecks := [][]int{make([]int, a), make([]int, b)}
			copy(newDecks[0], decks[0][:a])
			copy(newDecks[1], decks[1][:b])
			winner = play(newDecks, make(map[hashes]struct{}))
		}
		if winner {
			decks[0] = append(decks[0], []int{a, b}...)
			strs.p0 += hash([]int{a, b})
		} else {
			decks[1] = append(decks[1], []int{b, a}...)
			strs.p1 += hash([]int{b, a})
		}
	}

	return gameWinner || len(decks[0]) != 0
}

func hash(is []int) string {
	sb := strings.Builder{}
	for _, i := range is {
		sb.WriteString(fmt.Sprintf("%02d", i))
	}
	return sb.String()
}
