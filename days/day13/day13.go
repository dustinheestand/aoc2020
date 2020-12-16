package day13

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
	file, err := os.Open("input/day13.txt")
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
	for i := 0; ; i++ {
		t := input.tm + i
		for _, b := range input.buses {
			if b == 0 {
				continue
			}
			if t%b == 0 {
				return strconv.Itoa(i * b)
			}
		}
	}
}

// Solve2 solves.
func Solve2() string {
	res := 0
	m := 1
	for i, b := range input.buses {
		if b == 0 {
			continue
		}
		for res%b != (b-(i%b))%b {
			res += m
		}
		m *= b
	}
	return strconv.Itoa(res)
}
