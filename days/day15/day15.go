package day15

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var input []int

func init() {
	file, err := os.Open("input/day15.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		fs := strings.FieldsFunc(txt, func(r rune) bool { return r == ',' })
		for _, f := range fs {
			i, err := strconv.Atoi(f)
			if err != nil {
				log.Fatal(err)
			}
			input = append(input, i)
		}
	}
}

// Solve1 solves.
func Solve1() string {
	return solve(2020)
}

// Solve2 solves.
func Solve2() string {
	return solve(30000000)
}

func solve(pos int) string {
	seen := make(map[int]int)
	next := -1
	for i := 0; ; i++ {
		if i < len(input) {
			seen[input[i]] = i + 1
			continue
		}
		if i == pos {
			return fmt.Sprint(next)
		}
		last, ok := seen[next]
		since := i - last
		if i != len(input) {
			seen[next] = i
		}
		if !ok {
			next = 0
		} else {
			next = since
		}
	}
}
