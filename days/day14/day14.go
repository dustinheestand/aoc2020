package day14

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var input schedule

type schedule struct {
	tm    int
	buses []int
}

func init() {
	file, err := os.Open("input/day14test.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fst := scanner.Text()
	i, err := strconv.Atoi(fst)
	if err != nil {
		log.Fatal(err)
	}
	input.tm = i
	scanner.Scan()
	bs := scanner.Text()
	fs := strings.FieldsFunc(bs, func(r rune) bool { return r == ',' })
	for _, f := range fs {
		i, err := strconv.Atoi(f)
		if err != nil {
			input.buses = append(input.buses, 0)
		} else {
			input.buses = append(input.buses, i)
		}
	}
}

// Solve1 solves.
func Solve1() string {
	res := 0
	return strconv.Itoa(res)
}

// Solve2 solves.
func Solve2() string {
	res := 0
	return strconv.Itoa(res)
}
