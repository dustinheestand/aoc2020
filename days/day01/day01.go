package day01

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

var input sort.IntSlice

func init() {
	file, err := os.Open("input/day01.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		i, err := strconv.Atoi(txt)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, i)
	}
	input.Sort()
}

// Solve1 solves.
func Solve1() string {
	i, j := 0, len(input)-1
	for {
		a, b := input[i], input[j]
		if a+b < 2020 {
			i++
		} else if a+b > 2020 {
			j--
		} else {
			return strconv.Itoa(a + b)
		}
	}
}

// Solve2 solves.
func Solve2() string {
	for end := len(input) - 1; ; end-- {
		i, j := 0, end-1
		stable := input[end]
		target := 2020 - stable
		for i < j {
			a, b := input[i], input[j]
			if a+b < target {
				i++
			} else if a+b > target {
				j--
			} else if a+b == target {
				return strconv.Itoa(a * b * stable)
			}
		}
	}
}
