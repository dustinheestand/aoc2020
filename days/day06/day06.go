package day06

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type group struct {
	people []map[rune]bool
}

var input []group

func init() {
	file, err := os.Open("input/day06.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	g := group{}
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			input = append(input, g)
			g = group{}
			continue
		}
		p := map[rune]bool{}
		for _, r := range txt {
			p[r] = true
		}
		g.people = append(g.people, p)
	}
	input = append(input, g)
}

// Solve1 solves.
func Solve1() string {
	var res int
	for _, g := range input {
		m := map[rune]bool{}
		for _, p := range g.people {
			for r := range p {
				m[r] = true
			}
		}
		res += len(m)
	}
	return strconv.Itoa(res)
}

// Solve2 solves.
func Solve2() string {
	var res int
	for _, g := range input {
		m := g.people[0]
		for _, p := range g.people[1:] {
			for r := range m {
				if !p[r] {
					delete(m, r)
				}
			}
		}
		res += len(m)
	}
	return strconv.Itoa(res)
}
