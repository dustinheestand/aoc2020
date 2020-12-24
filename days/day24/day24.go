package day24

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var input []tile

type tile []string

func init() {
	file, err := os.Open("input/day24.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		s := ""
		t := tile{}
		for _, r := range txt {
			s += string(r)
			if r == 'e' || r == 'w' {
				t = append(t, s)
				s = ""
			}
		}
		input = append(input, t)
	}
}

type coord struct {
	rise, run int
}

func (c *coord) move(dir string) {
	switch dir {
	case "ne":
		c.rise++
	case "nw":
		c.rise++
		c.run--
	case "e":
		c.run++
	case "w":
		c.run--
	case "se":
		c.rise--
		c.run++
	case "sw":
		c.rise--
	}
}

// Solve1 solves.
func Solve1() string {
	changes := make(map[coord]bool)
	for _, t := range input {
		c := &coord{}
		for _, mv := range t {
			c.move(mv)
		}
		changes[*c] = !changes[*c]
	}

	res := 0
	for _, c := range changes {
		if c {
			res++
		}
	}
	return fmt.Sprint(res)
}

func (c *coord) adjs() []coord {
	var res []coord
	for _, dir := range []string{"e", "w", "ne", "nw", "se", "sw"} {
		newC := coord{c.rise, c.run}
		newC.move(dir)
		res = append(res, newC)
	}
	return res
}

// Solve2 solves.
func Solve2() string {
	blackTiles := make(map[coord]bool)
	for _, t := range input {
		c := &coord{}
		for _, mv := range t {
			c.move(mv)
		}
		blackTiles[*c] = !blackTiles[*c]
	}

	for i := 0; i < 100; i++ {
		flips := make(map[coord]bool)
		for c, isBlack := range blackTiles {
			if !isBlack {
				delete(blackTiles, c)
				continue
			}
			for _, t := range append(c.adjs(), c) {
				adjBlack := 0
				for _, adjT := range t.adjs() {
					if blackTiles[adjT] {
						adjBlack++
					}
				}
				if blackTiles[t] {
					if adjBlack < 1 || adjBlack > 2 {
						flips[t] = true
					}
				} else {
					if adjBlack == 2 {
						flips[t] = true
					}
				}
			}
		}
		for f := range flips {
			blackTiles[f] = !blackTiles[f]
		}
	}

	res := 0
	for _, c := range blackTiles {
		if c {
			res++
		}
	}
	return fmt.Sprint(res)
}
