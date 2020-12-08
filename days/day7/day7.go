package day7

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type contents map[string]int

var input = map[string]contents{}

func init() {
	file, err := os.Open("input/day7.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		fields := strings.Fields(txt)
		color := strings.Join(fields[:2], " ")
		c := make(contents)
		if fields[4] != "no" {
			for i := 4; i < len(fields); i += 4 {
				num, err := strconv.Atoi(fields[i])
				if err != nil {
					log.Fatal(err)
				}
				c[strings.Join([]string{fields[i+1], fields[i+2]}, " ")] = num
			}
		}
		if _, ok := input[color]; ok {
			log.Fatal("duplicate key")
		}
		input[color] = c
	}
}

// Solve1 solves.
func Solve1() string {
	memo := make(map[string]bool)
	for bag, c := range input {
		_ = canContain("shiny gold", bag, c, memo)
	}
	var res int
	for _, val := range memo {
		if val {
			res++
		}
	}
	return strconv.Itoa(res)
}

func canContain(target, bag string, c contents, m map[string]bool) bool {
	if val, ok := m[bag]; ok {
		return val
	}
	var val bool
	if len(c) == 0 {
		val = false
	} else if _, ok := c[target]; ok {
		val = true
	} else {
		for innerBag := range c {
			if canContain(target, innerBag, input[innerBag], m) {
				val = true
			}
		}
	}
	m[bag] = val
	return val
}

// Solve2 solves.
func Solve2() string {
	memo := make(map[string]int)
	res := contains("shiny gold", input["shiny gold"], memo)
	return strconv.Itoa(res)
}

func contains(bag string, c contents, m map[string]int) int {
	var val int
	for innerBag, count := range c {
		var perBag int
		if val, ok := m[innerBag]; ok {
			perBag = val
		} else {
			perBag = contains(innerBag, input[innerBag], m)
		}
		val += count * (perBag + 1)
	}
	m[bag] = val
	return val
}
