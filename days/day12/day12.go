package day12

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

type ship struct {
	x, y   int
	facing int
}

type move struct {
	dir byte
	num int
}

var input []move

func init() {
	file, err := os.Open("input/day12.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		m := move{}
		m.dir = txt[0]
		m.num, err = strconv.Atoi(txt[1:])
		if err != nil {
			log.Fatal(err)
		}
		if m.dir == 'L' {
			m.dir = 'R'
			m.num = 360 - m.num
		}
		input = append(input, m)
	}
}

// Solve1 solves.
func Solve1() string {
	s := ship{facing: 90}
	for _, m := range input {
		switch m.dir {
		case 'N':
			s.y += m.num
		case 'S':
			s.y -= m.num
		case 'E':
			s.x += m.num
		case 'W':
			s.x -= m.num
		case 'R':
			s.facing += m.num
		case 'F':
			s.x += m.num * int(math.Sin(float64(s.facing)*math.Pi/180))
			s.y += m.num * int(math.Cos(float64(s.facing)*math.Pi/180))
		}
	}
	return strconv.Itoa(int(math.Abs(float64(s.x)) + math.Abs(float64(s.y))))
}

// Solve2 solves.
func Solve2() string {
	s := ship{}
	w := ship{x: 10, y: 1}
	for _, m := range input {
		switch m.dir {
		case 'N':
			w.y += m.num
		case 'S':
			w.y -= m.num
		case 'E':
			w.x += m.num
		case 'W':
			w.x -= m.num
		case 'R':
			for m.num > 0 {
				w.x, w.y = w.y, -w.x
				m.num -= 90
			}
		case 'F':
			s.x += m.num * w.x
			s.y += m.num * w.y
		}
	}
	return strconv.Itoa(int(math.Abs(float64(s.x)) + math.Abs(float64(s.y))))
}
